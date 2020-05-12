// run test: go test -v
package gowebservice

import "testing"

func TestHelloworld(t *testing.T) {
	want := "Hello World"
	if got := Helloworld(); got != want {
		t.Errorf("Helloworld() return wrong output: got %q , want %q", got, want)
	}
}
