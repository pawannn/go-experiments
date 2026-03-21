package response

import (
	"fmt"
	"httpfromscratch/internal/header"
	"io"
)

type StatsCode int

const (
	StatusOk                  StatsCode = 200
	StatusBadRequest          StatsCode = 400
	StatusInternalServerError StatsCode = 500
)

type ResponseWriter struct {
	w io.Writer
}

func NewResponseWriter(conn io.Writer) *ResponseWriter {
	return &ResponseWriter{
		w: conn,
	}
}

func (rw *ResponseWriter) writeStatusLine(statsCode StatsCode) error {
	b := []byte{}

	switch statsCode {
	case StatusOk:
		b = []byte("HTTP/1.1 200 OK\r\n")
	case StatusBadRequest:
		b = []byte("HTTP/1.1 400 Bad Request\r\n")
	case StatusInternalServerError:
		b = []byte("HTTP/1.1 500 Internal Server Error\r\n")
	}

	_, err := rw.w.Write(b)
	return err
}

func (rw *ResponseWriter) getDefaultHeaders(contentLength int) header.Headers {
	h := header.NewHeader()

	h.SetContentLength(contentLength)
	h.Set("Connection", "Close")
	h.Set("Content-Type", "text/plaintext")

	return h
}

func (rw *ResponseWriter) writeHeaders(header header.Headers) error {
	b := []byte{}

	header.ForEach(func(k, v string) {
		b = fmt.Appendf(b, "%s: %s\r\n", k, v)
	})
	b = fmt.Append(b, "\r\n")

	_, err := rw.w.Write(b)
	return err
}

func (rw *ResponseWriter) SendResponse(statusCode StatsCode, body []byte) {
	contentLength := len(body)
	defaultHeader := rw.getDefaultHeaders(contentLength)

	rw.writeStatusLine(statusCode)
	rw.writeHeaders(defaultHeader)

	rw.w.Write(body)
}
