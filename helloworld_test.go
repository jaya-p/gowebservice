package gowebservice

import "testing"

func TestHelloworld(t *testing.T) {
	want := "Hello World"
	if got := Helloworld(); got != want {
		t.Errorf("Helloworld() = %q , want %q", got, want)
	}
}