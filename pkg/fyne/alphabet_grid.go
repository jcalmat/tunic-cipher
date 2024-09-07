package fyne

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	twidget "github.com/jcalmat/tunic-cipher/pkg/fyne/widget"
)

func formatAlphabetGrid(w fyne.Window, itemBundles []alphabetItems, deleteFn func(i int), scroll bool) fyne.CanvasObject {
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
		// grid := widget.NewGridWrap(
		// 	func() int {
		// 		return len(bundle)
		// 	},
		// 	func() fyne.CanvasObject {
		// 		return container.NewGridWrap(fyne.NewSize(75, 120), currentGrid...)
		// 	},
		// 	func(id widget.ListItemID, item fyne.CanvasObject) {
		// 		// item.(*fyne.Container).Objects[0] = currentGrid[id]
		// 	})

		if deleteFn != nil {
			objs[i] = container.NewBorder(nil, nil, nil,
				container.NewBorder(nil, nil, widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
					w.Clipboard().SetContent(currentRunes.String())
					fyne.CurrentApp().SendNotification(&fyne.Notification{
						Title:   "Tunic Cipher",
						Content: "Text has been copied to clipboard!",
					})
				}),
					widget.NewButtonWithIcon("", theme.DeleteIcon(), func() { deleteFn(i) })),
				grid)
		} else {
			objs[i] = grid
		}
	}

	// warning: using container.NewGridWrap will not allow the user to resize the window
	// smaller than the size of the grid
	if scroll {
		return container.NewVScroll(container.NewVBox(objs...))
	}

	return container.NewVBox(objs...)
}
