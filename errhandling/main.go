package main

import (
	"errors"
	"fmt"
)

func main() {
	err := makeCustomizeErr()
	if msg, ok := err.(Customize); ok {
		fmt.Println(msg)
		fmt.Println("Blah error")
	} else {
		fmt.Println("normal error")
	}

	err = makeNormalErr()
	if msg, ok := err.(Customize); ok {
		fmt.Println(msg)
		fmt.Println("Blah error")
	} else {
		fmt.Println("normal error")
	}
}

// Customize error
type Customize struct {
	msg string
}

func (blah Customize) Error() string {
	return blah.msg
}

func makeCustomizeErr() error {
	return Customize{msg: "Something went wrong"}
}

func makeNormalErr() error {
	return errors.New("Wow")
}
