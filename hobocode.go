/*
package hobocode provides some lazy convenience functions
for wrapping text in colors, detecting tty, terminal
size, and colored/formatted interactive inputs

The coloring choices are opinionated and the "levels" are based
on common logging patterns but with a focus on Console UX vs. logging style

Note: Most of the printers in this library automatically
new-line fields with Println variants
*/
package hobocode

import (
	"fmt"
	"os"
)

// Stdoutln prints with newline to os.Stdout
func Stdoutln(message string) {
	fmt.Fprintln(os.Stdout, message)
}

// Stderrln prints with newline to os.Stderr
func Stderrln(message string) {
	fmt.Fprintln(os.Stderr, message)
}

// Stdlin prints with newline to *os.File the given message
func Stdlin(fd *os.File, message string) {
	fmt.Fprintln(fd, message)
}
