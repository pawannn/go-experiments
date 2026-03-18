package server

import (
	"fmt"
	"io"
	"net"
)

type Server struct {
	port   uint16
	closed bool
}

func Serve(port uint16) (*Server, error) {
	s := &Server{
		port:   port,
		closed: false,
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	go s.createServer(listener)

	return s, nil
}

func (s *Server) Close() error {
	s.closed = true
	return nil
}

func (s *Server) createServer(listener net.Listener) {
	go func() {
		for {
			conn, err := listener.Accept()

			if s.closed {
				return
			}

			if err != nil {
				return
			}

			go s.runConnection(conn)
		}
	}()
}
func (s *Server) runConnection(conn io.ReadWriteCloser) {
	body := "HELO WORLD"

	data := fmt.Sprintf(
		"HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s",
		len(body),
		body,
	)

	conn.Write([]byte(data))
	conn.Close()
}
