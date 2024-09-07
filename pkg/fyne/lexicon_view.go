package fyne

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	twidget "github.com/jcalmat/tunic-cipher/pkg/fyne/widget"
)

func lexicon(w fyne.Window) fyne.CanvasObject {
	// return formatAlphabetGrid(w, []alphabetItems{defaultVowels, defaultConsonants}, nil, true)
	itemBundles := []alphabetItems{defaultVowels, defaultConsonants}
	objs := make([]fyne.CanvasObject, len(itemBundles))
	for i, bundle := range itemBundles {
		bundle := bundle
		currentGrid := make([]fyne.CanvasObject, len(bundle))
		currentRunes := strings.Builder{}
		for j, item := range bundle {
			item := item
			text := widget.NewLabel(item.Rune)
			text.Alignment = fyne.TextAlignCenter
			img := item.Img
			img.SetMinSize(fyne.NewSize(75, 75))
			w := twidget.NewClickableCanvasWithImage(img, func() {
				fyne.CurrentApp().SendNotification(&fyne.Notification{
					Title:   "Tunic Cipher",
					Content: "Clicked",
				})
				win1 := container.NewInnerWindow("Inner Demo", container.NewVBox(
					img,
					widget.NewEntry()),
				)

				multi := container.NewMultipleWindows()
				multi.Windows = []*container.InnerWindow{win1}
			})
			currentGrid[j] = container.NewBorder(nil, text, nil, nil, w)
			currentRunes.WriteString(item.Rune)
		}
		grid := container.NewGridWrap(fyne.NewSize(75, 120), currentGrid...)

		objs[i] = grid
	}

	return container.NewVScroll(container.NewVBox(objs...))
}
