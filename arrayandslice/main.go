package main

import "fmt"

func main() {
	arr := [2]int{0, 1}
	mutateArray(arr)
	fmt.Println(arr)

	slice := []int{0, 1}
	mutateSlice(slice)
	fmt.Println(slice)

	arr2 := arr
	arr2[0] = 21
	fmt.Println(arr2)

	slice2 := slice
	slice2[0] = 21
	fmt.Println(slice)
}

func mutateArray(arr [2]int) {
	arr[0] = 15
}

func mutateSlice(arr []int) {
	arr[0] = 15
}
