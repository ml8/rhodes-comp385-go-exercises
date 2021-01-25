package list

import (
	"testing"
)

func TestEmptyList(t *testing.T) {
	ll := NewLinkedIntList()
	if ll.String() != "[]" {
		t.Errorf("Empty list is not []: %v", ll.String())
	}
}

func TestEmptyListReturnsEmptyErr(t *testing.T) {
	ll := NewLinkedIntList()
	if _, err := ll.PopFront(); err != EmptyListErr {
		t.Errorf("Incorrect error returned from PopFront: %v", err)
	}
	if _, err := ll.PopBack(); err != EmptyListErr {
		t.Errorf("Incorrect error returned from PopBack: %v", err)
	}
}

func TestBuildList(t *testing.T) {
	ll := NewLinkedIntList()
	expected := "[1, 2, 3, 4, 5]"
	ll.PushFront(3)
	ll.PushBack(4)
	ll.PushFront(2)
	ll.PushBack(5)
	ll.PushFront(1)

	if actual := ll.String(); actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestBuildAndRemove(t *testing.T) {
	ll := NewLinkedIntList()
	expected := []int{5, 1, 4, 3, 2}
	ll.PushBack(1)
	ll.PushBack(2)
	ll.PushBack(3)
	ll.PushBack(4)
	ll.PushBack(5)

	errs := make([]error, len(expected))
	actual := make([]int, len(expected))
	actual[0], errs[0] = ll.PopBack()
	actual[1], errs[1] = ll.PopFront()
	actual[2], errs[2] = ll.PopBack()
	actual[3], errs[3] = ll.PopBack()
	actual[4], errs[4] = ll.PopBack()

	for i, v := range expected {
		if errs[i] != nil {
			t.Errorf("Unexpected error at %d: %v", i, errs[i])
		}
		if actual[i] != v {
			t.Errorf("Got %d, expected %d (at %d)", actual[i], v, i)
		}
	}
}

func TestPushBackPopFront(t *testing.T) {
	ll := &linkedIntList{}
	expected := 42
	ll.PushBack(expected)
	actual, err := ll.PopFront()
	if actual != expected {
		t.Errorf("Incorrect value popped, expected %d, got %d", expected, actual)
	}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if ll.Head != nil || ll.Tail != nil {
		t.Errorf("Leaked pointer. Head: %v, Tail: %v", ll.Head, ll.Tail)
	}
	if ll.String() != "[]" {
		t.Errorf("Empty list is not []: %v", ll.String())
	}
}

func TestPushBackPopBack(t *testing.T) {
	ll := &linkedIntList{}
	expected := 42
	ll.PushBack(expected)
	actual, err := ll.PopBack()
	if actual != expected {
		t.Errorf("Incorrect value popped, expected %d, got %d", expected, actual)
	}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if ll.Head != nil || ll.Tail != nil {
		t.Errorf("Leaked pointer. Head: %v, Tail: %v", ll.Head, ll.Tail)
	}
	if ll.String() != "[]" {
		t.Errorf("Empty list is not []: %v", ll.String())
	}
}

func TestPushFrontPopFront(t *testing.T) {
	ll := &linkedIntList{}
	expected := 42
	ll.PushFront(expected)
	actual, err := ll.PopFront()
	if actual != expected {
		t.Errorf("Incorrect value popped, expected %d, got %d", expected, actual)
	}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if ll.Head != nil || ll.Tail != nil {
		t.Errorf("Leaked pointer. Head: %v, Tail: %v", ll.Head, ll.Tail)
	}
	if ll.String() != "[]" {
		t.Errorf("Empty list is not []: %v", ll.String())
	}
}

func TestPushFrontPopBack(t *testing.T) {
	ll := &linkedIntList{}
	expected := 42
	ll.PushFront(expected)
	actual, err := ll.PopBack()
	if actual != expected {
		t.Errorf("Incorrect value popped, expected %d, got %d", expected, actual)
	}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if ll.Head != nil || ll.Tail != nil {
		t.Errorf("Leaked pointer. Head: %v, Tail: %v", ll.Head, ll.Tail)
	}
	if ll.String() != "[]" {
		t.Errorf("Empty list is not []: %v", ll.String())
	}
}
