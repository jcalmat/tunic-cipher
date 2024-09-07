package fyne

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// formatAlphabetGrid creates a grid of alphabet items with extra parameters.
// The optional deleteFn parameter is a function that will be called when the delete button is pressed.
// The optional scroll parameter is a boolean that will determine if the grid should be scrollable.
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
			newImg := canvas.NewImageFromImage(item.Img.Image)
			currentGrid[j] = container.NewBorder(nil, text, nil, nil, newImg)
			currentRunes.WriteString(item.Rune)
		}
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
				container.NewGridWrap(fyne.NewSize(75, 120), currentGrid...))
		} else {
			objs[i] = container.NewGridWrap(fyne.NewSize(75, 120), currentGrid...)
		}
	}

	// warning: using container.NewGridWrap will not allow the user to resize the window
	// smaller than the size of the grid
	if scroll {
		return container.NewVScroll(container.NewVBox(objs...))
	}

	return container.NewVBox(objs...)
}
