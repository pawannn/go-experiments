package response

import (
	"fmt"
	"io"

	"github.com/pawannn/httpfromscratch/internal/header"
	"github.com/pawannn/httpfromscratch/internal/request"
)

type ResponseWriter struct {
	Writer io.Writer
}

func NewResponseWriter(w io.Writer) *ResponseWriter {
	return &ResponseWriter{
		Writer: w,
	}
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

type Handler func(w *ResponseWriter, req *request.Request)

func (w *ResponseWriter) WriteStatusLine(statusCode statusCode) error {
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

	_, err := w.Writer.Write(statusLine)
	return err
}

func (w *ResponseWriter) GetDefaultHeaders(contentLength int) *header.Header {
	h := header.NewHeader()
	h.Set("Content-Length", fmt.Sprintf("%d", contentLength))
	h.Set("Connection", "Closed")
	h.Set("Content-type", "text/plaintext")

	return h
}

func (w *ResponseWriter) WriteHeaders(h *header.Header) error {
	b := []byte{}

	h.Foreach(func(k, v string) {
		b = fmt.Appendf(b, "%s: %s\r\n", k, v)
	})
	b = fmt.Append(b, "\r\n")

	_, err := w.Writer.Write(b)
	return err
}

func (w *ResponseWriter) SendResponse(statusCode statusCode, data []byte) {
	length := len(data)
	headers := w.GetDefaultHeaders(length)
	w.WriteHeaders(headers)
	w.WriteStatusLine(statusCode)
	w.Writer.Write(data)
}
