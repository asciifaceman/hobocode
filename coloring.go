package hobocode

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/text"
)

// Icolor prints the given string at the given indent
// to the given os.File with the given text.Color, if colorable
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
