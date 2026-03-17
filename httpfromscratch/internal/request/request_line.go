package request

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
func parseRequestLine(b []byte) (*RequestLine, int, error) {

	idx := bytes.Index(b, SEPARATOR)
	if idx == -1 {
		return nil, 0, nil
	}

	start_line := b[:idx]
	restOfMessage := idx + len(SEPARATOR)

	parts := bytes.Split(start_line, []byte(" "))
	if len(parts) != 3 {
		return nil, 0, errMalfunctionedLine
	}

	version, err := getHTTPVersion(string(parts[2]))
	if err != nil {
		return nil, 0, err
	}

	rl := &RequestLine{
		Method:        string(parts[0]),
		RequestTarget: string(parts[1]),
		HTTPVersion:   version,
	}

	return rl, restOfMessage, nil
}

func getHTTPVersion(httpVersion string) (string, error) {
	if httpVersion != "HTTP/1.1" {
		return "", errUnsupportedHTTPVersion
	}

	_, version, _ := strings.Cut(httpVersion, "/")
	return version, nil
}
