package list

import (
	"errors"
)

// Empty list error.
var EmptyListErr = errors.New("list is empty")

// A list of integers.
type IntList interface {
	// Push value onto front of list.
	PushFront(val int)

	// Push value onto back of list.
	PushBack(val int)

	// Remove element from front of list.
	PopFront() (val int, err error)

	// Remove element from rear of list.
	PopBack() (val int, err error)

	// All IntLists are also Stringers.
	String() string
}
