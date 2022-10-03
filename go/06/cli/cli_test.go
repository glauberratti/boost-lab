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
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("%q", line)

	// Output:
	// "Sample"
}

func ExamplePrimpt_default() {
	sr := strings.NewReader("Sample\r\n")
	line, err := cli.Prompt(sr, "")

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("%q", line)

	// Output:
	// > "Sample"
}

func ExamplePrimpt_explicit() {
	sr := strings.NewReader("Sample\r\n")
	line, err := cli.Prompt(sr, "--> ")

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("%q", line)

	// Output:
	// --> "Sample"
}
