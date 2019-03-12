package main

import "fmt"

type PointMe struct {
	Val int
}

func main() {
	one := &PointMe{}
	two := one

	two.Val = 1
	(*two).Val = 2

	fmt.Println(one, two)
}
