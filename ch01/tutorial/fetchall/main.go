package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("fetchall: fetch err:%s, %v", url, err)
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("fetchall: read err:%s, %v", url, err)
	}
	ch <- fmt.Sprintf("%.2fs\t%7d\t%s", time.Since(start).Seconds(), nbytes, url)
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
