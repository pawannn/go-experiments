package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

var store [][]int

func HandleMemory() {
	data := make([]int, 1024*1024*5)

	for i := range 10 {
		data[i] = i
	}

	store = append(store, data)
}

func main() {

	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	for range 30 {
		HandleMemory()
	}

	fmt.Println("done")

	select {}
}
