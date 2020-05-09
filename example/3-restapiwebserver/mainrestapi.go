// build: go build example/3-restapiwebserver/mainrestapi.go
// run: go run example/3-restapiwebserver/mainrestapi.go
// test:
//				GET: curl http://localhost:8080/api/v1/helloworld
//				POST: curl -d '{"name":"indonesia"}' -H "Content-Type: application/json" -X POST http://localhost:8080/api/v1/helloworld
//				DELETE: curl -X DELETE http://localhost:8080/api/v1/helloworld

package main

import (
	"fmt"

	"github.com/jaya-p/gowebservice"
)

func main() {
	fmt.Println("REST API web server is running")
	gowebservice.HelloworldRestAPIWebserver()
}
