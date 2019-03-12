package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for true {
		result, _ := reader.ReadString('\n')
		writer.WriteString("wow")
		writer.WriteString("ok")
		writer.WriteString(result)
		writer.Flush()
	}
}
