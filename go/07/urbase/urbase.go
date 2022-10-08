package urbase

import "fmt"

func Run() {
	fmt.Println("All your base ...")
	fmt.Printf("Base   2   8  10  16\n")
	for i := 0; i < 16; i++ {
		fmt.Printf("%8b %3o %3d %3x\n", i, i, i, i)
	}
}
