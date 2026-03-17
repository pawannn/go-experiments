package request

import "errors"

var SEPARATOR []byte = []byte("\r\n")
var errMalfunctionedLine = errors.New("Malformed request line")
var errUnsupportedHTTPVersion = errors.New("unsupported HTTP version")
var errStateRequest = errors.New("Request in Error state")

type parserState int

const (
	stateInit parserState = iota
	stateHeader
	stateDone
	stateError
)
