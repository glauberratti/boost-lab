package main

import (
	"os"

	"github.com/glauberratti/boost-lab/go/09/keeper"
)

func main() {
	keeper.Run(os.Stdin)
}
