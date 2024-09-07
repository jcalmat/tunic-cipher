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

func transcript(w fyne.Window) fyne.CanvasObject {
	// #Init output row
	outputEntry := widget.NewEntry()
	outputEntry.Text = ""
	// outputEntry.Disable()
	outputEntry.SetPlaceHolder("Click on a symbol to transcribe it")

	currentQuery := make(alphabetItems, 0)

	// load saved queries
	savedQueries, err := storage.Load[alphabetItems]("saved_queries.json")
	if err != nil {
		if !errors.Is(err, storage.ErrNotFound) {
			fmt.Printf("failed to load saved queries: %v\n", err)
		}
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
		err := storage.Save("saved_queries.json", append([]alphabetItems{currentQuery}, savedQueries...))
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
	vowelsGrid := widget.NewGridWrap(
		func() int {
			return len(defaultVowels)
		},
		func() fyne.CanvasObject {
			return container.NewGridWrap(fyne.NewSize(100, 110), canvas.NewImageFromFile("resources/alphabet/1.png"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[0] = defaultVowels[id].Img
		},
	)
	vowelsGrid.OnSelected = func(id widget.ListItemID) {
		// when an item is selected, update the sidepanel
		currentQuery = append(currentQuery, defaultVowels[id])
		outputEntry.SetText(currentQuery.String())
		vowelsGrid.Unselect(id)
	}

	consonantsGrid := widget.NewGridWrap(
		func() int {
			return len(defaultConsonants)
		},
		func() fyne.CanvasObject {
			return container.NewGridWrap(fyne.NewSize(100, 110), canvas.NewImageFromFile("resources/alphabet/20.png"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[0] = defaultConsonants[id].Img
		},
	)
	consonantsGrid.OnSelected = func(id widget.ListItemID) {
		// when an item is selected, update the sidepanel
		currentQuery = append(currentQuery, defaultConsonants[id])
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
		container.NewTabItem("Queries",
			container.NewBorder(formatAlphabetGrid(w, []alphabetItems{currentQuery}, nil, false), nil, nil, formatAlphabetGrid(w, savedQueries, nil, true))),
	)

	deleteQueryFn := func(i int) {
		savedQueries = append(savedQueries[:i], savedQueries[i+1:]...)
		err := storage.Save("saved_queries.json", savedQueries)
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
