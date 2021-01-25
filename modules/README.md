# Part 2: Modules and Interfaces

In this exercise, you will create a struct that implements an interface, and use
that struct in another package. 

### Packages 

A _package_ in Go is a group of related interfaces, structs, and functions.
For example, `fmt` is a package with formatting functions, `math` is a package
with math-related functions, etc.

The visibility of types and functions _outside_ of a package is controlled by
whether identifier names start with a capital or lower case letter. 

If you open any file in the `list` directory, you will see that they all start
with `package list`, identifying that the file belongs to the `list` package.

If you open the file `bench/benchmark_test.go`, you can see that it imports the
`list` package (the project is defined as `modules`, so the package is
`modules/list`).

### Interfaces

In Go, an _interface_ is similar to an interface in Java or an entirely abstract
class in C++ (i.e., every function is pure virtual).

This is how Go achieves subtype polymorphism -- whenever an interface type is
used, it can refer to any type that implements that interface.

To observe this, first open `list/list.go` -- this file contains the definition
of the interface `IntList`, representing a list of integers that can be modified
by pushing/popping values from either end of a list.

Any struct or type that implements the functions defined in the interface now
_implements_ that interface and has a "_is-a_" relationship to the interface.

For example, open `list/linked_list.go`. This file contains a complete
linked list implementation of `IntList`. The struct `linkedIntList` (notice the
lower case; it cannot be referenced outside of this package) has each of the
methods that are required by the `IntList` interface.

Finally, open `bench/benchmark_test.go`. This is a
[_benchmark_](https://en.wikipedia.org/wiki/Benchmark_(computing)) that will
compare the runtime of two list implementations. Notice how the parameters for
the `benchmarkXXX` methods all accept a parameter of type `list.IntList`. Now,
notice how the actual value supplied to that parameter can either be a linked
list (e.g., in `BenchmarkLinkedListRandom4Kb`) or an array list (e.g., in
`BenchmarkArrayListRandom4Kb`).

### Assignment

Your job is to implement an array-based implementation of the `IntList`
interface. Your code will be similar to `linked_list.go`, except you will be
using array slicing and `append` to build your list. 

Note that `linked_list.go` has a complete set of test in `linked_list_test.go`.
You should write tests for your implementation as well, in `array_list_test.go`.

### Testing and running the benchmark

If you are in this folder, you can run `go test ./list` to run the tests in both
`array_list_test.go` and `linked_list_test.go`.

In GoLand, right click on the `list` folder and choose "go test list" from the
Run menu.

You can run the benchmark comparing the implementations by running: 

```
cd bench
go test -bench=. -benchtime=20s
```

In GoLand, right click on the `bench` folder and choose "gobench bench" from the
Run menu.

Note that it is likely that your array-based implementation will likely be
dramatically worse in performance. In __this__ file, Fill in the following table
of runtimes and write brief summary of why the performance is worse, and 2 ideas
for how you could improve the runtime of your implementation.

| test                          | iterations of the benchmar  | time/iteration (ns)
|---                            |---                          |---
| BenchmarkLinkedListRandom4Kb  |                             |
| BenchmarkArrayListRandom4Kb   |                             |
| BenchmarkLinkedListBack4Kb    |                             | 
| BenchmarkArrayListBack4Kb     |                             |
| BenchmarkLinkedListFront4Kb   |                             |
| BenchmarkArrayListFront4Kb    |                             |


## Submitting

When you are done with this exercise, __commit and push your working branch__ to
GitHub. __Continue working in the same branch for the rest of the exercises.__

## Tips

* The relevant chapters in Head First Go are 5+6 (arrays and slices), 8
  (structs), 9-10 (methods), and 11 (interfaces).
* Make sure you use [pointer receiver
  methods](https://tour.golang.org/methods/4) for your implementation, as you
  will be modifying the struct's contents.
* Using `append` for these methods is the simplest approach. Note that if you
  want to [append an array to
  another](https://stackoverflow.com/questions/16248241/concatenate-two-slices-in-go),
  you need to use the `...` syntax: `append(one, other...)`. If you use this
  approach, you don't even need to use `make` once.
* A simple implementation is 1-2 lines of code per method (excluding error
  checking for an empty list in `PopXXX` methods), except for `String()`.
