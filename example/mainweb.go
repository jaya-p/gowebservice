// build: go build example/mainWeb.go
// run: go run example/mainWeb.go
// test: curl http://localhost:8080/helloworld
package main

import (
	"fmt"

	"github.com/jaya-p/gowebservice"
)

func main() {
	fmt.Println("web server is running")
	gowebservice.HelloworldWebserver()
}
