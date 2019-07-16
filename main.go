package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

var repetitions = flag.Int("repetitions", 100, "repetitions count")

func main() {
	flag.Parse()

	http.HandleFunc("/leak", leaker{}.handle)
	http.HandleFunc("/noleak", noleaker{}.handle)

	fmt.Println("Starting at :8080 with repetitions:", *repetitions)
	http.ListenAndServe(":8080", nil)
}

var leakyArray = make([]string, 0)

type leaker struct{}

func (leaker) handle(_ http.ResponseWriter, r *http.Request) {
	for i := 0; i < *repetitions; i++ {
		s := strconv.Itoa(i)
		leakyArray = append(leakyArray, s)
	}
	fmt.Println("leaked")
}

type noleaker struct{}

func (noleaker) handle(_ http.ResponseWriter, r *http.Request) {
	var array = make([]string, 0)
	for i := 0; i < *repetitions; i++ {
		s := strconv.Itoa(i)
		array = append(array, s)
	}
	fmt.Println("not leaked")
}
