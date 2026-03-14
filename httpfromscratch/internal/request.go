package internal

import (
	"errors"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HTTPVersion   string
	RequestTarget string
	Method        string
}

func GetHTTPVersion(httpVersion string) (string, error) {
	if httpVersion != "HTTP/1.1" {
		return "", ERR_UNSUPPORTED_HTTP_VERSION
	}

	_, version, _ := strings.Cut(httpVersion, "/")
	return version, nil
}

func ParseRequestLine(b string) (*RequestLine, string, error) {
	before, after, ok := strings.Cut(b, SEPARATOR)
	if !ok {
		return nil, b, nil
	}

	start_line := before
	restOfMessage := after

	parts := strings.Split(start_line, " ")
	if len(parts) != 3 {
		return nil, restOfMessage, ERR_MALFORMED_REQ_LINE
	}

	version, err := GetHTTPVersion(parts[2])
	if err != nil {
		return nil, restOfMessage, err
	}

	rl := &RequestLine{
		Method:        parts[0],
		RequestTarget: parts[1],
		HTTPVersion:   version,
	}

	return rl, restOfMessage, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.New("unable to readall" + err.Error())
	}

	rl, _, err := ParseRequestLine(string(data))
	if err != nil {
		return nil, err
	}

	return &Request{
		RequestLine: *rl,
	}, err
}
