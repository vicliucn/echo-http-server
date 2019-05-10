package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"
)

func sysEchoHandler(w http.ResponseWriter, r *http.Request) {
	r.Write(w)
}

func userEchoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.RequestURI, r.Proto)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	for key, values := range r.Header {
		fmt.Fprintf(w, "%s: %s\n", key, strings.Join(values, ", "))
	}
	fmt.Fprintln(w)
	io.Copy(w, r.Body)
}

func main() {
	var addr string
	var url string
	flag.StringVar(&addr, "addr", "127.0.0.1:8000", "listen addr")
	flag.StringVar(&url, "url", "", "serve url")
	flag.Parse()
	http.HandleFunc(path.Join("/", url, "sys"), userEchoHandler)
	http.HandleFunc(path.Join("/", url, "user"), sysEchoHandler)
	http.HandleFunc(path.Join("/", url), sysEchoHandler)
	http.ListenAndServe(addr, nil)
}
