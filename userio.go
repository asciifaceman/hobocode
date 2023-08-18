package hobocode

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
)

/*
	TODO: User IO stuff is rough and needs attention
*/

// Input acquires user CLI input
// defaultOption is presented as the default and used on nil entry
func Input(defaultOption string, message string) string {
	fmt.Print(text.FgHiYellow.Sprintf("%s [%s]: ", message, defaultOption))
	fmt.Scanln(&defaultOption)
	return defaultOption
}

// Inputf acquires user CLI input using a formatted message
// defaultOption is presented as the default and used on nil entry
func Inputf(defaultOption string, format string, a ...interface{}) string {
	fmt.Print(text.FgHiYellow.Sprintf(format, a...))
	fmt.Scanln(&defaultOption)
	return defaultOption
}

// Confirm is a naive confirmation prompt in HiRed that won't break until an affirmative answer is given
func Confirm(message string) bool {
	var ret string

	for {
		fmt.Print(text.FgHiRed.Sprintf("%s [(y)es / (n)o]: ", message))
		fmt.Scanln(&ret)

		if strings.ToLower(ret) == "y" || strings.ToLower(ret) == "yes" {
			return true
		} else if strings.ToLower(ret) == "n" || strings.ToLower(ret) == "no" {
			return false
		} else {
			continue
		}
	}

}

// Confirmf is a naive confirmation prompt in HiRed that won't break until an affirmative answer is given
// (with a formatted message)
func Confirmf(format string, a ...interface{}) bool {
	msg := fmt.Sprintf(format, a...)
	return Confirm(msg)
}
