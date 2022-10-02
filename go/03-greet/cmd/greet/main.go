package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	greet "github.com/glauberratti/boost-lab/go/03-greet"
	"github.com/glauberratti/boost-lab/go/03-greet/internal"
)

func main() {
	var name string
	var err error

	if len(os.Args) > 1 {
		name = strings.Join(os.Args[1:], " ")
	}

	for name == "" {
		fmt.Println("Hello there, what's your name?")
		name, err = internal.ReadLine(os.Stdin)

		if err != nil {
			log.Println(err)
			return
		}

	}

	greet.Greet(name)
}
