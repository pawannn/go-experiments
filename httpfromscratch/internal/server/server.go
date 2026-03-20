package server

import (
	"bytes"
	"fmt"
	"io"
	"net"

	"github.com/pawannn/httpfromscratch/internal/request"
	"github.com/pawannn/httpfromscratch/internal/response"
)

type Server struct {
	port    uint16
	closed  bool
	handler response.Handler
}

func Serve(port uint16, handler response.Handler) (*Server, error) {
	s := &Server{
		port:    port,
		closed:  false,
		handler: handler,
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
	defer conn.Close()

	h := response.GetDefaultHeaders(0)
	r, err := request.RequestFromReader(conn)
	if err != nil {
		response.WriteStatusLine(conn, response.StatusBadRequest)
		response.WriteHeaders(conn, h)
		return
	}

	writer := bytes.NewBuffer([]byte{})
	handlerErr := s.handler(writer, r)

	statusCode := response.StatusOk
	body := []byte{}

	if handlerErr != nil {
		statusCode = handlerErr.Code
		body = []byte(handlerErr.Message)
	} else {
		body = writer.Bytes()
	}

	h.Replace("Content-Length", fmt.Sprintf("%d", len(body)))
	response.WriteStatusLine(conn, statusCode)
	response.WriteHeaders(conn, h)
	conn.Write(body)
}
