package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", Index)

	http.ListenAndServe(":8000", mux)
}
