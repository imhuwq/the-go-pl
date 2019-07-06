package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func indexHandler(rsp http.ResponseWriter, req *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	_, _ = fmt.Fprintf(rsp, "Request Path: %q\n", req.URL.Path)
}

func countHandler(rsp http.ResponseWriter, req *http.Request) {
	mu.Lock()
	_, _ = fmt.Fprintf(rsp, "Request Count: %d\n", count)
	mu.Unlock()
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/count", countHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
