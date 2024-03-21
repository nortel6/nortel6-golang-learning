package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) push(val T) {
	curr := l
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &List[T]{nil, val}
	fmt.Printf("Pushed %T %v\n", val, val)
}

func (l *List[T]) pop() (T, error) {
	val := l.val
	// Dereference to change the value of List
	*l = *l.next
	fmt.Printf("Popped %T %v\n", val, val)
	return val, nil
}

// func (l *List[T]) pushAt() {

// }

// func (l *List[T]) popAt() T {
// 	return nil
// }

func (l *List[T]) print() {
	curr := l
	for curr != nil {
		fmt.Printf("%T: %v ,", curr.val, curr.val)
		curr = curr.next
	}
	fmt.Println()
}

func main() {
	list := List[any]{nil, 1}
	list.print()
	list.push("Hello")
	list.push("World")
	list.print()
	list.pop()
	list.push(true)
	list.pop()
	list.print()
	list.pop()
	list.print()
}
