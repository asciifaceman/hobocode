package main

import (
	"fmt"
	"time"

	"github.com/asciifaceman/hobocode"
)

func main() {
	hobocode.Header("Info")
	hobocode.Info("This is an info message")
	hobocode.Infof("This is a formatted info message: %v", time.Now().Format("2006/01/02"))
	hobocode.Iinfo(1, "This is an info message indented once")
	hobocode.Iinfo(2, "This is an info message indented twice")
	hobocode.Iinfof(1, "This is a formatted info message indented once: %v", time.Now().Format("2006/01/02"))
	hobocode.Iinfof(2, "This is a formatted info message indented twice: %v", time.Now().Format("2006/01/02"))

	hobocode.Header("Warn")
	hobocode.Warn("This is a warn message")
	hobocode.Warnf("This is a formatted warn message: %v", time.Now().Format("2006/01/02"))

	hobocode.Header("Error")
	hobocode.Error("This is an error message")
	hobocode.Errorf("This is a formatted error message: %v", fmt.Errorf("error occured at %v", time.Now().Format("2006/01/02")))

	hobocode.Header("Success")
	hobocode.Success("This is a success message!")
	hobocode.Successf("This is a formatted successs message! %v", time.Now().Format("2006/01/02"))

	hobocode.Header("Debug")
	hobocode.Debug("This is a debug message!")
	hobocode.Debugf("This is a formatted debug message! %v", time.Now().Format("2006/01/02"))

	hobocode.HeaderLeft("A left justified header!?")
	hobocode.HeaderLeft("Whoa!")
	hobocode.HeaderLeft("That's cool")

}
