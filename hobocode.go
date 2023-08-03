package hobocode

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
)

// Note prints the given message in HiCyan
func Note(message string) {
	fmt.Println(text.FgHiCyan.Sprint(message))
}

// Notef prints the given formatted message in HiCyan
func Notef(format string, a ...interface{}) {
	fmt.Println(text.FgHiCyan.Sprintf(format, a...))
}

// Warn prints the given message in HiMagenta
func Warn(message string) {
	fmt.Println(text.FgHiMagenta.Sprint(message))
}

// Warnf prints the given formatted message in HiMagenta
func Warnf(format string, a ...interface{}) {
	fmt.Println(text.FgHiMagenta.Sprintf(format, a...))
}

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

// Error prints the given message in HiRed
func Error(message string) {
	fmt.Println(text.FgHiRed.Sprint(message))
}

// Errorf prints the given formatted message in HiRed
func Errorf(format string, a ...interface{}) {
	fmt.Println(text.FgHiRed.Sprintf(format, a...))
}

// Success prints the given message in HiGreen
func Success(message string) {
	fmt.Println(text.FgHiGreen.Sprint(message))
}

// Successf prints the given formatted message in HiGreen
func Successf(format string, a ...interface{}) {
	fmt.Println(text.FgHiGreen.Sprintf(format, a...))
}

// Debug prints the given message in HiBlue
func Debug(message string) {
	fmt.Println(text.FgHiBlue.Sprint(message))
}

// Debugf prints the given formatted message in HiBlue
func Debugf(message string, a ...interface{}) {
	fmt.Println(text.FgHiBlue.Sprintf(message, a...))
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
