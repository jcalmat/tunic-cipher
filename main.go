package main

import (
	"os"

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
}
