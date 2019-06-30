package main

import (
	"fmt"
	"io"
	"net/url"
	"net/http"
	"os"
	"strings"
	"time"
)

func fetch(uri string, ch chan<- string) {
	start := time.Now()
	if !strings.HasPrefix(uri, "http") {
		uri = "http://" + uri
	}
	resp, err := http.Get(uri)
	if err != nil {
		ch <- fmt.Sprintf("fetchall: fetch err:%s, %v", uri, err)
		return
	}
	file, err := os.Create("./" + url.QueryEscape(uri) + ".html")
	if err != nil {
		ch <- fmt.Sprintf("fetchall: save err: %s, %v", uri, err)
		return
	}

	nbytes, err := io.Copy(file, resp.Body)
	_ = file.Close()
	_ = resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("fetchall: read err:%s, %v", uri, err)
		return
	}
	ch <- fmt.Sprintf("%.2fs\t%7d\t%s", time.Since(start).Seconds(), nbytes, uri)
}

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for _, _ = range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("Time Costs: %.2fs\n", time.Since(start).Seconds())
}
