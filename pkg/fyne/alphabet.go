package fyne

import (
	"fyne.io/fyne/v2/canvas"
)

type alphabetItem struct {
	Img  *canvas.Image
	Rune string
}

var alphabetMap = []alphabetItem{
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232424.png"), Rune: "ɔː"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232430.png"), Rune: "ɑː"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232439.png"), Rune: "ɪ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232447.png"), Rune: "e"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240810144548.png"), Rune: "ʊ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232458.png"), Rune: "ʌ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232516.png"), Rune: "i:"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232524.png"), Rune: "uː"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232547.png"), Rune: "ə"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232559.png"), Rune: "ɔː"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232614.png"), Rune: "ɑː"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232623.png"), Rune: "ɪə"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232629.png"), Rune: "eɪ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232638.png"), Rune: "aɪ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232644.png"), Rune: ""},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232743.png"), Rune: "aʊ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232750.png"), Rune: "əʊ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232807.png"), Rune: "eə"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232816.png"), Rune: "m"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232822.png"), Rune: "n"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232828.png"), Rune: "ŋ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232835.png"), Rune: "p"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232840.png"), Rune: "b"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232846.png"), Rune: "t"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232851.png"), Rune: "d"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232855.png"), Rune: "k"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232907.png"), Rune: "g"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232914.png"), Rune: "ʤ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232922.png"), Rune: "tʃ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232929.png"), Rune: "f"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232938.png"), Rune: "v"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232946.png"), Rune: "θ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232953.png"), Rune: "ð"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808232959.png"), Rune: "s"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808233003.png"), Rune: "z"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808233008.png"), Rune: "ʃ"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808233011.png"), Rune: ""},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808233016.png"), Rune: "ˈh"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808233020.png"), Rune: "r"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808233024.png"), Rune: "j"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808233028.png"), Rune: "w"},
	{Img: canvas.NewImageFromFile("./pkg/fyne/data/alphabet/Pasted image 20240808233033.png"), Rune: "l"},
}
