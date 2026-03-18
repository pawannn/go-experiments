package request

import (
	"fmt"
	"io"

	"github.com/pawannn/httpfromscratch/internal/header"
)

// GET /coffee HTTP/1.1\r\n
// Host: localhost:42069\r\n
// User-Agent: curl/7.81.0\r\n
// Accept: */*\r\n\r\n

type Request struct {
	RequestLine RequestLine
	Header      *header.Header
	Body        string

	state parserState
}

func newRequest() *Request {
	return &Request{
		Header: header.NewHeader(),
		Body:   "",

		state: stateInit,
	}
}

func (r *Request) done() bool {
	return r.state == stateDone
}

func (r *Request) errorState() bool {
	return r.state == stateError
}

func (r *Request) parse(data []byte) (int, error) {
	read := 0
OUTER:
	for {
		curentData := data[read:]

		switch r.state {
		case stateError:
			return 0, errStateRequest

		case stateInit:
			rl, n, err := parseRequestLine(curentData)
			if err != nil {
				r.state = stateError
				return 0, err
			}
			if n == 0 {
				break OUTER
			}

			r.RequestLine = *rl
			read += n
			r.state = stateHeader

		case stateHeader:
			n, done, err := r.Header.Parse(curentData)
			if err != nil {
				r.state = stateError
				return 0, err
			}

			if n == 0 {
				break OUTER
			}

			read += n

			if done {
				r.state = stateBody
			}

		case stateBody:
			contentLength := r.Header.GetContentLength()
			if contentLength == 0 {
				r.state = stateDone
			}

			toRead := contentLength - len(r.Body)
			available := len(curentData)

			n := min(toRead, available)

			if n > 0 {
				fmt.Println(string(curentData[:n]))
				r.Body += string(curentData[:n])
				read += n
			}

			if len(r.Body) == contentLength {
				r.state = stateDone
			}

		case stateDone:
			break OUTER
		}
	}
	return read, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	request := newRequest()

	buf := make([]byte, 1024)
	bufLen := 0

	for !request.done() && !request.errorState() {
		n, err := reader.Read(buf[bufLen:])
		if err != nil && err != io.EOF {
			return nil, err
		}

		bufLen += n
		readN, err := request.parse(buf[:bufLen])
		if err != nil {
			return nil, err
		}

		copy(buf, buf[readN:bufLen])
		bufLen -= readN

		if err == io.EOF {
			break
		}
	}

	return request, nil
}
