package gowebservice

import (
	"encoding/json"
	"io"
	"net/http"
        "fmt"
)

// structure of POST json
type postRequest struct {
	Name string
}

// GetHandler handles GET method
func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, Helloworld())
}

// PostHandler handles POST method
func PostHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var pr postRequest
	err := decoder.Decode(&pr)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, Helloworld()+", "+pr.Name+"!")
}

// StatusNotFoundHandler handles unknown request method
func StatusNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, "Your requested method ("+r.Method+") is not found")
}

// HelloworldRestAPIHandler is handles REST API (now only GET and POST) requests/responses
func HelloworldRestAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Set HTTP Response "Content-Type" header as JSON
	w.Header().Set("Content-type", "text/plain; charset=utf-8")

	// Set HTTP Response Body, based on HTTP request method
	switch r.Method {
	case "GET":
		GetHandler(w, r)
	case "POST":
		PostHandler(w, r)
	default:
		StatusNotFoundHandler(w, r)
	}
}

// HelloworldRestAPIWebserver provides connectivity for handling REST API request
func HelloworldRestAPIWebserver(pNum uint) {
	http.HandleFunc("/api/v1/helloworld", HelloworldRestAPIHandler)

	http.ListenAndServe(":"+fmt.Sprint(pNum), nil)
}
