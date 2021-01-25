// RPC server for chat.Server.

package main

import (
	"github.com/golang/glog"

	"go-exercises/webchat/chat"

	"flag"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	var port string
	flag.StringVar(&port, "port", ":5000", "port to listen on")
	flag.Parse()

	// Create chat.Server object.
	srv := chat.New()

	// Register server as an RPC receiver.
	_ = rpc.Register(srv)

	// Register an HTTP handler for incoming RPCs.
	rpc.HandleHTTP()

	// Listen and serve RPCs.
	l, err := net.Listen("tcp", port)
	if err != nil {
		glog.Fatalf("Error listening on %s: %v", "5000", err)
	}
	glog.Infof("Listening on port %s", port)
	glog.Fatal(http.Serve(l, nil))
}
