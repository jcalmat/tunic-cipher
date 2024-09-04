package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jcalmat/tunic-cipher/pkg/fyne"
	"github.com/rs/zerolog"
)

func main() {

	// setup logger
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	fyneApp := fyne.New(&logger)

	fyneApp.MakeTray()
	fyneApp.CreateWindow()
	fyneApp.Run()

	// take input from the command line
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <input>\n", os.Args[0])
		return
	}

	args := os.Args[1:]

	writeLogs(fmt.Sprintf(strings.Join(args[0:], " ") + "\n"))
	defer writeLogs("\n")

	for _, arg := range args {
		// convert the string to an integer
		if arg == "." {
			writeLogs(" ")
			continue
		}

		// if the first char is a backslash, print the phoneme as is
		if arg[0] == '+' {
			// print the phoneme as is
			writeLogs(fmt.Sprintf("%s", arg[1:]))
			continue
		}

		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		rune := alphabetMap[num]
		if rune.phoneme == " " || rune.phoneme == "" {
			// print question mark in red
			writeLogs("\033[31m?\033[0m")
			continue
		}

		if rune.confident {
			// print in green
			writeLogs(fmt.Sprintf("\033[32m%s\033[0m", rune.phoneme))
		} else {
			// print in yellow
			writeLogs(fmt.Sprintf("\033[33m%s\033[0m", rune.phoneme))
		}
	}
	writeLogs("\n")
}

func writeLogs(log string) {
	fmt.Print(log)

	// open a file for writing
	f, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	f.WriteString(log)
}
