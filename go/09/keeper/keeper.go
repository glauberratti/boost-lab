package keeper

import (
	"fmt"
	"io"
	"strings"

	C "github.com/glauberratti/boost-lab/go/04/termcolors"
	"github.com/glauberratti/boost-lab/go/06/cli"
)

func Run(r io.Reader) {
	var name, colour string

	//var isTerminal = term.IsTerminal(int(os.Stdout.Fd()))

	C.OnIfTerminal(r)

	fmt.Println(C.Clear + C.C + "Answer these questions three, 'ere the other side ye see.\n" + C.X)

	fmt.Println(C.R + "What is your name?" + C.X)
	name, _ = cli.Prompt(r, C.M+"> "+C.Y)

	fmt.Println(C.R + "What is your quest?" + C.X)
	cli.Prompt(r, C.M+"> "+C.Y)

	if name == "Lancelot" || name == "Galahad" {
		fmt.Println(C.R + "What is yout favorite colour?" + C.X)
		colour, _ = cli.Prompt(r, C.M+"> "+C.Y)
		if name == "Galahad" && strings.HasSuffix(colour, "no") {
			fmt.Println(C.R + "💀 You are thrown into the gulf of eternal peril. 🔥" + C.X)
			return
		}
	}

	if name == "Robin" {
		fmt.Println(C.R + "What is the capital of Assyria?" + C.X)
		cli.Prompt(r, C.M+"> "+C.Y)
	}

	if name == "Arthur" {
		fmt.Println(C.R + "What is the air velocity of an unladen swallow?" + C.X)
		cli.Prompt(r, C.M+"> "+C.Y)
	}
}
