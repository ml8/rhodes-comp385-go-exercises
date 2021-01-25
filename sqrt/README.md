# Part 1: Control Structures and Functions

This exercise is taken directly from the [tour of
Go](https://tour.golang.org/flowcontrol/8).

Your work for this section is to implement a `sqrt` function in `main.go`. A
unit test has been written for this function in `sqrt_test.go`.

Your implementation of `sqrt` should make this test pass.

Run the program from the terminal with `go run .` when in this directory.

You can run the test from the terminal with `go test .` when in this directory.

In GoLand, right click on the `sqrt` folder and choose either "go build sqrt" or
"go test sqrt" from the Run menu.

## Submitting

When you are done with this exercise, __commit and push your working branch__ to
GitHub. __Continue working in the same branch for the rest of the exercises.__

## Tips

* You should look at how `sqrt` is (both in `main.go` and `sqrt_test.go`) used
  to determine what values it should return. 
* Errors are presented in Head First Go chapter 12.
* Notice the use of `Printf` and `Errorf` in both files--these are print/logging
  functions that use [format strings](https://golang.org/pkg/fmt/). If you
  learned C++ in 142/241 or C in 330, you will have probably seen these before.
  If not, read the documentation of the `fmt` package. `fmt.Sprintf` is a useful
  function in a lot of places. The logging library we will use later in this
  project and in other projects uses format strings, as well.
