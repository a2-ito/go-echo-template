package helloworld_test

import (
	"helloworld"
	"testing"
)

func TestHelloworld(t *testing.T) {
	if !(helloworld.Helloworld() == "Hello world!!!") {
		t.Error(`miss`)
	}
}
