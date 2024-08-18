package fyne

import (
	"fyne.io/fyne/v2/canvas"
)

type alphabetItemType int

const (
	Vowel alphabetItemType = iota
	Consonant
	Special
)

type alphabetItem struct {
	Img  *canvas.Image
	Rune string
	Type alphabetItemType
}

var vowels = []alphabetItem{
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/1.png"), Type: Vowel, Rune: "ɔː"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/2.png"), Type: Vowel, Rune: "ɑː"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/3.png"), Type: Vowel, Rune: "ɪ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/4.png"), Type: Vowel, Rune: "e"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/5.png"), Type: Vowel, Rune: "ʊ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/6.png"), Type: Vowel, Rune: "ʌ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/7.png"), Type: Vowel, Rune: "i:"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/8.png"), Type: Vowel, Rune: "uː"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/9.png"), Type: Vowel, Rune: "ə"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/10.png"), Type: Vowel, Rune: "ɔː"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/11.png"), Type: Vowel, Rune: "ɑː"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/12.png"), Type: Vowel, Rune: "ɪə"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/13.png"), Type: Vowel, Rune: "eɪ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/14.png"), Type: Vowel, Rune: "aɪ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/15.png"), Type: Vowel, Rune: ""},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/16.png"), Type: Vowel, Rune: "aʊ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/17.png"), Type: Vowel, Rune: "əʊ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/18.png"), Type: Vowel, Rune: "eə"},
}
var consonants = []alphabetItem{
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/19.png"), Type: Consonant, Rune: "m"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/20.png"), Type: Consonant, Rune: "n"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/21.png"), Type: Consonant, Rune: "ŋ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/22.png"), Type: Consonant, Rune: "p"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/23.png"), Type: Consonant, Rune: "b"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/24.png"), Type: Consonant, Rune: "t"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/25.png"), Type: Consonant, Rune: "d"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/26.png"), Type: Consonant, Rune: "k"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/27.png"), Type: Consonant, Rune: "g"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/28.png"), Type: Consonant, Rune: "ʤ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/29.png"), Type: Consonant, Rune: "tʃ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/30.png"), Type: Consonant, Rune: "f"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/31.png"), Type: Consonant, Rune: "v"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/32.png"), Type: Consonant, Rune: "θ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/33.png"), Type: Consonant, Rune: "ð"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/34.png"), Type: Consonant, Rune: "s"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/35.png"), Type: Consonant, Rune: "z"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/36.png"), Type: Consonant, Rune: "ʃ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/37.png"), Type: Consonant, Rune: ""},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/38.png"), Type: Consonant, Rune: "ˈh"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/39.png"), Type: Consonant, Rune: "r"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/40.png"), Type: Consonant, Rune: "j"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/41.png"), Type: Consonant, Rune: "w"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/42.png"), Type: Consonant, Rune: "l"},
}

var specialChars = []alphabetItem{
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/punctuation/space.png"), Type: Special, Rune: " "},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/punctuation/exclamation mark.png"), Type: Special, Rune: "!"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/punctuation/question mark.png"), Type: Special, Rune: "?"},
}
