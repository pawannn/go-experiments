package request

import "bytes"

type RequestLine struct {
	HTTPVersion   string
	Method        string
	RequestTarget string
}

func parseRequestLine(request_line []byte) (*RequestLine, int, error) {
	first_line, _, ok := bytes.Cut(request_line, rn)
	if !ok {
		return nil, 0, nil
	}

	parts := bytes.Split(first_line, []byte(" "))
	if len(parts) != 3 {
		return nil, -1, errMalfunctionedRequest
	}

	version, err := parseHttpVersion(parts[2])
	if err != nil {
		return nil, -1, err
	}

	if !validMethod(parts[0]) {
		return nil, -1, errInvalidMethod
	}

	return &RequestLine{
		HTTPVersion:   version,
		Method:        string(parts[0]),
		RequestTarget: string(parts[1]),
	}, len(first_line) + len(rn), nil
}

func parseHttpVersion(protocol []byte) (string, error) {
	p, version, ok := bytes.Cut(protocol, []byte("/"))
	if !ok {
		return "", errUnsupportedProtocol
	}

	if string(p) != "HTTP" {
		return "", errUnsupportedProtocol
	}

	return string(version), nil
}

func validMethod(method []byte) bool {
	switch string(method) {
	case "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS":
		return true
	}

	return false
}
