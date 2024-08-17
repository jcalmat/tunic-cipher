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
	img := canvas.NewImageFromFile("pkg/fyne/data/alphabet/Pasted image 20240808232424.png")
	label := widget.NewLabel("Select An Item From The List")
	vbox := container.NewVBox(img, label)

	grid := widget.NewGridWrap(
		func() int {
			return len(alphabetMap)
		},
		func() fyne.CanvasObject {
			return container.NewGridWrap(fyne.NewSize(100, 100), canvas.NewImageFromFile("pkg/fyne/data/alphabet/Pasted image 20240808232424.png"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			fmt.Printf("id = %v, item = %+v\n", id, item.(*fyne.Container).Objects)
			item.(*fyne.Container).Objects[0].(*canvas.Image).Resource = canvas.NewImageFromFile(alphabetMap[id].Path).Resource
			item.(*fyne.Container).Objects[0].(*canvas.Image).Refresh()
			// item.(*fyne.Container).Objects[1].(*widget.Label).SetText(alphabetMap[id].Rune)
		},
	)
	grid.OnSelected = func(id widget.ListItemID) {
		label.SetText(alphabetMap[id].Rune)
		img = canvas.NewImageFromFile(alphabetMap[id].Path)
	}
	grid.OnUnselected = func(id widget.ListItemID) {
		label.SetText("Select An Item From The List")
		icon.SetResource(nil)
	}
	grid.Select(0)

	split := container.NewHSplit(grid, container.NewCenter(vbox))
	split.Offset = 0.6
	return split
}

func lexicon(w fyne.Window) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Lexicon"),
	)
}
