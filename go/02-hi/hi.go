package main

import (
	"fmt"
	"os"
)

func main() {
	var name string

	for i := 1; i <= len(os.Args)-1; i++ {
		name += os.Args[i] + " "
	}

	if name == "" {
		name = "there"
	} else {
		name = string(name[0 : len(name)-1])
	}

	fmt.Printf("Hi %v!\n", name)
}
