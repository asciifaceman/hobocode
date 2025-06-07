package hobocode

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/text"
)

var (
	out = os.Stdout
	err = os.Stderr
)

// SetOutput sets the output for stdout and stderr
func SetOutput(stdout, stderr *os.File) {
	if stdout != nil {
		out = stdout
	}
	if stderr != nil {
		err = stderr
	}
}

// ResetOutput resets the output for stdout and stderr to os.Stdout and os.Stderr
func ResetOutput() {
	out = os.Stdout
	err = os.Stderr
}

/*
	Stdout helpers
*/

// Info prints the given message with bright cyan if Colorable
//
// Stdout
func Info(message string) {
	Icolor(0, out, text.FgHiCyan, message)
}

// Infof prints the given formatted message with bright cyan if colorable
//
// Stdout
func Infof(format string, a ...interface{}) {
	Icolorf(0, out, text.FgHiCyan, format, a...)
}

// Iinfo prints the given message with bright cyan if Colorable and an indent
//
// Stdout
func Iinfo(indent int, message string) {
	Icolor(indent, out, text.FgHiCyan, message)
}

// Iinfof prints the given formatted message with bright cyan if Colorable and an indent
//
// Stdout
func Iinfof(indent int, format string, a ...interface{}) {
	Icolorf(indent, out, text.FgHiCyan, format, a...)
}

// Success prints the given message with bright green if Colorable
//
// Stdout
func Success(message string) {
	Icolor(0, out, text.FgHiGreen, message)
}

// Successf prints the given formatted message with bright green if Colorable
//
// Stdout
func Successf(format string, a ...interface{}) {
	Icolorf(0, out, text.FgHiGreen, format, a...)
}

// Isuccess prints the given message with bright green if Colorable with an indent
//
// Stdout
func Isuccess(indent int, message string) {
	Icolor(indent, out, text.FgHiGreen, message)
}

// Isuccessf prints the given formatted message with bright green if Colorable with an indent
//
// Stdout
func Isuccessf(indent int, format string, a ...interface{}) {
	Icolorf(indent, out, text.FgHiGreen, format, a...)
}

// Debug prints the given message with bright green if Colorable
//
// Stdout
func Debug(message string) {
	Icolor(0, out, text.FgHiBlue, message)
}

// Debugf prints the given formatted message with bright green if Colorable
//
// Stdout
func Debugf(format string, a ...interface{}) {
	Icolorf(0, out, text.FgHiBlue, format, a...)
}

// Idebug prints the given message with bright green if Colorable with an indent
//
// Stdout
func Idebug(indent int, message string) {
	Icolor(indent, out, text.FgHiBlue, message)
}

// Idebugf prints the given formatted message with bright green if Colorable with an indent
//
// Stdout
func Idebugf(indent int, format string, a ...interface{}) {
	Icolorf(indent, out, text.FgHiBlue, format, a...)
}

/*
	Stderr helpers
*/

// Warn prints the given message with bright yellow if Colorable
//
// Stderr
func Warn(message string) {
	Icolor(0, err, text.FgHiYellow, message)
}

// Warnf prints the given formatted message with bright yellow if Colorable
//
// Stderr
func Warnf(format string, a ...interface{}) {
	Icolorf(0, err, text.FgHiYellow, format, a...)
}

// Iwarn prints the given message with bright yellow if Colorable and an indent
//
// Stderr
func Iwarn(indent int, message string) {
	Icolor(indent, err, text.FgHiYellow, message)
}

// Iwarnf prints the given formatted message with bright yellow if Colorable and an indent
//
// Stderr
func Iwarnf(indent int, format string, a ...interface{}) {
	Icolorf(indent, err, text.FgHiYellow, format, a...)
}

// Error prints the given message with bright red if Colorable
//
// Stderr
func Error(message string) {
	Icolor(0, err, text.FgHiRed, message)
}

// Errorf prints the given formatted message with bright red if Colorable
//
// Stderr
func Errorf(format string, a ...interface{}) {
	Icolorf(0, err, text.FgHiRed, format, a...)
}

// Ierror prints the given message with bright red if Colorable and an indent
//
// Stderr
func Ierror(indent int, message string) {
	Icolor(indent, err, text.FgHiRed, message)
}

// Ierrorf prints the given formatted message with bright red if Colorable and an indent
//
// Stderr
func Ierrorf(indent int, format string, a ...interface{}) {
	Icolorf(indent, err, text.FgHiRed, format, a...)
}
