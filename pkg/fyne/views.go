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
		"": {"welcome", "transcriptor", "lexicon"},
	}
)

func transcript(w fyne.Window) fyne.CanvasObject {
	// #Init output row
	outputEntry := widget.NewEntry()
	outputEntry.Text = ""
	outputEntry.SetPlaceHolder("Click on a symbol to transcribe it")

	// init delete button
	deleteButton := widget.NewButtonWithIcon("", theme.ContentClearIcon(), func() {
		if len(outputEntry.Text) == 0 {
			return
		}
		outputEntry.Text = outputEntry.Text[:len(outputEntry.Text)-1]
		outputEntry.Refresh()
	})
	// position delete button to the right of the outputEntry
	outputRow := container.NewBorder(nil, nil, nil, deleteButton, outputEntry)

	// #Init utils row (copy to clipboard and reset output)
	clipboardCopyButton := widget.NewButton("Copy me!", func() {
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
	resetButton := widget.NewButton("Reset", func() {
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
			return container.NewGridWrap(fyne.NewSize(100, 100), canvas.NewImageFromFile("pkg/fyne/data/alphabet/1.png"))
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
			return container.NewGridWrap(fyne.NewSize(100, 100), canvas.NewImageFromFile("pkg/fyne/data/alphabet/20.png"))
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

	// # Grid layout
	gridLayout := container.NewGridWithColumns(2,
		container.NewBorder(widget.NewLabel("Vowels"), nil, nil, nil, vowelsGrid),
		container.NewBorder(widget.NewLabel("Consonants"), nil, nil, nil, consonantsGrid),
	)

	view := container.NewBorder(container.NewBorder(outputRow, utilsRow, nil, nil), nil, nil, nil, gridLayout)

	return view
}

func lexicon(w fyne.Window) fyne.CanvasObject {
	// spImg := canvas.NewImageFromFile("pkg/fyne/data/alphabet/Pasted image 20240808232424.png")
	// spImg.SetMinSize(fyne.NewSize(100, 108))
	// spLabel := widget.NewLabel("Select An Item From The List")
	// spLabel.Alignment = fyne.TextAlignCenter
	// sidePanel := container.NewVBox(spImg, spLabel)

	// grid := widget.NewGridWrap(
	// 	func() int {
	// 		return len(alphabetMap)
	// 	},
	// 	func() fyne.CanvasObject {
	// 		return container.NewGridWrap(fyne.NewSize(100, 108), canvas.NewImageFromFile("pkg/fyne/data/alphabet/Pasted image 20240808232424.png"))
	// 	},
	// 	func(id widget.ListItemID, item fyne.CanvasObject) {
	// 		item.(*fyne.Container).Objects[0] = alphabetMap[id].Img
	// 	},
	// )
	// grid.OnSelected = func(id widget.ListItemID) {
	// 	// when an item is selected, update the sidepanel
	// 	spLabel.SetText(alphabetMap[id].Rune)
	// 	// spImg = alphabetMap[3].Img

	// 	spImg = alphabetMap[id].Img
	// 	spImg.SetMinSize(fyne.NewSize(100, 108))
	// 	sidePanel.Objects[0] = spImg

	// }

	// grid.Select(15)

	// split := container.NewHSplit(grid, container.NewCenter(sidePanel))
	// split.Offset = 0.6
	// return split
	return widget.NewLabel("Lexicon")
}
