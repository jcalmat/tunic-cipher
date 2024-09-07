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
	Img  *canvas.Image    `json:"-"`
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

var defaultVowels = []alphabetItem{
	{ID: "0", Img: canvas.NewImageFromResource(resource1Png), Type: Vowel, Rune: "ɔː"},
	{ID: "1", Img: canvas.NewImageFromResource(resource2Png), Type: Vowel, Rune: "ɑː"},
	{ID: "2", Img: canvas.NewImageFromResource(resource3Png), Type: Vowel, Rune: "ɪ"},
	{ID: "3", Img: canvas.NewImageFromResource(resource4Png), Type: Vowel, Rune: "e"},
	{ID: "4", Img: canvas.NewImageFromResource(resource5Png), Type: Vowel, Rune: "ʊ"},
	{ID: "5", Img: canvas.NewImageFromResource(resource6Png), Type: Vowel, Rune: "ʌ"},
	{ID: "6", Img: canvas.NewImageFromResource(resource7Png), Type: Vowel, Rune: "i:"},
	{ID: "7", Img: canvas.NewImageFromResource(resource8Png), Type: Vowel, Rune: "uː"},
	{ID: "8", Img: canvas.NewImageFromResource(resource9Png), Type: Vowel, Rune: "ə"},
	{ID: "9", Img: canvas.NewImageFromResource(resource10Png), Type: Vowel, Rune: "ɔː"},
	{ID: "10", Img: canvas.NewImageFromResource(resource11Png), Type: Vowel, Rune: "ɑː"},
	{ID: "11", Img: canvas.NewImageFromResource(resource12Png), Type: Vowel, Rune: "ɪə"},
	{ID: "12", Img: canvas.NewImageFromResource(resource13Png), Type: Vowel, Rune: "eɪ"},
	{ID: "13", Img: canvas.NewImageFromResource(resource14Png), Type: Vowel, Rune: "aɪ"},
	{ID: "14", Img: canvas.NewImageFromResource(resource15Png), Type: Vowel, Rune: "?"},
	{ID: "15", Img: canvas.NewImageFromResource(resource16Png), Type: Vowel, Rune: "aʊ"},
	{ID: "16", Img: canvas.NewImageFromResource(resource17Png), Type: Vowel, Rune: "əʊ"},
	{ID: "17", Img: canvas.NewImageFromResource(resource18Png), Type: Vowel, Rune: "eə"},
}
var defaultConsonants = []alphabetItem{
	{ID: "18", Img: canvas.NewImageFromResource(resource19Png), Type: Consonant, Rune: "m"},
	{ID: "19", Img: canvas.NewImageFromResource(resource20Png), Type: Consonant, Rune: "n"},
	{ID: "20", Img: canvas.NewImageFromResource(resource21Png), Type: Consonant, Rune: "ŋ"},
	{ID: "21", Img: canvas.NewImageFromResource(resource22Png), Type: Consonant, Rune: "p"},
	{ID: "22", Img: canvas.NewImageFromResource(resource23Png), Type: Consonant, Rune: "b"},
	{ID: "23", Img: canvas.NewImageFromResource(resource24Png), Type: Consonant, Rune: "t"},
	{ID: "24", Img: canvas.NewImageFromResource(resource25Png), Type: Consonant, Rune: "d"},
	{ID: "25", Img: canvas.NewImageFromResource(resource26Png), Type: Consonant, Rune: "k"},
	{ID: "26", Img: canvas.NewImageFromResource(resource27Png), Type: Consonant, Rune: "g"},
	{ID: "27", Img: canvas.NewImageFromResource(resource28Png), Type: Consonant, Rune: "ʤ"},
	{ID: "28", Img: canvas.NewImageFromResource(resource29Png), Type: Consonant, Rune: "tʃ"},
	{ID: "29", Img: canvas.NewImageFromResource(resource30Png), Type: Consonant, Rune: "f"},
	{ID: "30", Img: canvas.NewImageFromResource(resource31Png), Type: Consonant, Rune: "v"},
	{ID: "31", Img: canvas.NewImageFromResource(resource32Png), Type: Consonant, Rune: "θ"},
	{ID: "32", Img: canvas.NewImageFromResource(resource33Png), Type: Consonant, Rune: "ð"},
	{ID: "33", Img: canvas.NewImageFromResource(resource34Png), Type: Consonant, Rune: "s"},
	{ID: "34", Img: canvas.NewImageFromResource(resource35Png), Type: Consonant, Rune: "z"},
	{ID: "35", Img: canvas.NewImageFromResource(resource36Png), Type: Consonant, Rune: "ʃ"},
	{ID: "36", Img: canvas.NewImageFromResource(resource37Png), Type: Consonant, Rune: "?"},
	{ID: "37", Img: canvas.NewImageFromResource(resource38Png), Type: Consonant, Rune: "ˈh"},
	{ID: "38", Img: canvas.NewImageFromResource(resource39Png), Type: Consonant, Rune: "r"},
	{ID: "39", Img: canvas.NewImageFromResource(resource40Png), Type: Consonant, Rune: "j"},
	{ID: "40", Img: canvas.NewImageFromResource(resource41Png), Type: Consonant, Rune: "w"},
	{ID: "41", Img: canvas.NewImageFromResource(resource42Png), Type: Consonant, Rune: "l"},
}

var specialChars = []alphabetItem{
	{ID: "42", Img: canvas.NewImageFromResource(resourceSpacePng), Type: Special, Rune: " "},
	{ID: "43", Img: canvas.NewImageFromResource(resourceExclamationMarkPng), Type: Special, Rune: "!"},
	{ID: "44", Img: canvas.NewImageFromResource(resourceQuestionMarkPng), Type: Special, Rune: "?"},
}

func getAlphabetItems(ids []string) []alphabetItem {
	ret := make([]alphabetItem, 0)

	all := append(defaultVowels, append(defaultConsonants, specialChars...)...)

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

// ToSave implements the Saver interface
func (a alphabetItems) ToSave() string {
	ret := strings.Builder{}
	for _, item := range a {
		ret.WriteString(item.ID)
		ret.WriteString(" ")
	}
	return ret.String()
}

// FromSave implements the Saver interface
func (a alphabetItems) FromSave(s string) alphabetItems {
	ids := strings.Split(s, " ")
	return getAlphabetItems(ids)
}
