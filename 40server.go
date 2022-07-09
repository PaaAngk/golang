package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Hello World!") })
	http.HandleFunc("/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
		fmt.Fprintf(w, "Host = %q\n", r.Host)
		fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
		}
	})
	http.HandleFunc("/co", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		fmt.Fprintf(w, "Count %d\n", count)
		mu.Unlock()
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		mu.Unlock()
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" || r.Method == "put" || r.Method == "DELETE" {
			fmt.Fprintf(w, "Sorry, only POST methods this link are supported.")
			return
		}
		fmt.Print("Post from website! r.PostFrom")
		fmt.Print(r.Body)
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
