package header

import "errors"

var rn = []byte("\r\n")
var errMalfunctionedHeaderLine = errors.New("malfunctioned header line")
var errMalfunctionedHeaderName = errors.New("malfunctioned header name")
var errMalfunctionedHeaderValue = errors.New("malfunctioned header value")
