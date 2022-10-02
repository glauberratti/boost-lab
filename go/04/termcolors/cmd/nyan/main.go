package main

import (
	"fmt"
	"time"

	C "github.com/glauberratti/boost-lab/go/04/termcolors"
)

func main() {
	//defer fmt.Print(C.CurOn)

	//fmt.Print(C.CurOff)
	for {
		fmt.Print(C.Rand() + "nyan" + C.Reset)
		time.Sleep(100 * time.Microsecond)
	}
}
