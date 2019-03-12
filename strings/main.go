package main

import "fmt"

func main() {
	greeting := "hello world"
	results := Revert([]byte(greeting))

	fmt.Println(string(results))
}

func Revert(data []byte) []byte {
	if len(data) == 0 {
		return data
	}

	val := data[:1]
	data = data[1:]

	result := Revert(data)
	result = append(result, val[0])

	return result
}
