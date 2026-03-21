package server

import (
	"fmt"
	"httpfromscratch/internal/request"
	"httpfromscratch/internal/response"
	"io"
	"net"
)

type Handler func(*response.ResponseWriter, *request.Request)

type Server struct {
	port     uint16
	open     bool
	listener net.Listener
	handler  Handler
}

func Serve(port uint16, handler Handler) (*Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	s := &Server{
		port:     port,
		open:     true,
		listener: listener,
		handler:  handler,
	}

	go s.startListener()

	return s, err
}

func (s *Server) Close() {
	s.open = false
}

func (s *Server) startListener() {
	go func() {
		for {
			conn, err := s.listener.Accept()
			if err != nil || !s.open {
				return
			}

			go s.handleConnection(conn)
		}
	}()
}

func (s *Server) handleConnection(conn io.ReadWriter) {
	responseWriter := response.NewResponseWriter(conn)

	req, err := request.ParseFromReader(conn)
	if err != nil {
		responseWriter.SendResponse(response.StatusBadRequest, []byte("Invalid request"))
		return
	}

	go s.handler(responseWriter, req)
}
