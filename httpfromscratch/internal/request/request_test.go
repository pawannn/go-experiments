package request

import (
	"strings"
	"testing"
)

func TestRequestLineParser(t *testing.T) {
	// Test: Good GET Request line
	r, err := ParseFromReader(strings.NewReader(
		"GET / HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n",
	))

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if r == nil {
		t.Fatalf("expected request, got nil")
	}

	if r.RequestLine.Method != "GET" {
		t.Fatalf("expected method GET, got %s", r.RequestLine.Method)
	}

	if r.RequestLine.RequestTarget != "/" {
		t.Fatalf("expected target /, got %s", r.RequestLine.RequestTarget)
	}

	if r.RequestLine.HTTPVersion != "1.1" {
		t.Fatalf("expected HTTP version 1.1, got %s", r.RequestLine.HTTPVersion)
	}

	// Test: Good GET Request line with path
	r, err = ParseFromReader(strings.NewReader(
		"GET /coffee HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n",
	))

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if r == nil {
		t.Fatalf("expected request, got nil")
	}

	if r.RequestLine.Method != "GET" {
		t.Fatalf("expected method GET, got %s", r.RequestLine.Method)
	}

	if r.RequestLine.RequestTarget != "/coffee" {
		t.Fatalf("expected target /coffee, got %s", r.RequestLine.RequestTarget)
	}

	if r.RequestLine.HTTPVersion != "1.1" {
		t.Fatalf("expected HTTP version 1.1, got %s", r.RequestLine.HTTPVersion)
	}

	// Test: Invalid number of parts in request line
	_, err = ParseFromReader(strings.NewReader(
		"/coffee HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n",
	))

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
