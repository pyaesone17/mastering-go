package main

import "fmt"

func main() {
	one := 1
	two := &one
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("%T \n %p \n %d\n %v \n %#v \n", two, two, *two, two, two)
	fmt.Printf("%t\n", true)
	fmt.Printf("%q \n", slice)
}
