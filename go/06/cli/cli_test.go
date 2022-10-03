package cli_test

import (
	"fmt"
	"strings"

	"github.com/glauberratti/boost-lab/go/06/cli"
)

func ExampleReadLine() {
	sr := strings.NewReader("Sample\r\n")
	line, err := cli.ReadLine(sr)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%q", line)

	// Output:
	// "Sample"
}
