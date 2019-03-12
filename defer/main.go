package main

import "fmt"

func main() {
	func() {
		defer fmt.Println(1)
		defer fmt.Println(2)
	}()

	func() {
		defer fmt.Println(3)
		defer fmt.Println(4)
	}()
}
