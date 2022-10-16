package eightball_test

import (
	"fmt"
	"log"

	"github.com/glauberratti/boost-lab/go/06/eightball"
)

func ExampleRespond() {
	resp := eightball.Respond()
	fmt.Print(len(resp) > 0)

	// Output:
	// true
}

func ExampleGenRandom() {
	respLen := len(eightball.GetResponses())
	n := eightball.GenRandom(respLen)
	log.Println(n)
	fmt.Print(n >= 0 && n < respLen)

	// Output:
	// true
}
