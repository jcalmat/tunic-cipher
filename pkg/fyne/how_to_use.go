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

The **Tunic Cipher** is a tool designed to help you transcribe and decipher texts written in the Tunic script.

<br/><br/><br/><br/><br/><br/><br/><br/>

## Reading Tunic Texts

The Tunic language is written from left to right, with each glyph representing a sound that corresponds to phonemes translatable into English.
`)
	text1.Wrapping = fyne.TextWrapWord

	img1 := canvas.NewImageFromResource(resourceTunictextexamplePng)
	img1.FillMode = canvas.ImageFillContain
	img1.SetMinSize(fyne.NewSize(450, 450))

	text2 := widget.NewRichTextFromMarkdown(`Each glyph is composed of two distinct elements: the consonant and the vowel. The consonant is depicted by the outer part of the glyph, while the vowel is represented by the inner part. When reading a glyph, start from the outside with the consonant, and move inward to the vowel.`)
	text2.Wrapping = fyne.TextWrapWord

	img2 := canvas.NewImageFromResource(resourceTunictextexampleoutsideinsidePng)
	img2.FillMode = canvas.ImageFillContain
	img2.SetMinSize(fyne.NewSize(450, 450))

	text3 := widget.NewRichTextFromMarkdown("As you've seen, some glyphs have a horizontal line through the middle. This line is purely decorative and does not alter the reading of the glyph.")
	text3.Wrapping = fyne.TextWrapWord

	text4 := widget.NewRichTextFromMarkdown("Occasionally, you’ll encounter a small circle at the bottom of a glyph. This circle indicates that the reading direction is reversed. In these cases, start with the vowel on the inside and end with the consonant on the outside.")
	text4.Wrapping = fyne.TextWrapWord

	text5 := widget.NewRichTextFromMarkdown(`
<br/><br/><br/><br/><br/><br/><br/><br/>
---
## Using the Tunic Cipher Tool

Now that you understand how to read Tunic text, let's explore how to use the Tunic Cipher tool.

The tool is divided into three main sections:
- The "How to Use" view
- The Transcript view
- The Lexicon view

<br/><br/><br/><br/><br/><br/><br/><br/>
### The Transcriptor

The **Transcriptor** is the main feature of the Tunic Cipher tool, where you can transcribe and decipher Tunic texts. It has two tabs:

- **Glyphs Tab**: This tab offers a list of glyphs separated by their inner and outer parts, which you can click to transcribe. Your transcription will appear in the text area provided.`)
	text5.Wrapping = fyne.TextWrapWord

	img5 := canvas.NewImageFromResource(resourceTranscriptorPng)
	img5.FillMode = canvas.ImageFillContain
	img5.SetMinSize(fyne.NewSize(450, 450))

	text6 := widget.NewRichTextFromMarkdown("- **Queries Tab**: This tab allows you to manage saved queries, alongside your current transcription query.")
	text6.Wrapping = fyne.TextWrapWord

	img6 := canvas.NewImageFromResource(resourceQueriesPng)
	img6.FillMode = canvas.ImageFillContain
	img6.SetMinSize(fyne.NewSize(450, 450))

	text7 := widget.NewRichTextFromMarkdown(`
<br/><br/><br/><br/><br/><br/><br/><br/>
### The Lexicon

The **Lexicon** view contains a comprehensive list of all the glyphs in the Tunic alphabet. You can modify these glyphs, and any changes made will automatically update both the Transcriptor view and your saved queries.`)
	text7.Wrapping = fyne.TextWrapWord

	img7 := canvas.NewImageFromResource(resourceLexiconPng)
	img7.FillMode = canvas.ImageFillContain
	img7.SetMinSize(fyne.NewSize(450, 450))

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
		widget.NewLabelWithStyle("\nEnjoy!\n", fyne.TextAlignCenter, fyne.TextStyle{Italic: true}),
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
