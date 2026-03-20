package response

import (
	"fmt"
	"io"

	"github.com/pawannn/httpfromscratch/internal/header"
	"github.com/pawannn/httpfromscratch/internal/request"
)

type Response struct {
}

type statusCode int

const (
	StatusOk                  statusCode = 200
	StatusBadRequest          statusCode = 400
	StatusInternalServerError statusCode = 500
)

type HandlerError struct {
	Code    statusCode
	Message string
}

type Handler func(w io.Writer, req *request.Request) *HandlerError

func WriteStatusLine(w io.Writer, statusCode statusCode) error {
	statusLine := []byte{}
	switch statusCode {
	case StatusOk:
		statusLine = []byte("HTTP/1.1 200 OK\r\n")
	case StatusBadRequest:
		statusLine = []byte("HTTP/1.1 400 Bad Request\r\n")
	case StatusInternalServerError:
		statusLine = []byte("HTTP/1.1 500 Internal Server Error\r\n")
	default:
		return fmt.Errorf("unrecognized error code")
	}

	_, err := w.Write(statusLine)
	return err
}

func GetDefaultHeaders(contentLength int) *header.Header {
	h := header.NewHeader()
	h.Set("Content-Length", fmt.Sprintf("%d", contentLength))
	h.Set("Connection", "Closed")
	h.Set("Content-type", "text/plaintext")

	return h
}

func WriteHeaders(w io.Writer, h *header.Header) error {
	b := []byte{}

	h.Foreach(func(k, v string) {
		b = fmt.Appendf(b, "%s: %s\r\n", k, v)
	})
	b = fmt.Append(b, "\r\n")

	_, err := w.Write(b)
	return err
}
