package greet_test

import greet "github.com/glauberratti/boost-lab/go/03-greet"

func ExampleGreet() {
	greet.Greet("", "en")

	// Output:
	// Hi, nice to meet you!
}

func ExampleGreet_with_arguments() {
	greet.Greet("Glauber", "en")

	// Output:
	// Hi Glauber, nice to meet you!
}
