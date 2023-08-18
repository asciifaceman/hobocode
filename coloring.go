package hobocode

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/text"
)

/*
	TODO: Should possibly check terminal width and try to nicely wrap to the indent if width is going to exceed
*/

// Icolor prints the given string at the given indent
// to the given os.File with the given text.Color, if colorable
// with newline
func Icolor(indent int, fd *os.File, color text.Color, message string) {
	id := Tindent(indent)
	if Colorable(fd) {
		fmt.Fprint(fd, id)
		Stdlin(fd, color.Sprint(message))
	} else {
		fmt.Fprint(fd, id)
		Stdlin(fd, message)
	}
}

// Icolorf prints the given format at the given indent
// to the given os.File with the given text.Color, if colorable
// with newline
func Icolorf(indent int, fd *os.File, color text.Color, format string, a ...interface{}) {
	id := Tindent(indent)
	if Colorable(fd) {
		fmt.Fprint(fd, id)
		Stdlin(fd, color.Sprintf(format, a...))
	} else {
		fmt.Fprint(fd, id)
		Stdlin(fd, fmt.Sprintf(format, a...))
	}
}

// Ficolor prints the given string at the given indent
// to the given os.File with the given text.Color, if colorable
// and does not newline
func Ficolor(indent int, fd *os.File, color text.Color, format string, a ...interface{}) {
	id := Tindent(indent)
	if Colorable(fd) {
		fmt.Fprint(fd, id)
		Std(fd, color.Sprintf(format, a...))
	} else {
		fmt.Fprint(fd, id)
		Std(fd, fmt.Sprintf(format, a...))
	}
}

// Ficolorf prints the given format at the given indent
// to the given os.File with the given text.Color, if colorable
// and does not newline
func Ficolorf(indent int, fd *os.File, color text.Color, format string, a ...interface{}) {
	id := Tindent(indent)
	if Colorable(fd) {
		fmt.Fprint(fd, id)
		Std(fd, color.Sprintf(format, a...))
	} else {
		fmt.Fprint(fd, id)
		Std(fd, fmt.Sprintf(format, a...))
	}
}
