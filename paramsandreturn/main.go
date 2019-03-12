package main

import "fmt"

func main() {
	value := NamingReturn("1", "2")
	Variadic("1", "2")
	fmt.Println(value)
}

func NamingReturn(input1, input2 string) (val string) {
	val = "ok"
	return
}

func Variadic(input ...string) {
	fmt.Println(input)
}
