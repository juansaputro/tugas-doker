package main

import (
	"io"
	"net/http"
)

func endpoint1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello this is endpoint1")
}

func endpoint2(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello this is endpoint2")
}

func getroot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Choose between /endpoint1 or /endpoint2")
}

func main() {

	http.HandleFunc("/", getroot)
	http.HandleFunc("/endpoint1", endpoint1)
	http.HandleFunc("/endpoint2", endpoint2)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}