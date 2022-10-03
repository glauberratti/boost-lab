package eightball

import (
	"fmt"
	"math/rand"
	"time"
)

const art = `
        ____
    ,dP9CGG88@b,
  ,IP  _   Y888@@b,
 dIi  (_)   G8888@b
dCII  (_)   G8888@@b
GCCIi     ,GG8888@@@
GGCCCCCCCGGG88888@@@
GGGGCCCGGGG88888@@@@...
Y8GGGGGG8888888@@@@P.....
 Y88888888888@@@@@P......
  Y8888888@@@@@@@P'......
    @@@@@@@@@P'.......
        """"........
`

var responses = []string{
	"Yes",
	"No",
	"Maybe",
}

func GetResponses() []string {
	return responses
}

func GenRandom(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func Print8BallArt() {
	fmt.Print(art)
}

// Prompt prints the oprional prompt (> by default) and returns the string entered by the user.
func Prompt(args ...string) string {
	var val string
	p := "> "
	if len(args) > 0 {
		p = args[0]
	}
	fmt.Print(p)
	return val
}

// Respond will return a random response from a list of Responses.
func Respond() string {
	i := GenRandom(len(GetResponses()))
	res := GetResponses()
	return res[i]
}
