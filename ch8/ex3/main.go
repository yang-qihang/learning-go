package main

import "fmt"

type LinkedListNode[T comparable] struct {
	data T
	// can be null
	next *LinkedListNode[T]
}

type LinkedList[T comparable] struct {
	// can be null
	head *LinkedListNode[T]
}

func newLinkedList[T comparable]() LinkedList[T] {
	return LinkedList[T]{
		head: nil,
	}
}

func (l *LinkedList[T]) Insert(elem T, pos int) {
	if pos < 0 {
		panic("invalid position")
	}

	// The case where pos == 0 requires special handling!
	if pos == 0 {
		l.head = &LinkedListNode[T]{
			data: elem,
			next: l.head,
		}
		return
	}

	curNode := l.head

	// e.g. pos = 1 -> 0 iterations
	for range pos - 1 {
		if curNode == nil {
			panic("invalid position")
		}
		curNode = curNode.next
	}

	if curNode == nil {
		panic("invalid position")
	}

	curNode.next = &LinkedListNode[T]{
		data: elem,
		next: curNode.next,
	}
}

func (l *LinkedList[T]) Add(elem T) {
	if l.head == nil {
		l.head = &LinkedListNode[T]{
			data: elem,
			next: nil,
		}
		return
	}

	curNode := l.head
	for curNode.next != nil {
		curNode = curNode.next
	}

	curNode.next = &LinkedListNode[T]{
		data: elem,
		next: nil,
	}
}

func (l LinkedList[T]) Index(elem T) int {
	tmpNode := l.head
	for count := 0; tmpNode != nil; tmpNode, count = tmpNode.next, count+1 {
		if tmpNode.data == elem {
			return count
		}
	}
	return -1
}

func main() {
	list := newLinkedList[string]()
	list.Add("Hello")      // 0
	list.Add("World")      // 1 -> 2
	list.Add("English")    // 2 -> 3
	list.Insert("Fast", 1) // 1

	fmt.Println(list.Index("Hello"))   // 0
	fmt.Println(list.Index("World"))   // 2
	fmt.Println(list.Index("Fast"))    // 1
	fmt.Println(list.Index("English")) // 3
}
