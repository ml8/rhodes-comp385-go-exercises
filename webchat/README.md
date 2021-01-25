# Part 4: Using RPCs

In this final exercise, you will use Go RPCs to build a web-based multi-user
chat application.

## Exercise structure

The structure of the project looks like this:

```
                 index.html
                 +-------------------------+
                 |                         |
                 |      Web frontend       |
                 |   (javascript + html)   |
                 |                         |
                 +-----------+-------------+
                             |
                             |    http
                             |  requests
                             |
      cmd/webchat/webchat.go v
                 +-----------+-------------+
                 |                         |
                 |     Web server with     |  cmd/chat/chat.go
                 |  index, /get and /send  |    +----------------+
                 |    (Go HTTP server)     |    |                |
                 |                         |    |  Command line  |
                 +-----------+-------------+    |   chat tool    |
                             |                  |  (Go Program)  |
                             |     Go           |                |
                             |    RPCs          +-------+--------+
                chat/chat.go |                          |
cmd/chatserver/chatserver.go v                          |   Go
                 +-----------+-------------+            |  RPCs
                 |                         |            |
                 |       Chat server       +<-----------+
                 |     (Go RPC server)     |
                 |                         |
                 +-------------------------+

```

* The web frontend displays user messages and allows users to type a new message
  in the UI.
* When the web frontend needs to get new messages for a user or send a message,
  it makes HTTP requests to the web server, which knows how to handle these
  requests.
* The web server, when receiving a request to get or send messages, makes an RPC
  to an RPC server that exposes the chat object in `chat/chat.go` so that it
  can be interacted with via RPC.
* There is also a separate command line utility that can be used to send and
  receive messages.

Most of this project has been supplied for you. There are only a few things that
you will be required to do.

1. The chat object defined in `chat/chat.go` is not thread-safe and will
   have errors if users attempt to make concurrent calls to its methods. You
   need to make sure there cannot be concurrent modifications to the
   `chat.Server`'s state (`q`). 
1. Modify `cmd/chat/chat.go` to actually make the RPCs to the chat server.
1. Modify `cmd/webcat/webchat.go` to actually make the RPCs to the chat server.

When you are ready submit, commit your changes and push them to GitHub.

Since this is the last section of the lab, create a pull request and set me as
the reviewer. This will be how I grade your lab.

## Tips

* Before you begin, take some time to read through the code. What is the purpose
  of each file? What is each function doing? Inspect the source of `index.html`
  -- how is this file using the web server?
* Each part of this section only requires a few lines of code. 
* Once you have completed part 2, you should have all the code required for part
  3.
* Since this project requires running several processes, the __easiest__ way to
  test everything end to end is to use the command line. If you are using
  GoLand, you can open a terminal by clicking the "Terminal" button along the
  bottom of the window. You can create multiple terminals by clicking the plus
  button "new session" at the top of the terminal window. From this terminal,
  you can `cd` to the correct directory:

  ```
  cd webchat
  ```

  Now the commands in the following section will work.
* To test your part 2, first run the chat server:

  ```
  go run ./cmd/chatserver -logtostderr
  ```

  Then, in a separate window, try sending and receiving messages:

  ```
  # send a DM: matt -> elise
  go run ./cmd/chat -logtostderr -me matt -to elise -m "hey! hows it going?"

  # send a global message: matt -> all
  go run ./cmd/chat -logtostderr -me matt -m "hey everyone"

  # send a DM: elise -> matt
  go run ./cmd/chat -logtostderr -me elise -m "hey"
  
  # get matt's messages (one global
  go run ./cmd/chat -logtostderr -me matt -g

  # get elise's messages (should be one global, one direct)
  go run ./cmd/chat -logtostderr -me elise -g
  ```
* Once you have this working, you should be able to basically reuse the same
  code in `cmd/webchat/webchat.go`.

  Then you can test the full web chat app.

  If the chat server is not already running, run it.

  ```
  go run ./cmd/chatserver -logtostderr
  ```

  Then, run the web server:

  ```
  go run ./cmd/webchat -logtostderr
  ```

  When the web server is running, you should be able to connect to it by
  visiting [http://localhost:8080](http://localhost:8080) in any browser.

  The UI is very simple. To start chatting, make sure the "My id" field is
  filled, and then you can use the message field to send messages (optionally
  DMs by using the "To" field).

  Open another browser window and be another user, or find your IP address with
  `ifconfig | grep inet` and look for the interface with four octets (and is not
  127.0.0.1 -- e.g., 10.10.47.67) and ask a friend to open a browser to
  `http://<your ip address>:8080`.

  Note that the way that the frontend works is by polling for new messages a few
  times a second. This means that if you are logging to standard error, you will
  see quite a big of logs (perhaps these log messages are candidates for
  verbosity levels...).

  ![Screenshot of
  interaction](https://storage.googleapis.com/distributed-files/screenshot.png)
