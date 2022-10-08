package eightball

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/glauberratti/boost-lab/go/06/cli"
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
	"It is certain.",
	"It is decidedly so",
	"Without a doubt.",
	"Yes definitely.",
	"You may rely on it.",
	"As I see it, yes.",
	"Most likely.",
	"Outlook good.",
	"Yes",
	"Signs point to yes.",
	"Reply hazy, try again.",
	"Ask again later.",
	"Better not tell you now.",
	"Cannot predict now.",
	"Concentrate and ask again.",
	"Don't count on it.",
	"My reply is no.",
	"My sources say no.",
	"Outlook not so good.",
	"Very doubtful.",
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

// Respond will return a random response from a list of Responses.
func Respond() string {
	i := GenRandom(len(GetResponses()))
	res := GetResponses()
	return res[i]
}

// Run starts an intereactive eightball session prompting the user for imput and answering and then repeating until interrupted.
func Run() {
	fmt.Print(art)
	fmt.Println("🎱 Welcome to the magic eightball!")
	fmt.Println("(Enter your yes or no question)")

	for {
		cli.Prompt(os.Stdin, "🎱 ")
		fmt.Println(Respond())
	}
}
