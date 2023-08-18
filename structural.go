package hobocode

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/jedib0t/go-pretty/v6/text"
)

/*
	TODO: Header right justification gets wonky when the terminal width is odd
*/

const (
	defaultWidth = 0
)

// Header prints a header title wrapped by "=" characters sized to the terminal width
//
// Stdout
func Header(title string) {

	/*
		TODO: I've wasted a lot of time here trying to get the header widths to reliably always
		be the same across headers based on screen width but I am too stupid to get it to work in every
		circumstance
	*/

	x, _, err := Size(os.Stdout)
	if err != nil {
		x = defaultWidth
	}
	// safety buffer
	x = x - 1
	//fmt.Printf("width: %d\n", x)

	titleLength := utf8.RuneCountInString(title)
	//fmt.Printf("title len: %d\n", titleLength)

	n := (x - titleLength - 4) / 2
	n2 := x - n*2
	//fmt.Printf("Remainder: %d\n", n2)
	//fmt.Printf("Difference: %d\n", titleLength*2-n2)

	//fmt.Printf("half: %d\n", n)
	offset := 0

	if titleLength*2 > n2 {
		offset--
	}

	if n%2 == 0 {
		if titleLength%2 != 0 {
			offset++
		}
	}

	//fmt.Printf("Offset: %d\n", offset)

	patternLeft := strings.Repeat("=", n)
	patternRight := strings.Repeat("=", n+offset)

	Icolor(0, os.Stdout, text.FgHiWhite, fmt.Sprintf("%s[ %s ]%s", patternLeft, title, patternRight))

}

// HeaderLeft prints a header title offset to the left wrapped by "=" characters sized to the terminal width
//
// Stdout
func HeaderLeft(title string) {

	x, _, err := Size(os.Stdout)
	if err != nil {
		x = defaultWidth
	}

	x = x - 1
	//fmt.Println(x)

	titleLength := utf8.RuneCountInString(title)
	//fmt.Printf("title len: %d\n", titleLength)

	n := (x - titleLength) / 10
	n2 := x - titleLength - n - 1

	offset := 0
	if titleLength%2 != 0 {
		offset++
	}

	patternLeft := ""
	if n-4 > 0 {
		patternLeft = strings.Repeat("=", n-4)
	}

	patternRight := ""
	if n-2+offset > 0 {
		patternRight = strings.Repeat("=", n2-4+offset)
	}

	Icolor(0, os.Stdout, text.FgHiWhite, fmt.Sprintf("%s[ %s ]%s", patternLeft, title, patternRight))

}
