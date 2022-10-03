package cli

import (
	"bufio"
	"io"
	"strings"
)

/*
func ReadLine(r io.Reader) (string, error) {
	return bufio.NewReader(r).ReadString('\n')
}
*/

//Readline takes any io.Reader and retuns a trimmed string (initial and trailing white space) or an empty string and error if any error is encountered.
func ReadLine(r io.Reader) (string, error) {
	line, err := bufio.NewReader(r).ReadString('\n')
	line = strings.TrimSpace(line)
	return line, err
}
