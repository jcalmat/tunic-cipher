package fyne

import (
	"strings"

	"fyne.io/fyne/v2/canvas"
)

type alphabetItemType int

const (
	Vowel alphabetItemType = iota
	Consonant
	Special
)

type alphabetItem struct {
	ID   string           `json:"id"`
	Img  *canvas.Image    `json:"img"`
	Rune string           `json:"rune"`
	Type alphabetItemType `json:"type"`
}

type alphabetItems []alphabetItem

func (a alphabetItems) String() string {
	var s string
	for _, item := range a {
		s += item.Rune
	}
	return s
}

var vowels = []alphabetItem{
	{ID: "0", Img: canvas.NewImageFromFile("resources/alphabet/1.png"), Type: Vowel, Rune: "ɔː"},
	{ID: "1", Img: canvas.NewImageFromFile("resources/alphabet/2.png"), Type: Vowel, Rune: "ɑː"},
	{ID: "2", Img: canvas.NewImageFromFile("resources/alphabet/3.png"), Type: Vowel, Rune: "ɪ"},
	{ID: "3", Img: canvas.NewImageFromFile("resources/alphabet/4.png"), Type: Vowel, Rune: "e"},
	{ID: "4", Img: canvas.NewImageFromFile("resources/alphabet/5.png"), Type: Vowel, Rune: "ʊ"},
	{ID: "5", Img: canvas.NewImageFromFile("resources/alphabet/6.png"), Type: Vowel, Rune: "ʌ"},
	{ID: "6", Img: canvas.NewImageFromFile("resources/alphabet/7.png"), Type: Vowel, Rune: "i:"},
	{ID: "7", Img: canvas.NewImageFromFile("resources/alphabet/8.png"), Type: Vowel, Rune: "uː"},
	{ID: "8", Img: canvas.NewImageFromFile("resources/alphabet/9.png"), Type: Vowel, Rune: "ə"},
	{ID: "9", Img: canvas.NewImageFromFile("resources/alphabet/10.png"), Type: Vowel, Rune: "ɔː"},
	{ID: "10", Img: canvas.NewImageFromFile("resources/alphabet/11.png"), Type: Vowel, Rune: "ɑː"},
	{ID: "11", Img: canvas.NewImageFromFile("resources/alphabet/12.png"), Type: Vowel, Rune: "ɪə"},
	{ID: "12", Img: canvas.NewImageFromFile("resources/alphabet/13.png"), Type: Vowel, Rune: "eɪ"},
	{ID: "13", Img: canvas.NewImageFromFile("resources/alphabet/14.png"), Type: Vowel, Rune: "aɪ"},
	{ID: "14", Img: canvas.NewImageFromFile("resources/alphabet/15.png"), Type: Vowel, Rune: "?"},
	{ID: "15", Img: canvas.NewImageFromFile("resources/alphabet/16.png"), Type: Vowel, Rune: "aʊ"},
	{ID: "16", Img: canvas.NewImageFromFile("resources/alphabet/17.png"), Type: Vowel, Rune: "əʊ"},
	{ID: "17", Img: canvas.NewImageFromFile("resources/alphabet/18.png"), Type: Vowel, Rune: "eə"},
}
var consonants = []alphabetItem{
	{ID: "18", Img: canvas.NewImageFromFile("resources/alphabet/19.png"), Type: Consonant, Rune: "m"},
	{ID: "19", Img: canvas.NewImageFromFile("resources/alphabet/20.png"), Type: Consonant, Rune: "n"},
	{ID: "20", Img: canvas.NewImageFromFile("resources/alphabet/21.png"), Type: Consonant, Rune: "ŋ"},
	{ID: "21", Img: canvas.NewImageFromFile("resources/alphabet/22.png"), Type: Consonant, Rune: "p"},
	{ID: "22", Img: canvas.NewImageFromFile("resources/alphabet/23.png"), Type: Consonant, Rune: "b"},
	{ID: "23", Img: canvas.NewImageFromFile("resources/alphabet/24.png"), Type: Consonant, Rune: "t"},
	{ID: "24", Img: canvas.NewImageFromFile("resources/alphabet/25.png"), Type: Consonant, Rune: "d"},
	{ID: "25", Img: canvas.NewImageFromFile("resources/alphabet/26.png"), Type: Consonant, Rune: "k"},
	{ID: "26", Img: canvas.NewImageFromFile("resources/alphabet/27.png"), Type: Consonant, Rune: "g"},
	{ID: "27", Img: canvas.NewImageFromFile("resources/alphabet/28.png"), Type: Consonant, Rune: "ʤ"},
	{ID: "28", Img: canvas.NewImageFromFile("resources/alphabet/29.png"), Type: Consonant, Rune: "tʃ"},
	{ID: "29", Img: canvas.NewImageFromFile("resources/alphabet/30.png"), Type: Consonant, Rune: "f"},
	{ID: "30", Img: canvas.NewImageFromFile("resources/alphabet/31.png"), Type: Consonant, Rune: "v"},
	{ID: "31", Img: canvas.NewImageFromFile("resources/alphabet/32.png"), Type: Consonant, Rune: "θ"},
	{ID: "32", Img: canvas.NewImageFromFile("resources/alphabet/33.png"), Type: Consonant, Rune: "ð"},
	{ID: "33", Img: canvas.NewImageFromFile("resources/alphabet/34.png"), Type: Consonant, Rune: "s"},
	{ID: "34", Img: canvas.NewImageFromFile("resources/alphabet/35.png"), Type: Consonant, Rune: "z"},
	{ID: "35", Img: canvas.NewImageFromFile("resources/alphabet/36.png"), Type: Consonant, Rune: "ʃ"},
	{ID: "36", Img: canvas.NewImageFromFile("resources/alphabet/37.png"), Type: Consonant, Rune: "?"},
	{ID: "37", Img: canvas.NewImageFromFile("resources/alphabet/38.png"), Type: Consonant, Rune: "ˈh"},
	{ID: "38", Img: canvas.NewImageFromFile("resources/alphabet/39.png"), Type: Consonant, Rune: "r"},
	{ID: "39", Img: canvas.NewImageFromFile("resources/alphabet/40.png"), Type: Consonant, Rune: "j"},
	{ID: "40", Img: canvas.NewImageFromFile("resources/alphabet/41.png"), Type: Consonant, Rune: "w"},
	{ID: "41", Img: canvas.NewImageFromFile("resources/alphabet/42.png"), Type: Consonant, Rune: "l"},
}

var specialChars = []alphabetItem{
	{ID: "42", Img: canvas.NewImageFromFile("resources/punctuation/space.png"), Type: Special, Rune: " "},
	{ID: "43", Img: canvas.NewImageFromFile("resources/punctuation/exclamation mark.png"), Type: Special, Rune: "!"},
	{ID: "44", Img: canvas.NewImageFromFile("resources/punctuation/question mark.png"), Type: Special, Rune: "?"},
}

func getAlphabetItems(ids []string) []alphabetItem {
	ret := make([]alphabetItem, 0)

	all := append(vowels, append(consonants, specialChars...)...)

	// worst code ever, put some protection glasses on
	for _, id := range ids {
		for _, a := range all {
			if a.ID == id {
				ret = append(ret, a)
				break
			}
		}
	}
	return ret
}

func (a alphabetItems) ToSave() string {
	ret := strings.Builder{}
	for _, item := range a {
		ret.WriteString(item.ID)
		ret.WriteString(" ")
	}
	return ret.String()
}

func (a alphabetItems) FromSave(s string) alphabetItems {
	ids := strings.Split(s, " ")
	return getAlphabetItems(ids)
}
