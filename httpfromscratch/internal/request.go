package internal

import (
	"io"
)

// GET /coffee HTTP/1.1\r\n
// Host: localhost:42069\r\n
// User-Agent: curl/7.81.0\r\n
// Accept: */*\r\n\r\n

type Request struct {
	RequestLine RequestLine
	State       parserState
}

func NewRequest() *Request {
	return &Request{
		State: stateInit,
	}
}

func (r *Request) Done() bool {
	return r.State == stateDone
}

func (r *Request) ErrorState() bool {
	return r.State == stateError
}

func (r *Request) Parse(data []byte) (int, error) {
	read := 0
OUTER:
	for {
		switch r.State {
		case stateError:
			return 0, ERR_REQUEST_ERRSTATE
		case stateInit:
			rl, n, err := ParseRequestLine(data)
			if err != nil {
				r.State = stateError
				return 0, err
			}
			if n == 0 {
				break OUTER
			}

			r.RequestLine = *rl
			read += n
			r.State = stateDone
		case stateDone:
			break OUTER
		}
	}
	return read, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	request := NewRequest()

	buf := make([]byte, 1024)
	bufLen := 0

	for !request.Done() && !request.ErrorState() {
		n, err := reader.Read(buf[bufLen:])
		if err != nil {
			return nil, err
		}

		bufLen += n
		readN, err := request.Parse(buf[:bufLen])
		if err != nil {
			return nil, err
		}

		copy(buf, buf[readN:bufLen])
		bufLen -= readN
	}

	return request, nil
}
