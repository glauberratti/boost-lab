package main

import (
	"fmt"
)

/*
func output() string {
	return "test"
}
*/

func output(args ...string) string {
	var buf string
	for n, val := range args {
		buf += fmt.Sprintf("$%v --> %q\n", n, val)
	}

	return buf
}

func main() {
	fmt.Print(output())
}
