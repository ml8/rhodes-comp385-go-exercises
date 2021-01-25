package bench

import (
	"go-exercises/modules/list"

	"math/rand"
	"testing"
)

// Build a list, randomly pushing to front or back.
func buildList(lst list.IntList, iters int) {
	for i := 0; i < iters; i++ {
		switch rand.Int() % 2 {
		case 0:
			lst.PushFront(rand.Int())
		case 1:
			lst.PushBack(rand.Int())
		}
	}
}

// Run a random pattern of push/pop operations.
func benchmarkRandom(lst list.IntList, iters int) {
	for i := 0; i < iters; i++ {
		switch rand.Int() % 4 {
		case 0:
			lst.PushFront(rand.Int())
		case 1:
			lst.PushBack(rand.Int())
		case 2:
			_, _ = lst.PopFront()
		case 3:
			_, _ = lst.PopBack()
		}
	}
}

// Repeatedly modify the back.
func benchmarkBack(lst list.IntList, iters int) {
	for i := 0; i < iters; i++ {
		switch rand.Int() % 2 {
		case 0:
			lst.PushBack(rand.Int())
		case 1:
			_, _ = lst.PopBack()
		}
	}
}

// Repeatedly modify the front.
func benchmarkFront(lst list.IntList, iters int) {
	for i := 0; i < iters; i++ {
		switch rand.Int() % 2 {
		case 0:
			lst.PushFront(rand.Int())
		case 1:
			_, _ = lst.PopFront()
		}
	}
}

func benchmarkList(lst list.IntList, f func(list.IntList, int), iters int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		f(lst, iters)
	}
}

func BenchmarkLinkedListRandom4Kb(b *testing.B) {
	lst := list.NewLinkedIntList()
	iters := 100_000
	buildList(lst, iters)
	b.ResetTimer()
	benchmarkList(lst, benchmarkRandom, iters, b)
}

// Must be run with a longer timeout to get multiple runs.
// e.g., -benchtime=20s
func BenchmarkArrayListRandom4Kb(b *testing.B) {
	lst := list.NewArrayIntList()
	iters := 100_000
	buildList(lst, iters)
	b.ResetTimer()
	benchmarkList(lst, benchmarkRandom, iters, b)
}

func BenchmarkLinkedListBack4Kb(b *testing.B) {
	lst := list.NewLinkedIntList()
	iters := 100_000
	buildList(lst, iters)
	b.ResetTimer()
	benchmarkList(lst, benchmarkBack, iters, b)
}

func BenchmarkArrayListBack4Kb(b *testing.B) {
	lst := list.NewArrayIntList()
	iters := 100_000
	buildList(lst, iters)
	b.ResetTimer()
	benchmarkList(lst, benchmarkBack, iters, b)
}

func BenchmarkLinkedListFront4Kb(b *testing.B) {
	lst := list.NewLinkedIntList()
	iters := 100_000
	buildList(lst, iters)
	b.ResetTimer()
	benchmarkList(lst, benchmarkFront, iters, b)
}

func BenchmarkArrayListFront4Kb(b *testing.B) {
	lst := list.NewArrayIntList()
	iters := 100_000
	buildList(lst, iters)
	b.ResetTimer()
	benchmarkList(lst, benchmarkFront, iters, b)
}
