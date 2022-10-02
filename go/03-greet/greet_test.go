package greet_test

import greet "github.com/glauberratti/boost-lab/go/03-greet"

func ExampleGreet() {
	greet.Greet()

	// Output:
	// Hi, what's your name?
}

func ExampleGreet_with_arguments() {
	greet.Greet("Glauber")

	// Output:
	// Hi Glauber, nice to meet you!
}
