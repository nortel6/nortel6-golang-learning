package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// Only the sender should close the channel
	close(c)
}

func fibonacci_select(c, quit chan int) {
	x, y := 0, 1
	for {
		// Blocks until the cases can run
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	// Oh this is cool...
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	// Runs in another thread
	go say("world")
	say("hello")

	// Use channel to synchronize memory
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	// Make with buffer size
	// Block send if full, block wait if empty
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fibc := make(chan int, 10)
	go fibonacci(cap(fibc), fibc)
	// Will continue forever until channel closed
	// Normally, you don't have to close channels
	// But for cases like this, you must
	// Can also do a check like v, ok := <-ch
	for i := range fibc {
		fmt.Println(i)
	}

	c_select := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c_select)
		}
		quit <- 0
	}()
	fibonacci_select(c_select, quit)

	// Test mutex
	counter := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go counter.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(counter.Value("somekey"))

	// If no other cases can run, runs default
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
