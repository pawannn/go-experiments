package header

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Headers map[string]string

func NewHeader() Headers {
	return make(map[string]string)
}

func (h Headers) Set(key string, value string) {
	name := strings.ToLower(key)
	val, ok := h[name]
	if !ok {
		h[name] = value
	} else {
		h[name] = fmt.Sprintf("%s, %s", val, value)
	}
}

func (h Headers) Replace(key string, value string) {
	h[strings.ToLower(key)] = value
}

func (h Headers) ContentLength() int {
	val, ok := h["Content-Length"]
	if !ok {
		return 0
	}

	contentLength, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}

	return contentLength
}

func (h Headers) SetContentLength(contentLength int) {
	h["Content-Length"] = fmt.Sprintf("%d", contentLength)
}

func (h Headers) ParseHeaders(headerLines []byte) (int, bool, error) {
	read := 0
	done := false

	for {
		idx := bytes.Index(headerLines[read:], rn)
		if idx == -1 {
			break
		}

		if idx == 0 {
			done = true
			read += len(rn)
			break
		}

		key, value, err := parse(headerLines[idx:])
		if err != nil {
			return -1, false, err
		}
		h.Set(key, value)

		read += idx + len(rn)
	}

	return read, done, nil
}

func (h Headers) ForEach(cb func(k, v string)) {
	for key, val := range h {
		cb(key, val)
	}
}

func parse(headerLine []byte) (string, string, error) {
	parts := bytes.SplitN(headerLine, []byte(":"), 2)
	if len(parts) != 2 {
		return "", "", errMalfunctionedHeaderLine
	}

	key := string(parts[0])
	value := string(parts[1])
	return key, value, nil
}

func validHeaderName(headerName []byte) bool {
	for _, ch := range headerName {
		found := false
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9' {
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

	return false
}
