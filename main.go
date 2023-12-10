package main

import (
	"fmt"
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	b, _ := io.ReadAll(req.Body)
	w.Write(b)
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe("0.0.0.0:8080", nil)
	//eh
}
