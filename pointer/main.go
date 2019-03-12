package main

import "fmt"

func main() {
	one := 1
	two := &one
	*two = 3

	fmt.Println(one, *two)
}
