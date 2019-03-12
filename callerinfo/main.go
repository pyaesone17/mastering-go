package main

import (
	"fmt"
	"runtime"
)

func main() {
	call()
}

func call() {
	_, file, no, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("called from %s#%d\n", file, no)
	}
}
