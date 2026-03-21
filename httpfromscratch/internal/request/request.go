package request

import (
	"httpfromscratch/internal/header"
	"io"
)

type Request struct {
	RequestLine RequestLine
	Headers     header.Headers
	Body        string

	parserState parserState
}

func newRequest() *Request {
	return &Request{
		parserState: stateInit,
		Headers:     header.NewHeader(),
		Body:        "",
	}
}

func (r *Request) done() bool {
	return r.parserState == stateDone
}

func (r *Request) error() bool {
	return r.parserState == stateError
}

func (r *Request) setState(parserState parserState) {
	r.parserState = parserState
}

func (r *Request) parseRequest(data []byte) (int, error) {
	read := 0

OUTER:
	for {
		currentdata := data[read:]

		switch r.parserState {
		case stateInit:
			rl, n, err := parseRequestLine(currentdata)
			if err != nil {
				r.setState(stateError)
				break
			}

			if n == 0 {
				break OUTER
			}

			r.RequestLine = *rl
			read += n

			r.setState(stateHeaders)

		case stateHeaders:
			n, done, err := r.Headers.ParseHeaders(currentdata)
			if err != nil {
				r.setState(stateError)
				break
			}

			if n == 0 {
				break OUTER
			}

			read += n
			if done {
				r.setState(stateBody)
			}

		case stateBody:
			contentLength := r.Headers.ContentLength()
			if contentLength == 0 {
				r.setState(stateDone)
				break
			}

			toRead := contentLength - len(r.Body)
			available := len(currentdata)

			n := min(toRead, available)

			if n > 0 {
				r.Body += string(currentdata)
				read += n
			}

			if len(r.Body) == contentLength {
				r.setState(stateDone)
			}

		case stateDone:
			break OUTER

		case stateError:
			return -1, errParserState

		default:
			panic("invalid parser state")
		}
	}

	return read, nil
}

func ParseFromReader(r io.Reader) (*Request, error) {
	request := newRequest()
	buf := make([]byte, 1024)
	bufLen := 0

	for !request.done() && !request.error() {
		n, err := r.Read(buf[bufLen:])
		if err != nil || err == io.EOF {
			return nil, err
		}

		bufLen += n

		readN, err := request.parseRequest(buf)
		if err != nil {
			return nil, err
		}

		bufLen -= readN
		copy(buf, buf[:bufLen])
	}

	return request, nil
}
