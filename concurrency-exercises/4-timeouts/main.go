package main

import (
	"flag"
	"time"
)

// User of the service.
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  time.Duration
}

// HandleRequest runs the processes requested by users. Returns true iff the
// process completes.
func HandleRequest(process func(), u *User) bool {
	// TODO: time out and return false if process() is taking more than the user's
	// remaining time.
	now := time.Now()
	process()
	u.TimeUsed += (time.Now().Sub(now))

	return true
}

func main() {
	flag.Parse()
	RunMockServer()
}
