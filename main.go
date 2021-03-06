package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const (
	port = "0.0.0.0:8080"
)

var calls = 0
var hostname = "host with no name"

func helloWorld(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "a host with no name"
	}
	if !strings.HasSuffix(r.URL.Path, ".ico") {
		calls++
	}
	fmt.Fprintf(w, "Hello, %s! I'm %s and I've served %d requests.\n", r.RemoteAddr, hostname, calls)
}

func init() {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "a host with no name"
	}

	fmt.Printf("Started server: '%s' at http://%v.\n", hostname, port)
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(port, nil)
}

func main() {}
