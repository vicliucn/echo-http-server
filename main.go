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
	var b strings.Builder
	r.Write(&b)
	var s = b.String()
	w.Write([]byte(s))
	fmt.Println("++++++++")
	fmt.Println(s)
	fmt.Println("++++++++")
}

func userEchoHandler(w http.ResponseWriter, r *http.Request) {
	var b strings.Builder
	fmt.Fprintf(&b, "%s %s %s\n", r.Method, r.RequestURI, r.Proto)
	fmt.Fprintf(&b, "Host: %s\n", r.Host)
	for key, values := range r.Header {
		fmt.Fprintf(&b, "%s: %s\n", key, strings.Join(values, ", "))
	}
	fmt.Fprintln(&b)
	io.Copy(&b, r.Body)
	var s = b.String()
	w.Write([]byte(s))
	fmt.Println("++++++++")
	fmt.Println(s)
	fmt.Println("++++++++")
}

func main() {
	var addr string
	var url string
	flag.StringVar(&addr, "addr", "127.0.0.1:8000", "listen addr")
	flag.StringVar(&url, "url", "", "serve url")
	flag.Parse()
	http.HandleFunc(path.Join("/", url), sysEchoHandler)
	http.HandleFunc(path.Join("/", url, "sys"), sysEchoHandler)
	http.HandleFunc(path.Join("/", url, "user"), userEchoHandler)
	http.ListenAndServe(addr, nil)
}
