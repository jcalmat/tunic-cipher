package fyne

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
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
	icon := widget.NewIcon(nil)
	spImg := canvas.NewImageFromFile("pkg/fyne/data/alphabet/Pasted image 20240808232424.png")
	spImg.SetMinSize(fyne.NewSize(100, 100))
	spLabel := widget.NewLabel("Select An Item From The List")
	spLabel.Alignment = fyne.TextAlignCenter
	sidePanel := container.NewVBox(spImg, spLabel)

	grid := widget.NewGridWrap(
		func() int {
			return len(alphabetMap)
		},
		func() fyne.CanvasObject {
			return container.NewGridWrap(fyne.NewSize(100, 100), canvas.NewImageFromFile("pkg/fyne/data/alphabet/Pasted image 20240808232424.png"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[0] = alphabetMap[id].Img
		},
	)
	grid.OnSelected = func(id widget.ListItemID) {
		// when an item is selected, update the sidepanel
		fmt.Printf("Selected: %d\n", id)
		spLabel.SetText(alphabetMap[id].Rune)
		// spImg = alphabetMap[3].Img

		spImg = alphabetMap[id].Img
		spImg.SetMinSize(fyne.NewSize(100, 100))
		sidePanel.Objects[0] = spImg

	}
	grid.OnUnselected = func(id widget.ListItemID) {
		spLabel.SetText("Select An Item From The List")
		icon.SetResource(nil)
	}
	grid.Select(15)

	split := container.NewHSplit(grid, container.NewCenter(sidePanel))
	split.Offset = 0.6
	return split
}

func lexicon(w fyne.Window) fyne.CanvasObject {
	icon := widget.NewIcon(nil)
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
	grid.OnUnselected = func(id widget.ListItemID) {
		spLabel.SetText("Select An Item From The List")
		icon.SetResource(nil)
	}
	grid.Select(15)

	split := container.NewHSplit(grid, container.NewCenter(sidePanel))
	split.Offset = 0.6
	return split
}
