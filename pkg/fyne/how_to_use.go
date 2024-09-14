package fyne

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// howToUse returns a CanvasObject that displays the howToUse view
func howToUse(w fyne.Window) fyne.CanvasObject {
	logo := canvas.NewImageFromResource(resourceLogobigPng)
	logo.FillMode = canvas.ImageFillContain
	logo.SetMinSize(fyne.NewSize(256, 256))

	footer := container.NewHBox(
		layout.NewSpacer(),
		widget.NewHyperlink("github", parseURL("https://github.com/jcalmat/tunic-cipher")),
		layout.NewSpacer(),
	)

	text1 := widget.NewRichTextFromMarkdown(`
# Tunic Cipher

Tunic Cipher is a tool to help you transcribe and decipher texts written in the Tunic script.

<br/><br/><br/><br/><br/><br/><br/><br/>

## How to read a Tunic text

The Tunic Language is written from left to right, and each glyph represents a single sound that ultimately corresponds to phonemes that can be translated to english. 
`)
	text1.Wrapping = fyne.TextWrapWord

	img1 := canvas.NewImageFromResource(resourceTunictextexamplePng)
	img1.FillMode = canvas.ImageFillContain
	img1.SetMinSize(fyne.NewSize(256, 256))

	text2 := widget.NewRichTextFromMarkdown(`Each glyph contains 2 distinct parts: the consonant and the vowel. The consonant is represented by the external part of the glyph, and the vowel is represented by the inside part of the glyph.
A glyph is read from the outside to the inside, starting with the consonant and ending with the vowel.`)
	text2.Wrapping = fyne.TextWrapWord

	img2 := canvas.NewImageFromResource(resourceTunictextexampleoutsideinsidePng)
	img2.FillMode = canvas.ImageFillContain
	img2.SetMinSize(fyne.NewSize(256, 256))

	text3 := widget.NewRichTextFromMarkdown("As seen in the previous screen, sometimes a glyph is separated by an horizontal line in its center. It seems to be purely aesthetic and does not change the reading of the glyph.")
	text3.Wrapping = fyne.TextWrapWord

	text4 := widget.NewRichTextFromMarkdown("Finally, at the bottom of the second glyph, you can find a small circle. This circle reverses the reading of the glyph. When you encounter a circle, you should read the glyph from the inside to the outside, starting with the vowel and ending with the consonant.")
	text4.Wrapping = fyne.TextWrapWord

	text5 := widget.NewRichTextFromMarkdown(`
<br/><br/><br/><br/><br/><br/><br/><br/>
---
## How to use the Tunic Cipher tool

Now that you know how to read a Tunic text, let's see how to use the Tunic Cipher tool.

The Tunic Cipher tool is divided into 3 main sections:
- the "how to use" view
- the transcript view
- the lexicon view.

<br/><br/><br/><br/><br/><br/><br/><br/>
### The transcriptor

The transcriptor is the main view of the Tunic Cipher tool, and it allows you to transcribe and decipher Tunic texts.

It is divided into 2 tabs: the glyphs tab and the queries tab.

The glyphs tab provides you with a list of glyphs, split by inside and outside parts, that you can click on to transcribe them, and a text area where you can see the result of your transcription.`)
	text5.Wrapping = fyne.TextWrapWord

	img5 := canvas.NewImageFromResource(resourceTranscriptorPng)
	img5.FillMode = canvas.ImageFillContain
	img5.SetMinSize(fyne.NewSize(256, 256))

	text6 := widget.NewRichTextFromMarkdown("The queries tab provides you with a list of saved queries alonside your current query.")
	text6.Wrapping = fyne.TextWrapWord

	img6 := canvas.NewImageFromResource(resourceQueriesPng)
	img6.FillMode = canvas.ImageFillContain
	img6.SetMinSize(fyne.NewSize(256, 256))

	text7 := widget.NewRichTextFromMarkdown(`
<br/><br/><br/><br/><br/><br/><br/><br/>
### The lexicon

The lexicon view provides you with a list of all the glyphs of the Tunic alphabet, and allows you to modify them.
Modifying a glyph will update the glyph in the transcriptor view and will update the saved queries.`)
	text7.Wrapping = fyne.TextWrapWord

	img7 := canvas.NewImageFromResource(resourceLexiconPng)
	img7.FillMode = canvas.ImageFillContain
	img7.SetMinSize(fyne.NewSize(256, 256))

	content := container.NewVBox(
		widget.NewLabelWithStyle("\n\nWelcome to the Tunic (de)Cipher tool!", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		logo,
		text1,
		layout.NewSpacer(),
		img1,
		layout.NewSpacer(),
		text2,
		layout.NewSpacer(),
		img2,
		layout.NewSpacer(),
		text3,
		text4,
		text5,
		layout.NewSpacer(),
		img5,
		layout.NewSpacer(),
		text6,
		layout.NewSpacer(),
		img6,
		layout.NewSpacer(),
		text7,
		layout.NewSpacer(),
		img7,
		layout.NewSpacer(),
		widget.NewLabelWithStyle("\nThank you for using Tunic Cipher!\n", fyne.TextAlignCenter, fyne.TextStyle{Italic: true}),
	)
	scroll := container.NewScroll(content)

	return container.NewBorder(nil, footer, nil, nil, scroll)
}

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}
