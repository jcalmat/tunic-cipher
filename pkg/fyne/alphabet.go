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
	{Img: canvas.NewImageFromFile("resources/alphabet/1.png"), Type: Vowel, Rune: "ɔː"},
	{Img: canvas.NewImageFromFile("resources/alphabet/2.png"), Type: Vowel, Rune: "ɑː"},
	{Img: canvas.NewImageFromFile("resources/alphabet/3.png"), Type: Vowel, Rune: "ɪ"},
	{Img: canvas.NewImageFromFile("resources/alphabet/4.png"), Type: Vowel, Rune: "e"},
	{Img: canvas.NewImageFromFile("resources/alphabet/5.png"), Type: Vowel, Rune: "ʊ"},
	{Img: canvas.NewImageFromFile("resources/alphabet/6.png"), Type: Vowel, Rune: "ʌ"},
	{Img: canvas.NewImageFromFile("resources/alphabet/7.png"), Type: Vowel, Rune: "i:"},
	{Img: canvas.NewImageFromFile("resources/alphabet/8.png"), Type: Vowel, Rune: "uː"},
	{Img: canvas.NewImageFromFile("resources/alphabet/9.png"), Type: Vowel, Rune: "ə"},
	{Img: canvas.NewImageFromFile("resources/alphabet/10.png"), Type: Vowel, Rune: "ɔː"},
	{Img: canvas.NewImageFromFile("resources/alphabet/11.png"), Type: Vowel, Rune: "ɑː"},
	{Img: canvas.NewImageFromFile("resources/alphabet/12.png"), Type: Vowel, Rune: "ɪə"},
	{Img: canvas.NewImageFromFile("resources/alphabet/13.png"), Type: Vowel, Rune: "eɪ"},
	{Img: canvas.NewImageFromFile("resources/alphabet/14.png"), Type: Vowel, Rune: "aɪ"},
	{Img: canvas.NewImageFromFile("resources/alphabet/15.png"), Type: Vowel, Rune: ""},
	{Img: canvas.NewImageFromFile("resources/alphabet/16.png"), Type: Vowel, Rune: "aʊ"},
	{Img: canvas.NewImageFromFile("resources/alphabet/17.png"), Type: Vowel, Rune: "əʊ"},
	{Img: canvas.NewImageFromFile("resources/alphabet/18.png"), Type: Vowel, Rune: "eə"},
}
var consonants = []alphabetItem{
	{Img: canvas.NewImageFromFile("resources/alphabet/19.png"), Type: Consonant, Rune: "m"},
	{Img: canvas.NewImageFromFile("resources/alphabet/20.png"), Type: Consonant, Rune: "n"},
	{Img: canvas.NewImageFromFile("resources/alphabet/21.png"), Type: Consonant, Rune: "ŋ"},
	{Img: canvas.NewImageFromFile("resources/alphabet/22.png"), Type: Consonant, Rune: "p"},
	{Img: canvas.NewImageFromFile("resources/alphabet/23.png"), Type: Consonant, Rune: "b"},
	{Img: canvas.NewImageFromFile("resources/alphabet/24.png"), Type: Consonant, Rune: "t"},
	{Img: canvas.NewImageFromFile("resources/alphabet/25.png"), Type: Consonant, Rune: "d"},
	{Img: canvas.NewImageFromFile("resources/alphabet/26.png"), Type: Consonant, Rune: "k"},
	{Img: canvas.NewImageFromFile("resources/alphabet/27.png"), Type: Consonant, Rune: "g"},
	{Img: canvas.NewImageFromFile("resources/alphabet/28.png"), Type: Consonant, Rune: "ʤ"},
	{Img: canvas.NewImageFromFile("resources/alphabet/29.png"), Type: Consonant, Rune: "tʃ"},
	{Img: canvas.NewImageFromFile("resources/alphabet/30.png"), Type: Consonant, Rune: "f"},
	{Img: canvas.NewImageFromFile("resources/alphabet/31.png"), Type: Consonant, Rune: "v"},
	{Img: canvas.NewImageFromFile("resources/alphabet/32.png"), Type: Consonant, Rune: "θ"},
	{Img: canvas.NewImageFromFile("resources/alphabet/33.png"), Type: Consonant, Rune: "ð"},
	{Img: canvas.NewImageFromFile("resources/alphabet/34.png"), Type: Consonant, Rune: "s"},
	{Img: canvas.NewImageFromFile("resources/alphabet/35.png"), Type: Consonant, Rune: "z"},
	{Img: canvas.NewImageFromFile("resources/alphabet/36.png"), Type: Consonant, Rune: "ʃ"},
	{Img: canvas.NewImageFromFile("resources/alphabet/37.png"), Type: Consonant, Rune: ""},
	{Img: canvas.NewImageFromFile("resources/alphabet/38.png"), Type: Consonant, Rune: "ˈh"},
	{Img: canvas.NewImageFromFile("resources/alphabet/39.png"), Type: Consonant, Rune: "r"},
	{Img: canvas.NewImageFromFile("resources/alphabet/40.png"), Type: Consonant, Rune: "j"},
	{Img: canvas.NewImageFromFile("resources/alphabet/41.png"), Type: Consonant, Rune: "w"},
	{Img: canvas.NewImageFromFile("resources/alphabet/42.png"), Type: Consonant, Rune: "l"},
}

var specialChars = []alphabetItem{
	{Img: canvas.NewImageFromFile("resources/punctuation/space.png"), Type: Special, Rune: " "},
	{Img: canvas.NewImageFromFile("resources/punctuation/exclamation mark.png"), Type: Special, Rune: "!"},
	{Img: canvas.NewImageFromFile("resources/punctuation/question mark.png"), Type: Special, Rune: "?"},
}
