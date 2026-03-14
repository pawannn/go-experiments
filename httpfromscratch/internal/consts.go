package internal

import "errors"

var SEPARATOR = "\r\n"
var ERR_MALFORMED_REQ_LINE = errors.New("Malformed request line")
var ERR_UNSUPPORTED_HTTP_VERSION = errors.New("unsupported HTTP version")
