package main

import (
	"io"
	"os"
)

// tmp/output 2>&1
// 2> err.log
// 1> output.log

func main() {

	io.WriteString(os.Stderr, "STD error \n")
	io.WriteString(os.Stdout, "STD output \n")
}
