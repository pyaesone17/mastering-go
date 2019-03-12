package main

import "fmt"

func main() {
	val := doIt()

	fmt.Println(val)
}

func doIt() (val int) {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg)
			val = 2
		}
	}()

	makePanic()
	return 1
}

func makePanic() {
	panic("error")
}
