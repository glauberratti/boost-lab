package args_test

import (
	"testing"

	"github.com/glauberratti/boost-lab/go/05/args/internal/args"
)

/*
func TestPrint(t *testing.T) {
	got := output()
	if got != "test" {
		t.Errorf("\nWant:	%q\nGot:	%q", "this", got)
	}
}
*/

const testOne = `$0 --> "./args"
$1 --> "first"
$2 --> "second"
$3 --> "third"
`

func TestPrint(t *testing.T) {
	got := args.Output("./args", "first", "second", "third")
	if got != testOne {
		t.Errorf("\nWant:	%q\nGot:	%q\n", testOne, got)
	}
}
