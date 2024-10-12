package fyne

import (
	"errors"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jcalmat/tunic-cipher/pkg/fyne/storage"
)

// transcript returns a CanvasObject that displays the transcript view
func transcript(w fyne.Window) fyne.CanvasObject {
	// #Init output row
	outputEntry := widget.NewEntry()
	outputEntry.Text = ""
	outputEntry.SetPlaceHolder("Click on a symbol to transcribe it")

	currentQuery := make(alphabetItems, 0)

	// get the current alphabet, either the default one or the one saved by the user
	// it will be used all along the transcript view
	currentAlphabet := currentAlphabet()

	// load saved queries
	savedQueries, err := storage.Load[alphabetItems]("saved_queries.tnc")
	if err != nil {
		if !errors.Is(err, storage.ErrNotFound) {
			fmt.Printf("failed to load saved queries: %v\n", err)
		}
	}
	savedQueries = currentAlphabet.Normalize(savedQueries)

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
				Title:   "Tunic Cipher",
				Content: "Nothing to copy mate",
			})
			return
		}
		w.Clipboard().SetContent(currentQuery.String())
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "Tunic Cipher",
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
		err := storage.Save("saved_queries.tnc", append([]alphabetItems{currentQuery}, savedQueries...))
		if err != nil {
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Tunic Cipher",
				Content: "Failed to save: " + err.Error(),
			})
			return
		}
		savedQueries = append([]alphabetItems{currentQuery}, savedQueries...)
	})

	// position the buttons at the right of the window
	utilsRow := container.NewBorder(nil, container.NewPadded(widget.NewSeparator()), nil, container.NewBorder(nil, nil, container.NewBorder(nil, nil, saveButton, clipboardCopyButton), resetButton))

	// # Handle letters grids
	vowels := currentAlphabet.ByType(Vowel)
	vowelsGrid := widget.NewGridWrap(
		func() int {
			return len(vowels)
		},
		func() fyne.CanvasObject {
			return container.NewGridWrap(fyne.NewSize(100, 110), &canvas.Image{})
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

	consonants := currentAlphabet.ByType(Consonant)
	consonantsGrid := widget.NewGridWrap(
		func() int {
			return len(consonants)
		},
		func() fyne.CanvasObject {
			return container.NewGridWrap(fyne.NewSize(100, 110), &canvas.Image{})
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
			return container.NewGridWrap(fyne.NewSize(75, 75), &canvas.Image{})
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
		container.NewTabItem("Glyphs", container.NewBorder(
			nil,
			container.NewPadded(
				layout.NewSpacer(),
				specialCharsGrid,
			),
			nil,
			nil,
			container.NewPadded(gridLayout),
		)),
		container.NewTabItem("Queries",
			container.NewBorder(formatAlphabetGrid(w, []alphabetItems{currentQuery}, nil, false), nil, nil, formatAlphabetGrid(w, savedQueries, nil, true))),
	)

	deleteQueryFn := func(i int) {
		savedQueries = append(savedQueries[:i], savedQueries[i+1:]...)
		err := storage.Save("saved_queries.tnc", savedQueries)
		if err != nil {
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Tunic Cipher",
				Content: "Failed to delete: " + err.Error(),
			})
		}
		// unfortunately, refreshing the current tab doesn't work as expected
		// and since we're in the deleteQueryFn, we can't fall formatAlphabetGrid and pass the delete function.
		// to manually refresh the tab, we need to select another tab and then reselect the current tab
		tabs.SelectIndex(0)
		tabs.SelectIndex(1)
	}

	tabs.OnSelected = func(item *container.TabItem) {
		if item.Text == "Queries" {
			if len(currentQuery) == 0 {
				item.Content = container.NewBorder(widget.NewLabel("Saved queries"), nil, nil, nil, formatAlphabetGrid(w, savedQueries, deleteQueryFn, true))
				return
			}
			item.Content = container.NewBorder(
				container.NewBorder(widget.NewLabel("Current query"), nil, nil, nil, formatAlphabetGrid(w, []alphabetItems{currentQuery}, nil, false)),
				nil, nil, nil,
				container.NewBorder(widget.NewLabel("Saved queries"), nil, nil, nil, formatAlphabetGrid(w, savedQueries, deleteQueryFn, true)))
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
