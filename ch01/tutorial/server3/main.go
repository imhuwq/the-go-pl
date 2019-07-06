package main

import (
	"fmt"
	"log"
	"net/http"
)

func debugHandler(rsp http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(rsp, "%s %s %s\n", req.Method, req.URL, req.Proto)
	for k, v := range req.Header {
		_, _ = fmt.Fprintf(rsp, "Header[%q] = %q\n", k, v)
	}
	_, _ = fmt.Fprintf(rsp, "Host = %q\n", req.Host)
	_, _ = fmt.Fprintf(rsp, "Remote Address = %q\n", req.RemoteAddr)
	if err := req.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range req.Form {
		_, _ = fmt.Fprintf(rsp, "Form[%q] = %q\n", k, v)
	}
}

func main() {
	http.HandleFunc("/", debugHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
