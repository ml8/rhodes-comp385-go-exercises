# Exercise 5: Background threading

# Session management

A [session](https://en.wikipedia.org/wiki/Session_(computer_science)) is a
period of interaction between a client and a server. Frequently, the use of
session refers to a persistent store associated with a user's interaction with a
web server, mostly often in the form of a key-value store. This way, a server
can maintain different elements of state for a particular user while they
interact with the server.

In `main.go` there is a `SessionManager` struct that could be used by a
webserver to manage user sessions. An individual `Session` is a map from keys to
values and a timestamp indicating when the `Session` was last updated.

Given that sessions are kept in memory, we want to limit the number of sessions
that are stored. It doesn't make sense, for example, for a webserver to maintain
a session for a user that has been idle for a year, ... or a month, ... or even
a day.

## Implement a session cleaner

For this exercise, you should implement a "background thread" (a thread that is
tasked with secondary maintainence or computation) that periodically cleans
sessions that have been idle for more than 5 seconds (of course, in reality,
sessions are much longer).

You will need to implement `CleanupThread` to remove idle sessions. You need to
ensure that this function attempts to clean idle sessions every second, and that
this function is started in a thread when a `SessionManager` is created.

In order to do this, you will need to pull together ideas from previous exercies
(including regular periodic actions and preventing concurrent access to shared
state).

## Hint

* You used `time.Ticker` in exercise 1.
* You used `go ...` to start threads in exercise 2.
* You used `sync.Mutex` in exercise 3.
* You used time calculations in exercise 4.

## Test your solution

To complete this exercise, you must pass two test. The normal `go
test` test cases as well as the race condition test.

Use the following commands to test your solution:
```
go test
go test -race
```
