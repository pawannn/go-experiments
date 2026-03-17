package request

import "errors"

var SEPARATOR []byte = []byte("\r\n")
var ERR_MALFORMED_REQ_LINE = errors.New("Malformed request line")
var ERR_UNSUPPORTED_HTTP_VERSION = errors.New("unsupported HTTP version")
var ERR_REQUEST_ERRSTATE = errors.New("Request in Error state")

type parserState int

const (
	stateInit parserState = iota
	stateDone
	stateError
)
