package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
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

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sysEchoHandler(w, r)
}

func main() {
	var addr string
	flag.StringVar(&addr, "addr", "127.0.0.1:8000", "listen addr")
	flag.Parse()
	var h handler
	http.ListenAndServe(addr, h)
}
