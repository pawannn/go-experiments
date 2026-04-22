package main

import (
	"fmt"
	"math"
	"net/http"
	_ "net/http/pprof"
)

func hotCalc() {
	for {
		for i := 0; i < 1_000_000; i++ {
			_ = math.Sqrt(float64(i)) * math.Sin(float64(i))
		}
	}
}

func main() {
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	for range 4 {
		go hotCalc()
	}

	fmt.Println("CPU load started")
	select {}
}
