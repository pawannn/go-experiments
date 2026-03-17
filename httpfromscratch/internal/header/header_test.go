package header

import "testing"

func TestParseHeader(t *testing.T) {
	data := []byte("Host: localhost:8080\r\n\r\n")
	headers := NewHeader()
	_, done, err := headers.Parse(data)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if !done {
		t.Errorf("expected done: true, got : %v", done)
	}

	host := headers.Get("Host")

	if host != "localhost:8080" {
		t.Errorf("expected host: localhost:8080, got : %s", host)
	}

}
