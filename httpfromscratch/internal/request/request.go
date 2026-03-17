package request

import (
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
	state       parserState
}

func newRequest() *Request {
	return &Request{
		state:  stateInit,
		Header: header.NewHeader(),
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
		if err != nil {
			return nil, err
		}

		bufLen += n
		readN, err := request.parse(buf[:bufLen])
		if err != nil {
			return nil, err
		}

		copy(buf, buf[readN:bufLen])
		bufLen -= readN
	}

	return request, nil
}
