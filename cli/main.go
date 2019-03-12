package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		log.Fatal("Please enter the flag")
	}

	number, err := strconv.Atoi(arguments[1])
	if err != nil {
		log.Fatal("Please enter the correct integer")
	}

	fmt.Println(number)
}
