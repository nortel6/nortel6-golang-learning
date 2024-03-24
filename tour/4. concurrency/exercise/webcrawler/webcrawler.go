package main

import (
	"fmt"
	"sync"
)

type SafeCache struct {
	mu sync.Mutex
	v  map[string]struct{}
}

func NewSafeCache() SafeCache {
	return SafeCache{v: make(map[string]struct{})}
}

func (c *SafeCache) Add(key string) {
	c.mu.Lock()
	c.v[key] = struct{}{}
	c.mu.Unlock()
}

func (c *SafeCache) Has(key string) bool {
	c.mu.Lock()
	_, ok := c.v[key]
	c.mu.Unlock()
	return ok
}

// This is needed because there's possibility of race condition
// Between if Has() and Add() in normal usage
func (c *SafeCache) AddIfNotExists(key string) bool {
	c.mu.Lock()
	_, ok := c.v[key]
	if !ok {
		c.v[key] = struct{}{}
	}
	c.mu.Unlock()
	// Return true if added
	return !ok
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
// Using channels for waiting, PAIN
func Crawl(url string, depth int, fetcher Fetcher, ret chan string) {
	defer close(ret)

	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		ret <- fmt.Sprint(err)
		return
	}
	ret <- fmt.Sprintf("found: %s %q", url, body)
	channels := []chan string{}
	for _, u := range urls {
		if cache.AddIfNotExists(u) {
			channel := make(chan string)
			channels = append(channels, channel)
			go Crawl(u, depth-1, fetcher, channel)
		}
	}

	// Channels is slice
	// Wait for all subchannels to close
	for _, c := range channels {
		for u := range c {
			ret <- u
		}
	}
	return
}

// I love workgroup
func CrawlWG(url string, depth int, fetcher Fetcher) {
	defer wg.Done()
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if cache.AddIfNotExists(u) {
			wg.Add(1)
			go CrawlWG(u, depth-1, fetcher)
		}
	}

	return
}

// I will just make this global
var cache = NewSafeCache()

// Using another method... seems more canonical
var wg sync.WaitGroup

func main() {
	cache.Add("https://golang.org/")
	// I spent such a long time scratching my head with channels alone...
	// urls := make(chan string)
	// go Crawl("https://golang.org/", 4, fetcher, urls)

	// I didn't want to do this, but it doesn't seem like there's another choice
	// to wait with channels... at least from my searches
	// for u := range urls {
	// 	fmt.Println(u)
	// }

	// Work Group is top tier solution for this
	wg.Add(1)
	go CrawlWG("https://golang.org/", 4, fetcher)
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
