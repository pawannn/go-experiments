package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/pawannn/httpfromscratch/internal"
)

func getLines(f io.ReadCloser) <-chan string {
	out := make(chan string)

	go func() {
		defer f.Close()
		defer close(out)
		str := ""
		for {
			data := make([]byte, 8)
			n, err := f.Read(data)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Closed due to error : ", err.Error())
				return
			}

			data = data[:n]
			if i := bytes.IndexByte(data, '\n'); i != -1 {
				str += string(data[:i])
				data = data[i+1:]
				out <- str
				str = ""
			}

			str += string(data)
		}

		if str != "" {
			out <- str
		}
	}()

	return out
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			break
		}

		r, err := internal.RequestFromReader(conn)
		if err != nil {
			break
		}

		fmt.Println("HTTP Version : ", r.RequestLine.HTTPVersion)
		fmt.Println("Method : ", r.RequestLine.Method)
		fmt.Println("Target : ", r.RequestLine.RequestTarget)
	}
}
