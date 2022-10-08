package termcolors

import (
	"fmt"
	"math/rand"
	"time"
)

const Black = "\033[30m"
const K = Black
const Red = "\033[31m"
const R = Red
const Green = "\033[32m"
const G = Green
const Yellow = "\033[33m"
const Y = Yellow
const Blue = "\033[34m"
const B = Blue
const Magenta = "\033[35m"
const M = Magenta
const Cyan = "\033[36m"
const C = Cyan
const White = "\033[37m"
const W = White
const Reset = "\033[0m"
const X = Reset
const Clear = "\033[H\033[2J"
const CurOff = "\033[?25l"
const CurOn = "\033[?25h"

// Rand returns a random color ANSI escape between Black (30) and White (37).
func Rand() string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(7)
	return fmt.Sprintf("\033[3%vm", n)
}
