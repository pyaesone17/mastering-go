package main

import "fmt"

type Person struct {
	name string
}

// Pointer receiver
func (person *Person) UpdateName(val string) {
	person.name = val
}

// No Pointer
func (person Person) UpdateName2(val string) {
	person.name = val
}

func UpdateName(person *Person, val string) {
	person.name = val
}

func main() {
	p := &Person{name: "Ve Ve"}
	p.UpdateName("Nyan Win")
	p.UpdateName2("Pyae Sone Nyein")

	fmt.Println(p.name)
	fmt.Println(p.name)
}
