package main

import (
	"fmt"
	"os"

	"github.com/glauberratti/boost-lab/go/05/args/internal/args"
)

func main() {
	fmt.Print(args.Output(os.Args...))
}
