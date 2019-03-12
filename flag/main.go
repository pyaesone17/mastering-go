package main

import (
	"flag"
	"fmt"
	"strings"
)

// go run main.go -len=14 -i=false -names=veve,nyan lmao ok

func main() {
	i := flag.Bool("i", true, "i")
	length := flag.Int("len", 0, "len")
	var manyNames NamesFlag
	flag.Var(&manyNames, "names", "Comma-separated list")
	flag.Parse()

	fmt.Println(*i, *length)

	for i, item := range manyNames.GetNames() {
		fmt.Println(i, item)
	}

	fmt.Println("Remaining command-line arguments:")
	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}
}

type NamesFlag struct {
	Names []string
}

func (s *NamesFlag) GetNames() []string {
	return s.Names
}

func (s *NamesFlag) String() string {
	return fmt.Sprint(s.Names)
}

func (s *NamesFlag) Set(v string) error {
	if len(s.Names) > 0 {
		return fmt.Errorf("Cannot use names flag more than once!")
	}
	names := strings.Split(v, ",")
	for _, item := range names {
		s.Names = append(s.Names, item)
	}
	return nil
}
