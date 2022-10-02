package greet

import (
	"fmt"
	"log"
	"strings"

	"github.com/glauberratti/boost-lab/go/03-greet/internal"
)

func Greet(name string, lang string) {
	msg, err := internal.MsgHiNiceToMeetYou(lang)

	if err != nil {
		log.Println(err)
	}

	msg = strings.Replace(msg, "{1}", name, 1)

	fmt.Println(msg)
}
