package greet

import (
	"fmt"
	"log"

	"github.com/glauberratti/boost-lab/go/03-greet/internal"
)

func Greet(name string, lang string) {
	msg, err := internal.MsgHiNiceToMeetYou(lang, name)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(msg)
}
