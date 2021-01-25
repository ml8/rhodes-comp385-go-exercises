// Multi-user chat server.
// chat.Server can be used as an RPC server.
package chat

import (
	"github.com/golang/glog"

	"fmt"
)

// The server, maintaining a queue of messages per user.
type Server struct {
	q map[string][]string // Map user name -> array of messages
}

// Put/send message request. If To is empty, the message is to everyone.
type PutRequest struct {
	From string // Sender
	To   string // Receiver, empty for all users
	Msg  string // Message to send
}

// Put/send response
type PutResponse struct {
}

// Get messages request.
type GetRequest struct {
	For string // User to retrieve messages for
}

// Get messages response.
type GetResponse struct {
	Msgs []string // Received messages, in order
}

// Create a new chat server
func New() *Server {
	return &Server{q: make(map[string][]string)}
}

// Put a message.
func (s *Server) Put(req *PutRequest, _ *PutResponse) error {
	glog.Infof("Put for %s from %s: %s", req.To, req.From, req.Msg)

	// Empty To -> message is for everyone
	if req.To == "" {
		// If sender is not in message queue map, add them.
		if _, ok := s.q[req.From]; !ok {
			s.q[req.From] = nil
		}
		// Send message to everyone.
		for user := range s.q {
			s.q[user] = append(s.q[user], fmt.Sprintf("%s: %s", req.From, req.Msg))
		}
	} else {
		// Private message
		s.q[req.To] = append(s.q[req.To], fmt.Sprintf("-> %s: %s", req.From, req.Msg))
	}
	return nil
}

func (s *Server) Get(req *GetRequest, resp *GetResponse) error {
	resp.Msgs = s.q[req.For]
	// Note: not the same as delete(s.q, req.For) -- this will remove the key from
	// the map (rather than leave the key in the map with a nil slice), which
	// causes the user to no longer receive global messages.
	s.q[req.For] = nil
	glog.Infof("Get for %s: %d messages", req.For, len(resp.Msgs))
	return nil
}
