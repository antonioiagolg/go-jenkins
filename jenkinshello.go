package main

import (
	"fmt"
	"io"
	"os"
)

func Hello(o io.Writer, name string) {
	fmt.Fprintf(o, "Hello %v", name)
}

func main() {
	Hello(os.Stdout, "Antonio")
}
