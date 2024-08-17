package fyne

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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
	spEntry := widget.NewEntry()
	spEntry.Text = ""
	spEntry.SetPlaceHolder("Click on a symbol to transcribe it")
	// spSymbols := []canvas.Image{}
	clipboardCopyButton := widget.NewButton("Copy me!", func() {
		if len(spEntry.Text) == 0 {
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Tunic Translator",
				Content: "Nothing to copy mate",
			})
			return
		}
		w.Clipboard().SetContent(spEntry.Text)
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "Tunic Translator",
			Content: "Text has been copied to clipboard!",
		})
	})

	sidePanel := container.NewVBox(
		layout.NewSpacer(),
		container.NewGridWithColumns(3,
			layout.NewSpacer(),
			container.NewVBox(spEntry, clipboardCopyButton),
			layout.NewSpacer()),
		layout.NewSpacer(),
	)

	grid := widget.NewGridWrap(
		func() int {
			return len(alphabetMap)
		},
		func() fyne.CanvasObject {
			return container.NewGridWrap(fyne.NewSize(100, 108), canvas.NewImageFromFile("pkg/fyne/data/alphabet/Pasted image 20240808232424.png"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[0] = alphabetMap[id].Img
		},
	)
	grid.OnSelected = func(id widget.ListItemID) {
		// when an item is selected, update the sidepanel
		spEntry.Text = fmt.Sprintf("%s%s", spEntry.Text, alphabetMap[id].Rune)
		spEntry.Refresh()
	}

	split := container.NewVSplit(grid, sidePanel)
	split.Offset = 0.8
	return split
}

func lexicon(w fyne.Window) fyne.CanvasObject {
	spImg := canvas.NewImageFromFile("pkg/fyne/data/alphabet/Pasted image 20240808232424.png")
	spImg.SetMinSize(fyne.NewSize(100, 108))
	spLabel := widget.NewLabel("Select An Item From The List")
	spLabel.Alignment = fyne.TextAlignCenter
	sidePanel := container.NewVBox(spImg, spLabel)

	grid := widget.NewGridWrap(
		func() int {
			return len(alphabetMap)
		},
		func() fyne.CanvasObject {
			return container.NewGridWrap(fyne.NewSize(100, 108), canvas.NewImageFromFile("pkg/fyne/data/alphabet/Pasted image 20240808232424.png"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[0] = alphabetMap[id].Img
		},
	)
	grid.OnSelected = func(id widget.ListItemID) {
		// when an item is selected, update the sidepanel
		spLabel.SetText(alphabetMap[id].Rune)
		// spImg = alphabetMap[3].Img

		spImg = alphabetMap[id].Img
		spImg.SetMinSize(fyne.NewSize(100, 108))
		sidePanel.Objects[0] = spImg

	}

	grid.Select(15)

	split := container.NewHSplit(grid, container.NewCenter(sidePanel))
	split.Offset = 0.6
	return split
}
