package request

import "fmt"

var rn = []byte("\r\n")

var errMalfunctionedRequest error = fmt.Errorf("malfunctioned request line")
var errInvalidMethod error = fmt.Errorf("invalid method")
var errUnsupportedProtocol error = fmt.Errorf("invalid protocol")

type parserState int

const (
	stateInit parserState = iota
	stateHeaders
	stateBody
	stateDone
	stateError
)

var errParserState error = fmt.Errorf("parser in error state")
