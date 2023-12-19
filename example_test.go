package options_test

import (
	"fmt"
	"time"
)

// example test guide of FunctionalOptions pattern in Go via https://sagikazarmark.hu/blog/functional-options-on-steroids/
func Example_basic() {
	type Server struct {
		addr string
	}

	// NewServer initializes a new Server listening on addr.
	NewServer := func(addr string) *Server {
		return &Server{
			addr: addr,
		}
	}

	// Now we can use NewServer to create a new Server.
	server := NewServer(":8080")
	fmt.Println(server.addr)
	// Output:
	// :8080
}

func Example_withTimeoutOption() {

	type Server struct {
		addr string

		// default: no timeout
		timeout time.Duration
	}

	// Timeout configures a maximum length of idle connection in Server.
	Timeout := func(timeout time.Duration) func(*Server) { // would normally be defined in options.go with something like: func Timeout(timeout time.Duration) func(*Server)
		return func(s *Server) {
			s.timeout = timeout
		}
	}

	// NewServer initializes a new Server listening on addr with optional configuration.
	NewServer := func(addr string, opts ...func(*Server)) *Server { // would normally be defined in options.go with something like: func NewServer(addr string, opts ...func(*Server)) *Server
		server := &Server{
			addr: addr,
		}

		// apply the list of options to Server
		for _, opt := range opts {
			opt(server)
		}

		return server
	}

	// Now we can use NewServer to create new servers with different timeouts.
	// no optional paramters, use defaults
	server := NewServer(":8080")

	// configure a timeout in addition to the address
	server = NewServer(":8080", Timeout(10*time.Second))
	fmt.Println(server.addr, server.timeout)
	// Output:	:8080 10s
}

func Example_optionType() {

	type Server struct {
		addr string

		// default: no timeout
		timeout time.Duration
	}

	type Option func(*Server)

	// Timeout configures a maximum length of idle connection in Server.
	Timeout := func(timeout time.Duration) func(*Server) { // would normally be defined in options.go with something like: func Timeout(timeout time.Duration) func(*Server)
		return func(s *Server) {
			s.timeout = timeout
		}
	}

	// NewServer initializes a new Server listening on addr with optional configuration.
	NewServer := func(addr string, opts ...Option) *Server { // would normally be defined in options.go with something like: func NewServer(addr string, opts ...func(*Server)) *Server
		server := &Server{
			addr: addr,

			// default values
			timeout: 0,

			// other default values

		}

		// apply the list of options to Server
		for _, opt := range opts {
			opt(server)
		}

		return server
	}

	// Now we can use NewServer to create new servers with different timeouts.
	// no optional paramters, use defaults
	server := NewServer(":8080")

	// configure a timeout in addition to the address
	server = NewServer(":8080", Timeout(10*time.Second))
	fmt.Println(server.addr, server.timeout)
	// Output:	:8080 10s
}
