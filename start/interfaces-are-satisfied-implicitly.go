package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(b []byte) (n int, err error)
}
type Writer interface {
	Write(b []byte) (n int, err error)
}
type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var w Writer
	w = os.Stdout
	fmt.Println(w, "Hello writer\n")
}
