// Web server for web-based chat app.

package main

import (
	"github.com/golang/glog"

	"go-exercises/webchat/chat"

	"flag"
	"io/ioutil"
	"net/http"
	"net/rpc"
	"strings"
)

var client *rpc.Client // RPC client to chat server
var index []byte       // Index page

// Calls RPC on server to retrieve messages.
func getMessages(client *rpc.Client, who string) []string {

	var err error

	// Make Server.Get RPC
	// TODO: Make RPC using client.

	if err != nil {
		glog.Errorf("Error getting: %v", err)
		return nil
	}
	// TODO: Return retrieved messages.
	return nil
}

// Retrieves messages for a user using URL parameters.
// /get/username retrieves messages for the user named username.
func get(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Path[len("/get/"):]
	glog.Infof("get for %s %s", r.URL.Path, user)
	if user == "" {
		// Ignore empty usernames.
		return
	}
	msgs := getMessages(client, user)
	// Format as HTML with line breaks.
	var sb strings.Builder
	for _, msg := range msgs {
		sb.WriteString(msg)
		sb.WriteString("<br/>")
	}
	// Write result to client.
	_, _ = w.Write([]byte(sb.String()))
}

// Sends a message for a user using a form post request.
// from, to, and msg values are read from the form.
func send(_ http.ResponseWriter, r *http.Request) {
	from := r.FormValue("from")
	to := r.FormValue("to")
	msg := r.FormValue("msg")
	if from == "" || msg == "" {
		return
	}
	glog.Infof("send %v %v %v", from, to, msg)

	var err error

	// Make Put RPC.
	// TODO: Make RPC using client.

	if err != nil {
		glog.Errorf("Error sending: %v", err)
	}
	glog.Infof("Sent %v -> %v: %v", from, to, msg)
}

// Handler for / (index).
func handler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("Get for %v", r.URL.Path)
	_, _ = w.Write(index)
}

func main() {
	var srv string
	var idx string
	flag.StringVar(&srv, "server", "127.0.0.1:5000", "server ip:port")
	flag.StringVar(&idx, "indexPage", "index.html", "index page")
	flag.Parse()

	// Read index page.
	var err error
	index, err = ioutil.ReadFile(idx)
	if err != nil {
		glog.Fatalf("Error reading %s: %v", idx, err)
	}

	// Connect to chat server.
	client, err = rpc.DialHTTP("tcp", srv)
	if err != nil {
		glog.Fatalf("Error connecting to %s: %v", srv, err)
	}

	// Requests for the root retrieve the index page.
	http.HandleFunc("/", handler)
	// Requests for /get should retrieve messages.
	http.HandleFunc("/get/", get)
	// Requests for /send should send messages.
	http.HandleFunc("/send/", send)

	glog.Info("Listening on port :8080")
	// Start HTTP server.
	glog.Fatal(http.ListenAndServe(":8080", nil))
}
