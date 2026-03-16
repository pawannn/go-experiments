package internal

import (
	"bytes"
	"strings"
)

type RequestLine struct {
	HTTPVersion   string
	RequestTarget string
	Method        string
}

// GET /coffee HTTP/1.1\r\n
func ParseRequestLine(b []byte) (*RequestLine, int, error) {

	idx := bytes.Index(b, SEPARATOR)
	if idx == -1 {
		return nil, 0, nil
	}

	start_line := b[:idx]
	restOfMessage := idx + len(SEPARATOR)

	parts := bytes.Split(start_line, []byte(" "))
	if len(parts) != 3 {
		return nil, idx, ERR_MALFORMED_REQ_LINE
	}

	version, err := GetHTTPVersion(string(parts[2]))
	if err != nil {
		return nil, idx, err
	}

	rl := &RequestLine{
		Method:        string(parts[0]),
		RequestTarget: string(parts[1]),
		HTTPVersion:   version,
	}

	return rl, restOfMessage, nil
}

func GetHTTPVersion(httpVersion string) (string, error) {
	if httpVersion != "HTTP/1.1" {
		return "", ERR_UNSUPPORTED_HTTP_VERSION
	}

	_, version, _ := strings.Cut(httpVersion, "/")
	return version, nil
}
