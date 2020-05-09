// build: go build example/1-cli/main.go
// run: go run example/1-cli/main.go
package main

import (
	"fmt"

	"github.com/jaya-p/gowebservice"
)

func main() {
	fmt.Println(gowebservice.Helloworld())
}
