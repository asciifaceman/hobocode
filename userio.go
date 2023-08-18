package hobocode

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
)

// Input acquires user CLI input
func Input(message string) string {
	Ficolor(0, out, text.FgHiYellow, message)
	var ret string
	fmt.Scanln(&ret)
	return ret
}

// Inputd acquires user CLI input with a default
// defaultOption is presented as the default and used on nil entry
func Inputd(defaultOption string, message string) string {
	Ficolorf(0, out, text.FgHiYellow, "%s [%s]: ", message, defaultOption)
	fmt.Scanln(&defaultOption)
	return defaultOption
}

// Inputf acquires user CLI input using a formatted message
// defaultOption is presented as the default and used on nil entry
func Inputf(format string, a ...interface{}) string {
	var ret string
	Ficolorf(0, out, text.FgHiYellow, format, a...)
	fmt.Scanln(&ret)
	return ret
}

// Confirm is a naive confirmation prompt in HiRed that won't break until an explicit answer is given
func Confirm(message string) bool {
	var ret string

	for {
		Ficolorf(0, os.Stdout, text.FgHiWhite, "%s (Y/n): ", message)
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

// Select allows the user to select from a list of options
// the index of the options is used as the selection and the user
// selection is returned as that int
// a response of -1 indicates a failure to capture input
// message introduces the list selection
// prompt is the input line
func Selection(options []string, message string, prompt string) int {
	Icolor(0, out, text.FgHiWhite, message)
	for i, opt := range options {
		Icolorf(1, out, text.FgHiWhite, "%d): %s", i, opt)
	}

	resp := Input(prompt)

	respInt, err := strconv.Atoi(resp)
	if err != nil {
		return -1
	}

	return respInt
}

// Selectr allows the user to select from a list of options with a recommended option
// the index of the options is used as the selection and the user
// selection is returned as that int
// a response of -1 indicates a failure to capture input
// message introduces the list selection
// prompt is the input line
func Selectionr(options []string, recommendedOption int, message string, prompt string) int {
	Icolor(0, out, text.FgHiWhite, message)
	for i, opt := range options {
		if i == recommendedOption {
			Icolorf(1, out, text.FgHiWhite, "%d): %s <--", i, opt)
		} else {
			Icolorf(1, out, text.FgHiWhite, "%d): %s", i, opt)
		}
	}

	resp := Inputd(fmt.Sprintf("%d", recommendedOption), prompt)

	respInt, err := strconv.Atoi(resp)
	if err != nil {
		return -1
	}

	return respInt
}
