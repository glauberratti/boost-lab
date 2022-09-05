package main

import (
	"fmt"
	"os"
)

func main() {
	name := "there"

	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	}

	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	fmt.Printf("Hi %v!\n", name)
}
