package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("It prints Hello Antonio", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Hello(&buffer, "Antonio")
		got := string(buffer.String())
		want := "Hello, Antonio"
		if got != want {
			t.Errorf("Got %v, but want %v", got, want)
		}
	})
}
