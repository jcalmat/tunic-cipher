package fyne

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jcalmat/tunic/pkg/fyne/storage"
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
	// outputEntry.Disable()
	outputEntry.SetPlaceHolder("Click on a symbol to transcribe it")

	currentQuery := make(alphabetItems, 0)

	// load saved queries
	savedQueries, err := storage.LoadQuery[alphabetItems]("saved_queries.json")
	if err != nil {
		fmt.Printf("failed to load saved queries: %v\n", err)
	}

	// init delete button
	deleteButton := widget.NewButtonWithIcon("", theme.ContentUndoIcon(), func() {
		if len(currentQuery) == 0 {
			return
		}

		currentQuery = currentQuery[:len(currentQuery)-1]
		outputEntry.Text = currentQuery.String()
		outputEntry.Refresh()
	})
	// position delete button to the right of the outputEntry
	outputRow := container.NewBorder(nil, container.NewPadded(widget.NewSeparator()), nil, deleteButton, outputEntry)

	// #Init utils row (copy to clipboard and reset output)
	clipboardCopyButton := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		if len(currentQuery) == 0 {
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Tunic Translator",
				Content: "Nothing to copy mate",
			})
			return
		}
		w.Clipboard().SetContent(currentQuery.String())
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "Tunic Translator",
			Content: "Text has been copied to clipboard!",
		})
	})
	resetButton := widget.NewButtonWithIcon("", theme.ContentClearIcon(), func() {
		currentQuery = make(alphabetItems, 0)
		outputEntry.Text = ""
		outputEntry.Refresh()
	})
	saveButton := widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {
		if len(currentQuery) == 0 {
			return
		}
		err := storage.SaveQuery[alphabetItems]("saved_queries.json", append([]alphabetItems{currentQuery}, savedQueries...))
		if err != nil {
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Tunic Translator",
				Content: "Failed to save: " + err.Error(),
			})
			return
		}
		savedQueries = append(savedQueries, currentQuery)
	})

	// position the buttons at the right of the window
	utilsRow := container.NewBorder(nil, container.NewPadded(widget.NewSeparator()), nil, container.NewBorder(nil, nil, container.NewBorder(nil, nil, saveButton, clipboardCopyButton), resetButton))

	// # Handle letters grids
	vowelsGrid := widget.NewGridWrap(
		func() int {
			return len(vowels)
		},
		func() fyne.CanvasObject {
			return container.NewGridWrap(fyne.NewSize(100, 110), canvas.NewImageFromFile("resources/alphabet/1.png"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[0] = vowels[id].Img
		},
	)
	vowelsGrid.OnSelected = func(id widget.ListItemID) {
		// when an item is selected, update the sidepanel
		currentQuery = append(currentQuery, vowels[id])
		outputEntry.SetText(currentQuery.String())
		vowelsGrid.Unselect(id)
	}

	consonantsGrid := widget.NewGridWrap(
		func() int {
			return len(consonants)
		},
		func() fyne.CanvasObject {
			return container.NewGridWrap(fyne.NewSize(100, 110), canvas.NewImageFromFile("resources/alphabet/20.png"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[0] = consonants[id].Img
		},
	)
	consonantsGrid.OnSelected = func(id widget.ListItemID) {
		// when an item is selected, update the sidepanel
		currentQuery = append(currentQuery, consonants[id])
		outputEntry.SetText(currentQuery.String())
		consonantsGrid.Unselect(id)
	}

	specialCharsGrid := widget.NewGridWrap(
		func() int {
			return len(specialChars)
		},
		func() fyne.CanvasObject {
			return container.NewGridWrap(fyne.NewSize(75, 75), canvas.NewImageFromFile("resources/alphabet/1.png"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[0] = specialChars[id].Img
		},
	)
	specialCharsGrid.OnSelected = func(id widget.ListItemID) {
		// when an item is selected, update the sidepanel
		currentQuery = append(currentQuery, specialChars[id])
		outputEntry.SetText(currentQuery.String())
		specialCharsGrid.Unselect(id)
	}

	// # Grid layout
	gridLayout := container.NewGridWithColumns(2,
		container.NewBorder(widget.NewLabel("Vowels"), nil, nil, nil, vowelsGrid),
		container.NewBorder(widget.NewLabel("Consonants"), nil, nil, nil, consonantsGrid),
	)

	tabs := container.NewAppTabs(
		container.NewTabItem("Glyphes", container.NewBorder(
			nil,
			container.NewPadded(
				layout.NewSpacer(),
				specialCharsGrid,
			),
			nil,
			nil,
			container.NewPadded(gridLayout),
		)),
		container.NewTabItem("Saved queries", formatAlphabetGrid(append([]alphabetItems{currentQuery}, savedQueries...))),
	)

	tabs.OnSelected = func(item *container.TabItem) {
		if item.Text == "Saved queries" {
			if len(currentQuery) == 0 {
				item.Content = formatAlphabetGrid(savedQueries)
				return
			}
			item.Content = formatAlphabetGrid(append([]alphabetItems{currentQuery}, savedQueries...))
		}
	}

	view := container.NewBorder(
		container.NewBorder(outputRow, utilsRow, nil, nil),
		nil,
		nil,
		nil,
		tabs,
	)

	return view
}

func lexicon(w fyne.Window) fyne.CanvasObject {
	return formatAlphabetGrid([]alphabetItems{vowels, consonants})
}

func formatAlphabetGrid(itemBundles []alphabetItems) fyne.CanvasObject {
	objs := make([]fyne.CanvasObject, len(itemBundles))
	for i, bundle := range itemBundles {
		bundle := bundle
		currentGrid := make([]fyne.CanvasObject, len(bundle))
		for j, item := range bundle {
			item := item
			text := widget.NewLabel(item.Rune)
			text.Alignment = fyne.TextAlignCenter
			newImg := canvas.NewImageFromImage(item.Img.Image)
			currentGrid[j] = container.NewBorder(nil, text, nil, nil, newImg)
			_ = item
		}
		objs[i] = container.NewGridWrap(fyne.NewSize(75, 120), currentGrid...)
	}

	// warning: using container.NewGridWrap will not allow the user to resize the window
	// smaller than the size of the grid
	return container.NewVScroll(container.NewVBox(objs...))
}
