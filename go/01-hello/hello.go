package main

import (
	"fmt"
	"log"
)

func main() {
	println("Hello World! - \"println\"")         // goes to stderr (os.Stderr)
	fmt.Println("Hello World! - \"fmt.Println\"") // goes to stdout (os.Stdout)
	log.Println("Hello World! - \"log.Println\"") // goes to stderr (os.Stderr)
}
