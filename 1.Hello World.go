// https://gowebexamples.com/hello-world/

// Registering a Request Handler
// First, create a Handler which receives all incomming HTTP connections from browsers, HTTP clients or API requests. A handler in Go is a function with this signature:
// func (w http.ResponseWriter, r *http.Request)

// 1. An http.ResponseWriter which is where you write your text/html response to.
// 2. An http.Request which contains all information about this HTTP request including things like the URL or header fields.

// Registering a request handler to the default HTTP Server is as simple as this:

// http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
// })

package main

// Modul Go
import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // fungsi handler untuk menerima seluruh requests http
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":80", nil) //listen http connection

// go run 1.Hello\ World.go # buka localhost
// ctrl + x (untuk close/interrupt)