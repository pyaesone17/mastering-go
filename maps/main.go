package main

import "fmt"

func main() {
	data := make(map[string]int)
	data["one"] = 1
	data["two"] = 2
	data["three"] = 3

	delete(data, "two")

	if _, ok := data["two"]; !ok {
		fmt.Println("two doesn't exist")
	}

	for key, d := range data {
		fmt.Printf("%s : %d \n", key, d)
	}
}
