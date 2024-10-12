package fyne

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jcalmat/tunic-cipher/pkg/fyne/storage"
	twidget "github.com/jcalmat/tunic-cipher/pkg/fyne/widget"
)

// lexicon returns a CanvasObject that displays the lexicon of the alphabet
func lexicon(w fyne.Window) fyne.CanvasObject {
	alphabet, err := storage.Load[alphabetItems]("alphabet.tnc")
	if err != nil {
		alphabet = emptyAlphabet()
	}

	objs := make([]fyne.CanvasObject, len(alphabet))

	for i, bundle := range alphabet {
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
				entry := widget.NewEntry()
				entry.SetPlaceHolder(item.Rune)
				d := dialog.NewCustomConfirm("Enter the new glyph",
					"Confirm",
					"Abort",
					container.NewBorder(nil, layout.NewSpacer(), nil, nil, entry),
					func(b bool) {
						if b {
							// update the glyph and save it
							if entry.Text != "" && item.Rune != entry.Text {
								item.Rune = entry.Text
								bundle.Update(item.ID, entry.Text)
								alphabet[i] = bundle
								storage.Save("alphabet.tnc", alphabet)
								text.SetText(item.Rune)
								text.Refresh()
							}
						}
					}, w)
				d.Show()
				// auto focus the entry
				w.Canvas().Focus(entry)
			})
			currentGrid[j] = container.NewBorder(nil, text, nil, nil, w)
			currentRunes.WriteString(item.Rune)
		}
		grid := container.NewGridWrap(fyne.NewSize(75, 120), currentGrid...)

		objs[i] = grid
	}

	resetBtn := widget.NewButtonWithIcon("Reset to default", theme.DeleteIcon(), func() {
		// ask for confirmation
		dialog.ShowConfirm("Reset to default", "Are you sure you want to reset the alphabet to the default one?", func(b bool) {
			if b {
				// reset the alphabet to the default one
				storage.Save("alphabet.tnc", emptyAlphabet())
				// refresh the text of the labels
				for i, bundle := range emptyAlphabet() {
					for j, item := range bundle {
						objs[i].(*fyne.Container).Objects[j].(*fyne.Container).Objects[1].(*widget.Label).SetText(item.Rune)
						objs[i].(*fyne.Container).Objects[j].(*fyne.Container).Refresh()

					}
				}
			}
		}, w)
	})

	loadFinalAlphabetBtn := widget.NewButtonWithIcon("Load solution", theme.UploadIcon(), func() {
		// ask for confirmation
		dialog.ShowConfirm("Load solution", "Are you sure you want to load the final solution? It might spoil you!", func(b bool) {
			if b {
				// reset the alphabet to the default one
				storage.Save("alphabet.tnc", definitiveAlphabet())
				// refresh the text of the labels
				for i, bundle := range definitiveAlphabet() {
					for j, item := range bundle {
						objs[i].(*fyne.Container).Objects[j].(*fyne.Container).Objects[1].(*widget.Label).SetText(item.Rune)
						objs[i].(*fyne.Container).Objects[j].(*fyne.Container).Refresh()
					}
				}
			}
		}, w)
	})

	return container.NewBorder(nil, container.NewBorder(nil, nil, nil, container.NewBorder(nil, nil, loadFinalAlphabetBtn, resetBtn)), nil, nil, container.NewVScroll(container.NewVBox(objs...)))
}
