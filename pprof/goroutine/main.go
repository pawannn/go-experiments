package main

import (
	"net/http"
	_ "net/http/pprof"
)

func leak() {
	ch := make(chan int)
	go func() {
		<-ch
	}()
}

func main() {
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	leak()

	select {}
}
