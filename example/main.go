// build: cd example/ && go build ./main.go
// run: cd example/ && go run ./main.go
package main

import (
	"fmt"

	"github.com/jaya-p/gowebservice"
)

func main() {
	fmt.Println(gowebservice.Helloworld())
}
