package fyne

import (
	"fyne.io/fyne/v2"
)

// View defines the data structure for a view
type View struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
}

var (
	// Views defines the metadata for each view
	Views = map[string]View{
		"transcriptor": {"Transcriptor",
			"Transcriptor",
			transcript,
		},
		"lexicon": {"Lexicon",
			"Lexicon\n\nClick on a phoneme to update it",
			lexicon,
		},
	}

	// ViewIndex  defines how our views should be laid out in the index tree
	ViewIndex = map[string][]string{
		"": {"transcriptor", "lexicon"},
	}
)
