package main

import "fmt"

const (
	One = "One"
	Two = "Two"
)

const (
	Tit = iota
	Nit
)

func main() {
	fmt.Println(Tit)
	fmt.Println(Nit)
}
