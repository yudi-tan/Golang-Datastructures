package main

import "fmt"

// Linkedlist struct for type int only since Go has no generic type
type Linkedlist struct {
	head int
	tail *Linkedlist
}

//adds an item to the end of the linked list
func (ll *Linkedlist) addToEnd(item int) {
	if ll.tail == nil {
		ll.tail = &Linkedlist{head: item, tail: nil}
	} else {
		ll.tail.addToEnd(item)
	}
}

func (ll *Linkedlist) addAtIndex(item int, index int) {
	if index == 0 {
		tmp := *ll
		*ll = Linkedlist{head: item, tail: &tmp}
	} else if index < 0 || (index > 0 && ll.tail == nil) {
		fmt.Printf("%s\n", "\"Out of bounds\"")
	} else {
		ll.tail.addAtIndex(item, index-1)
	}
}

func (ll *Linkedlist) printLinkedlist() {
	if ll.tail == nil {
		fmt.Println(ll.head)
	} else {
		fmt.Println(ll.head)
		ll.tail.printLinkedlist()
	}
}

func (ll *Linkedlist) deleteAtIndex(index int) {
	if index == 0 {
		*ll = *ll.tail
	} else {
		ll.tail.deleteAtIndex(index - 1)
	}
}

func (ll *Linkedlist) deleteTail() {
	if ll.tail == nil {
	} else if ll.tail != nil && ll.tail.tail == nil {
		ll.tail = nil
	} else {
		ll.tail.deleteTail()
	}
}

func main() {
	test := Linkedlist{head: 1, tail: nil}
	fmt.Println("Test 1: ")
	test.printLinkedlist()
	test.addToEnd(2)
	test.addToEnd(3)
	test.addToEnd(4)
	fmt.Println("Test 2: ")
	test.printLinkedlist()
	test.addAtIndex(8, 0)
	fmt.Println("Test 3: ")
	test.printLinkedlist()
	fmt.Println("Test 4: ")
	test.deleteTail()
	test.printLinkedlist()
	fmt.Println("Test 5: ")
	test.deleteAtIndex(2)
	test.printLinkedlist()
}
