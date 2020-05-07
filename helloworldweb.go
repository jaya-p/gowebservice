package gowebservice

import (
	"fmt"
	"net/http"
)

// HelloworldHandler is helloworld that is access through http protocol
func HelloworldHandler(w http.responseWriter, r *http.request) {
	fmt.Fprintf(w, helloworld())
}

// HelloworldWebserver is handling serve web request
func HelloworldWebserver() {
	http.HandleFunc("/helloworld", HelloworldHandler)

	http.ListenAndServe(":8080", nil)
}
