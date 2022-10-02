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
	lang := "en"

	if len(os.Args) > 1 {
		name = strings.Join(os.Args[1:], " ")
	}

	for name == "" {
		var msg string

		msg, err = internal.MsgHelloThereWhatsYourName(lang)

		if err != nil {
			log.Println(err)
		}

		fmt.Println(msg)
		name, err = internal.ReadLine(os.Stdin)

		if err != nil {
			log.Println(err)
			return
		}

	}

	greet.Greet(name, lang)
}
