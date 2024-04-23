package main

import "fmt"

type LinkedList struct {
	Head *Node
	Size int
}

// Node represents a link (element) in our linked list
type Node struct {
	Value string
	Next  *Node
}

// Insert Adds a new element to the start of the linked list
func (l *LinkedList) Insert(elem string) {
	node := Node{
		Value: elem,
		Next:  l.Head,
	}
	l.Head = &node
	l.Size++
}

// What happens on Insert() ?
// "TestValue" is HEAD
// Insert("AnotherValue")
// "AnotherValue" -> "TestValue"
// HEAD is now "AnotherValue"
// Linked list looks like:
// "AnotherValue" -> "TestValue" -> nil

// DeleteFirst Removes the first element from the linked list
func (l *LinkedList) DeleteFirst() {
	l.Head = l.Head.Next
	l.Size--
}

// List Iterates through all linked list elements and prints them
func (l *LinkedList) List() {
	current := l.Head
	for current != nil {
		fmt.Printf("%+v\n", current)
		current = current.Next
	}
}

// Search Searches through every element of the linked list
// and returns the first element that matches the string (elem)
func (l *LinkedList) Search(elem string) *Node {
	current := l.Head
	for current != nil {
		if current.Value == elem {
			return current
		}
		current = current.Next
	}

	return nil
}

// Delete Removes an element from the linked list if it matches
func (l *LinkedList) Delete(elem string) {
	previous := l.Head
	current := l.Head
	for current != nil {
		if current.Value == elem {
			// Updates the Head of the list
			previous.Next = current.Next
			l.Size--
		}
		previous = current
		current = current.Next
	}
}

func main() {
	// A linked list is a series of nodes.
	// Each node has a reference to the next one.
	// First node is the "Head", the last is the "Tail".
	// Note: In a doubly linked list a single node has a reference
	// to both the nexy and previous node. This current file is a singly linked list.

	ll := LinkedList{}
	ll.Insert("one")
	ll.Insert("two")
	ll.Insert("three")

	ll.Delete("two")
	ll.List() // &{Value:three Next:0xc00000c030}

	element := ll.Search("one")
	if element != nil {
		fmt.Printf("Found -> %+v\n", element) // Found -> &{Value:one Next:<nil>}
	}

	fmt.Println(ll.Size) // 2
}
