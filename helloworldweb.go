package gowebservice

import (
	"fmt"
	"net/http"
)

// HelloworldHandler is helloworld that is access through http protocol
func HelloworldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, Helloworld())
}

// HelloworldWebserver is handling serve web request
func HelloworldWebserver() {
	http.HandleFunc("/helloworld", HelloworldHandler)

	http.ListenAndServe(":8080", nil)
}
