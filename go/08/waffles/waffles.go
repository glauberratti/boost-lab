package waffles

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	C "github.com/glauberratti/boost-lab/go/04/termcolors"
	"github.com/glauberratti/boost-lab/go/06/cli"
)

func prompt(in io.Reader) string {
	resp, err := cli.Prompt(in, C.M+"> "+C.Y)

	if err != nil {
		log.Println(err)
		return ""
	}

	fmt.Print(C.Reset)
	return resp
}

func end() {
	fmt.Println("The end.")
	os.Exit(1)
}

var YES = regexp.MustCompile(`(?i)^(y(|es|eah|ep|eap)$|da|affirmative|sure|ok|si)`)

func Run(in io.Reader) {
	fmt.Println(C.R + "Do you lihe waffles?" + C.Y)
	resp := prompt(in)

	//if !strings.HasPrefix(strings.ToLower(resp), "yes") {

	if !YES.MatchString(resp) {
		end()
	}

	fmt.Println(C.R + "Do you lihe pancakes?" + C.Y)

	resp = prompt(in)

	if !YES.MatchString(resp) {
		end()
	}

	log.Println("would continue")
}
