package internal

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HTTPVersion string
	RequestType string
	Method      string
}
