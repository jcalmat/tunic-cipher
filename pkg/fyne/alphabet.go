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

var alphabetMap = []alphabetItem{
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232424.png"), Type: Vowel, Rune: "ɔː"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232430.png"), Type: Vowel, Rune: "ɑː"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232439.png"), Type: Vowel, Rune: "ɪ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232447.png"), Type: Vowel, Rune: "e"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240810144548.png"), Type: Vowel, Rune: "ʊ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232458.png"), Type: Vowel, Rune: "ʌ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232516.png"), Type: Vowel, Rune: "i:"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232524.png"), Type: Vowel, Rune: "uː"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232547.png"), Type: Vowel, Rune: "ə"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232559.png"), Type: Vowel, Rune: "ɔː"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232614.png"), Type: Vowel, Rune: "ɑː"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232623.png"), Type: Vowel, Rune: "ɪə"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232629.png"), Type: Vowel, Rune: "eɪ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232638.png"), Type: Vowel, Rune: "aɪ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232644.png"), Type: Vowel, Rune: ""},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232743.png"), Type: Vowel, Rune: "aʊ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232750.png"), Type: Vowel, Rune: "əʊ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232807.png"), Type: Vowel, Rune: "eə"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232816.png"), Type: Vowel, Rune: "m"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232822.png"), Type: Consonant, Rune: "n"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232828.png"), Type: Consonant, Rune: "ŋ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232835.png"), Type: Consonant, Rune: "p"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232840.png"), Type: Consonant, Rune: "b"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232846.png"), Type: Consonant, Rune: "t"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232851.png"), Type: Consonant, Rune: "d"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232855.png"), Type: Consonant, Rune: "k"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232907.png"), Type: Consonant, Rune: "g"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232914.png"), Type: Consonant, Rune: "ʤ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232922.png"), Type: Consonant, Rune: "tʃ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232929.png"), Type: Consonant, Rune: "f"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232938.png"), Type: Consonant, Rune: "v"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232946.png"), Type: Consonant, Rune: "θ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232953.png"), Type: Consonant, Rune: "ð"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232959.png"), Type: Consonant, Rune: "s"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808233003.png"), Type: Consonant, Rune: "z"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808233008.png"), Type: Consonant, Rune: "ʃ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808233011.png"), Type: Consonant, Rune: ""},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808233016.png"), Type: Consonant, Rune: "ˈh"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808233020.png"), Type: Consonant, Rune: "r"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808233024.png"), Type: Consonant, Rune: "j"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808233028.png"), Type: Consonant, Rune: "w"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808233033.png"), Type: Consonant, Rune: "l"},
}
