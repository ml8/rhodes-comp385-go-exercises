package list

import (
	"fmt"
	"strings"
)

// Uppper case -> visible outside of this package.
// Lower case -> not visible outside of this package.

// Creates a new IntList, implemented with a linked list.
func NewLinkedIntList() IntList {
	// Note that a pointer to an linkedIntList is an IntList, because all methods
	// have pointer receivers (func (l *linkedIntList) ...).
	return &linkedIntList{}
}

// Node type used in linked list.
type node struct {
	val  int
	next *node
	prev *node
}

// Linked list implementation of IntList.
type linkedIntList struct {
	Head *node
	Tail *node
}

func (l *linkedIntList) PushFront(val int) {
	// Create node, n is a pointer to it (see the &).
	n := &node{val: val} // next and prev are nil by default

	// Head == nil -> list is empty, set Tail pointer
	if l.Head == nil {
		l.Tail = n
	}

	n.next = l.Head
	l.Head = n

	// link next.prev to this.
	if n.next != nil {
		n.next.prev = n
	}
}

func (l *linkedIntList) PushBack(val int) {
	n := &node{val: val}

	// Head == nil -> list is empty, set Head pointer
	if l.Head == nil {
		l.Head = n
	}

	n.prev = l.Tail
	l.Tail = n

	// link prev.next to this.
	if n.prev != nil {
		n.prev.next = n
	}
}

func (l *linkedIntList) PopFront() (val int, err error) {
	if l.Head == nil {
		err = EmptyListErr
		return
	}

	val = l.Head.val

	if l.Head == l.Tail {
		// If list is a singleton, ensure Tail no longer points to node.
		l.Tail = nil
	}

	l.Head = l.Head.next

	if l.Head != nil {
		// free previous Head.
		l.Head.prev = nil
	}

	return
}

func (l *linkedIntList) PopBack() (val int, err error) {
	if l.Head == nil {
		err = EmptyListErr
		return
	}

	val = l.Tail.val

	if l.Head == l.Tail {
		// Ensure Head no longer points to node.
		l.Head = nil
	}

	l.Tail = l.Tail.prev

	if l.Tail != nil {
		// free previous Tail.
		l.Tail.next = nil
	}
	return
}

func (l *linkedIntList) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	tmp := l.Head
	for tmp != nil {
		sb.WriteString(fmt.Sprintf("%d", tmp.val))
		tmp = tmp.next
		if tmp != nil {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}
