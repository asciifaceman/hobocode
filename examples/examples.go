package main

import (
	"fmt"
	"os"
	"time"

	"github.com/asciifaceman/hobocode"
)

func main() {
	if len(os.Args) == 1 {
		help()
	}

	arg := os.Args[1]

	if arg == "prints" {
		prints()
	} else if arg == "io" {
		userio()
	} else {
		help()
	}
}

func help() {
	hobocode.Error("Please select an example to see")
	hobocode.Error("prints | io")
	os.Exit(0)
}

func prints() {
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

func userio() {
	in1 := hobocode.Inputd("nvim", "What is your preferred editor? ")
	hobocode.Successf("You chose %s", in1)

	options := []string{
		"Continue seeing user IO examples",
		"Quit seeing user IO examples",
		"What does the owl say",
	}

	resp := hobocode.Selection(options, "Please select from the list", "Your selection (empty is none): ")
	if resp == -1 {
		hobocode.Error("Error fetching input, exiting")
		os.Exit(1)
	} else if resp == 0 {
		hobocode.Info("Continuing...")
	} else if resp == 1 {
		hobocode.Info("User exited")
		os.Exit(0)
	} else if resp == 2 {
		hobocode.Success("Hoooo!")
	}

	options = []string{
		"Continue (again!) seeing user IO examples",
		"See curerent date and continue",
		"Quit seeing user IO examples",
	}

	resp = hobocode.Selectionr(options, 0, "Please select again from the list", "Your selection: ")
	if resp == -1 {
		hobocode.Error("Error fetching input, exiting")
		os.Exit(1)
	} else if resp == 0 {
		hobocode.Info("Continuing...")
	} else if resp == 1 {
		hobocode.Info(time.Now().Format("2006-01-02"))
	} else if resp == 2 {
		hobocode.Info("User exited...")
		os.Exit(0)
	}

	conf := hobocode.Confirm("All good?")
	if conf {
		hobocode.Success("Yay!")
	} else {
		hobocode.Warn("Oh no :(")
	}

}
