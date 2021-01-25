// Command-line client for chat server.
package main

import (
	"github.com/golang/glog"

	"go-exercises/webchat/chat"

	"flag"
	"fmt"
	"net/rpc"
)

// Send an RPC with the given message.
func send(client *rpc.Client, from string, to string, message string) {
	if from == "" || message == "" {
		return
	}

	var err error

	// Send Server.Put RPC.
	// TODO: Use client to send an RPC.

	if err != nil {
		glog.Errorf("Error sending: %v", err)
		return
	}
	fmt.Println("ok")
}

// Get all messages for the specified user.
func get(client *rpc.Client, who string) {
	if who == "" {
		return
	}

	var err error

	// Send Server.Get RPC.
	// TODO: Use client to make a Get RPC.

	if err != nil {
		glog.Errorf("Error getting: %v", err)
		return
	}
	// TODO: Print retrieved messages.
	// for _, msg := range resp.Msgs {
	// 	fmt.Println(msg)
	// }
}

func main() {
	var g bool
	var who string
	var to string
	var msg string
	var srv string

	flag.BoolVar(&g, "g", false, "Whether to get messages or send, false by default")
	flag.StringVar(&srv, "server", "127.0.0.1:5000", "server ip:port")
	flag.StringVar(&who, "me", "", "Identity of caller/sender")
	flag.StringVar(&to, "to", "", "Identity of receiver, only relevant when sending")
	flag.StringVar(&msg, "m", "", "Message to send, only relevant when sending")
	flag.Parse()

	// Connect to RPC server.
	client, err := rpc.DialHTTP("tcp", srv)
	if err != nil {
		glog.Fatalf("Error connecting to %s: %v", srv, err)
	}

	if g {
		// Retrieve message.
		get(client, who)
	} else {
		// Send message.
		send(client, who, to, msg)
	}
}
