package hobocode

import (
	"os"
	"strings"

	"golang.org/x/term"
)

/*
	These aren't really necessary to wrap
	but I done did it anyways
*/

// Colorable reports if the given *os.File is
// a terminal or not (os.Stdout, os.Stderr)
func Colorable(pipe *os.File) bool {
	return term.IsTerminal(int(pipe.Fd()))
}

// Size returns the terminal size or error
func Size(pipe *os.File) (int, int, error) {
	return term.GetSize(int(pipe.Fd()))
}

// Tindent returns a tabbed indent repeated n times
func Tindent(n int) string {
	return strings.Repeat("\t", n)
}
