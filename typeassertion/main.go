package main

import "fmt"

func main() {
	var blah interface{}
	blah = 5

	if val, ok := blah.(int); ok {
		fmt.Println(val)
	}

	switch val := blah.(type) {
	case int:
		fmt.Println("it is int")
		fmt.Println(val)
	case string:
		fmt.Println("it is string")
		fmt.Println(val)
	}
}
