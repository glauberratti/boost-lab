package internal_test

import (
	"fmt"
	"strings"

	"github.com/glauberratti/boost-lab/go/03-greet/internal"
)

func ExampleReadLine() {
	sr := strings.NewReader("Sample\r\n")
	line, err := internal.ReadLine(sr)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%q", line)

	// Output:
	// "Sample"
}
