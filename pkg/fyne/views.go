package fyne

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// View defines the data structure for a view
type View struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
}

var (
	// Views defines the metadata for each view
	Views = map[string]View{
		// "welcome": {"Welcome", "", welcomeScreen},
		// "canvas": {"Canvas",
		// 	"See the canvas capabilities.",
		// 	canvasScreen,
		// },
		"transcriptor": {"Transcriptor",
			"Transcript Tunic text to phonemes.",
			transcript,
		},
		"lexicon": {"Lexicon",
			"List of all symbols with their individual corresponding phoneme.",
			lexicon,
		},
	}

	// ViewIndex  defines how our views should be laid out in the index tree
	ViewIndex = map[string][]string{
		"": {"transcriptor", "lexicon"},
	}
)

func transcript(w fyne.Window) fyne.CanvasObject {
	// #Init output row
	outputEntry := widget.NewEntry()
	outputEntry.Text = ""
	outputEntry.SetPlaceHolder("Click on a symbol to transcribe it")

	// init delete button
	deleteButton := widget.NewButtonWithIcon("", theme.ContentUndoIcon(), func() {
		if len(outputEntry.Text) == 0 {
			return
		}
		outputEntry.Text = outputEntry.Text[:len(outputEntry.Text)-1]
		outputEntry.Refresh()
	})
	// position delete button to the right of the outputEntry
	outputRow := container.NewBorder(nil, nil, nil, deleteButton, outputEntry)

	// #Init utils row (copy to clipboard and reset output)
	clipboardCopyButton := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		if len(outputEntry.Text) == 0 {
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Tunic Translator",
				Content: "Nothing to copy mate",
			})
			return
		}
		w.Clipboard().SetContent(outputEntry.Text)
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "Tunic Translator",
			Content: "Text has been copied to clipboard!",
		})
	})
	resetButton := widget.NewButtonWithIcon("", theme.ContentClearIcon(), func() {
		outputEntry.Text = ""
		outputEntry.Refresh()
	})
	// position the buttons at the right of the window
	utilsRow := container.NewBorder(nil, nil, nil, container.NewBorder(nil, nil, clipboardCopyButton, resetButton))

	// # Handle letters grids
	vowelsGrid := widget.NewGridWrap(
		func() int {
			return len(vowels)
		},
		func() fyne.CanvasObject {
			return container.NewGridWrap(fyne.NewSize(100, 110), canvas.NewImageFromFile("pkg/fyne/data/alphabet/1.png"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[0] = vowels[id].Img
		},
	)
	vowelsGrid.OnSelected = func(id widget.ListItemID) {
		// when an item is selected, update the sidepanel
		outputEntry.SetText(fmt.Sprintf("%s%s", outputEntry.Text, vowels[id].Rune))
		vowelsGrid.Unselect(id)
	}

	consonantsGrid := widget.NewGridWrap(
		func() int {
			return len(consonants)
		},
		func() fyne.CanvasObject {
			return container.NewGridWrap(fyne.NewSize(100, 110), canvas.NewImageFromFile("pkg/fyne/data/alphabet/20.png"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[0] = consonants[id].Img
		},
	)
	consonantsGrid.OnSelected = func(id widget.ListItemID) {
		// when an item is selected, update the sidepanel
		outputEntry.SetText(fmt.Sprintf("%s%s", outputEntry.Text, consonants[id].Rune))
		consonantsGrid.Unselect(id)
	}

	specialCharsGrid := widget.NewGridWrap(
		func() int {
			return len(specialChars)
		},
		func() fyne.CanvasObject {
			return container.NewGridWrap(fyne.NewSize(75, 75), canvas.NewImageFromFile("pkg/fyne/data/alphabet/1.png"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[0] = specialChars[id].Img
		},
	)
	specialCharsGrid.OnSelected = func(id widget.ListItemID) {
		// when an item is selected, update the sidepanel
		outputEntry.SetText(fmt.Sprintf("%s%s", outputEntry.Text, specialChars[id].Rune))
		specialCharsGrid.Unselect(id)
	}

	// # Grid layout
	gridLayout := container.NewGridWithColumns(2,
		container.NewBorder(widget.NewLabel("Vowels"), nil, nil, nil, vowelsGrid),
		container.NewBorder(widget.NewLabel("Consonants"), nil, nil, nil, consonantsGrid),
	)

	view := container.NewBorder(
		container.NewBorder(outputRow, utilsRow, nil, nil),
		nil,
		nil,
		nil,
		container.NewBorder(
			nil,
			container.NewPadded(
				specialCharsGrid,
			),
			nil,
			nil,
			container.NewPadded(gridLayout),
		),
	)

	return view
}

func lexicon(w fyne.Window) fyne.CanvasObject {
	combinedList := append(vowels, consonants...)

	objs := make([]fyne.CanvasObject, len(combinedList))
	for i := range combinedList {
		text := widget.NewLabel(combinedList[i].Rune)
		text.Alignment = fyne.TextAlignCenter
		img := combinedList[i].Img
		objs[i] = container.NewBorder(nil, text, nil, nil, img)
	}

	// warning: using container.NewGridWrap will not allow the user to resize the window
	// smaller than the size of the grid
	alphabetGrid := container.NewGridWrap(fyne.NewSize(75, 120), objs...)

	return alphabetGrid
}
