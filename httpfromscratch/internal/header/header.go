package header

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Header struct {
	Header map[string]string
}

func NewHeader() *Header {
	h := &Header{
		Header: make(map[string]string),
	}

	return h
}

func (h *Header) Get(key string) string {
	return h.Header[strings.ToLower(key)]
}

func (h *Header) Set(key string, value string) {
	name := strings.ToLower(key)
	if v, ok := h.Header[name]; ok {
		h.Header[name] = fmt.Sprintf("%s,%s", v, value)
	} else {
		h.Header[name] = value
	}
}

func (h *Header) GetContentLength() int {
	contentLength := h.Get("content-length")
	if contentLength == "" {
		return 0
	}

	length, err := strconv.Atoi(contentLength)
	if err != nil {
		return 0
	}

	return length
}

func (h *Header) Parse(data []byte) (int, bool, error) {
	read := 0
	done := false

	for {
		idx := bytes.Index(data, rn)
		if idx == -1 {
			break
		}

		if idx == 0 {
			read += len(rn)
			done = true
			break
		}

		name, value, err := parseHeader(data[:idx])
		if err != nil {
			return 0, false, err
		}

		if !isToken([]byte(name)) {
			return 0, false, errMalfunctionedHeaderLine
		}

		read += idx + len(rn)
		h.Set(name, value)

		data = data[idx+len(rn):]
	}

	return read, done, nil
}

func parseHeader(data []byte) (string, string, error) {
	if len(data) == 0 {
		return "", "", nil
	}

	parts := bytes.SplitN(data, []byte(":"), 2)
	if len(parts) != 2 {
		return "", "", errMalfunctionedHeaderLine
	}

	field := bytes.TrimSpace(parts[0])
	value := bytes.TrimSpace(parts[1])

	return string(field), string(value), nil
}

func isToken(token []byte) bool {
	if len(token) == 0 {
		return false
	}

	for _, ch := range token {
		found := false

		if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9' {
			found = true
		}

		switch ch {
		case '!', '#', '$', '%', '&', '\'', '*', '+', '-', '.', '^', '_', '`', '|', '~', ':':
			found = true
		}

		if !found {
			return false
		}
	}

	return true
}
