package fyne

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type alphabetItemType int

const (
	Vowel alphabetItemType = iota
	Consonant
	Special
)

type alphabetItem struct {
	ID       string               `json:"id"`
	Img      *canvas.Image        `json:"-"`
	Resource *fyne.StaticResource `json:"-"` // better store both the image and the resource to pre-load them in the memory
	Rune     string               `json:"rune"`
	Type     alphabetItemType     `json:"type"`
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
	{ID: "0", Img: canvas.NewImageFromResource(resource1Png), Resource: resource1Png, Type: Vowel, Rune: "ɔː"},
	{ID: "1", Img: canvas.NewImageFromResource(resource2Png), Resource: resource2Png, Type: Vowel, Rune: "ɑː"},
	{ID: "2", Img: canvas.NewImageFromResource(resource3Png), Resource: resource3Png, Type: Vowel, Rune: "ɪ"},
	{ID: "3", Img: canvas.NewImageFromResource(resource4Png), Resource: resource4Png, Type: Vowel, Rune: "e"},
	{ID: "4", Img: canvas.NewImageFromResource(resource5Png), Resource: resource5Png, Type: Vowel, Rune: "ʊ"},
	{ID: "5", Img: canvas.NewImageFromResource(resource6Png), Resource: resource6Png, Type: Vowel, Rune: "ʌ"},
	{ID: "6", Img: canvas.NewImageFromResource(resource7Png), Resource: resource7Png, Type: Vowel, Rune: "i:"},
	{ID: "7", Img: canvas.NewImageFromResource(resource8Png), Resource: resource8Png, Type: Vowel, Rune: "uː"},
	{ID: "8", Img: canvas.NewImageFromResource(resource9Png), Resource: resource9Png, Type: Vowel, Rune: "ə"},
	{ID: "9", Img: canvas.NewImageFromResource(resource10Png), Resource: resource10Png, Type: Vowel, Rune: "ɔː"},
	{ID: "10", Img: canvas.NewImageFromResource(resource11Png), Resource: resource11Png, Type: Vowel, Rune: "ɑː"},
	{ID: "11", Img: canvas.NewImageFromResource(resource12Png), Resource: resource12Png, Type: Vowel, Rune: "ɪə"},
	{ID: "12", Img: canvas.NewImageFromResource(resource13Png), Resource: resource13Png, Type: Vowel, Rune: "eɪ"},
	{ID: "13", Img: canvas.NewImageFromResource(resource14Png), Resource: resource14Png, Type: Vowel, Rune: "aɪ"},
	{ID: "14", Img: canvas.NewImageFromResource(resource15Png), Resource: resource15Png, Type: Vowel, Rune: "?"},
	{ID: "15", Img: canvas.NewImageFromResource(resource16Png), Resource: resource16Png, Type: Vowel, Rune: "aʊ"},
	{ID: "16", Img: canvas.NewImageFromResource(resource17Png), Resource: resource17Png, Type: Vowel, Rune: "əʊ"},
	{ID: "17", Img: canvas.NewImageFromResource(resource18Png), Resource: resource18Png, Type: Vowel, Rune: "eə"},
}
var defaultConsonants = []alphabetItem{
	{ID: "18", Img: canvas.NewImageFromResource(resource19Png), Resource: resource19Png, Type: Consonant, Rune: "m"},
	{ID: "19", Img: canvas.NewImageFromResource(resource20Png), Resource: resource20Png, Type: Consonant, Rune: "n"},
	{ID: "20", Img: canvas.NewImageFromResource(resource21Png), Resource: resource21Png, Type: Consonant, Rune: "ŋ"},
	{ID: "21", Img: canvas.NewImageFromResource(resource22Png), Resource: resource22Png, Type: Consonant, Rune: "p"},
	{ID: "22", Img: canvas.NewImageFromResource(resource23Png), Resource: resource23Png, Type: Consonant, Rune: "b"},
	{ID: "23", Img: canvas.NewImageFromResource(resource24Png), Resource: resource24Png, Type: Consonant, Rune: "t"},
	{ID: "24", Img: canvas.NewImageFromResource(resource25Png), Resource: resource25Png, Type: Consonant, Rune: "d"},
	{ID: "25", Img: canvas.NewImageFromResource(resource26Png), Resource: resource26Png, Type: Consonant, Rune: "k"},
	{ID: "26", Img: canvas.NewImageFromResource(resource27Png), Resource: resource27Png, Type: Consonant, Rune: "g"},
	{ID: "27", Img: canvas.NewImageFromResource(resource28Png), Resource: resource28Png, Type: Consonant, Rune: "ʤ"},
	{ID: "28", Img: canvas.NewImageFromResource(resource29Png), Resource: resource29Png, Type: Consonant, Rune: "tʃ"},
	{ID: "29", Img: canvas.NewImageFromResource(resource30Png), Resource: resource30Png, Type: Consonant, Rune: "f"},
	{ID: "30", Img: canvas.NewImageFromResource(resource31Png), Resource: resource31Png, Type: Consonant, Rune: "v"},
	{ID: "31", Img: canvas.NewImageFromResource(resource32Png), Resource: resource32Png, Type: Consonant, Rune: "θ"},
	{ID: "32", Img: canvas.NewImageFromResource(resource33Png), Resource: resource33Png, Type: Consonant, Rune: "ð"},
	{ID: "33", Img: canvas.NewImageFromResource(resource34Png), Resource: resource34Png, Type: Consonant, Rune: "s"},
	{ID: "34", Img: canvas.NewImageFromResource(resource35Png), Resource: resource35Png, Type: Consonant, Rune: "z"},
	{ID: "35", Img: canvas.NewImageFromResource(resource36Png), Resource: resource36Png, Type: Consonant, Rune: "ʃ"},
	{ID: "36", Img: canvas.NewImageFromResource(resource37Png), Resource: resource37Png, Type: Consonant, Rune: "?"},
	{ID: "37", Img: canvas.NewImageFromResource(resource38Png), Resource: resource38Png, Type: Consonant, Rune: "ˈh"},
	{ID: "38", Img: canvas.NewImageFromResource(resource39Png), Resource: resource39Png, Type: Consonant, Rune: "r"},
	{ID: "39", Img: canvas.NewImageFromResource(resource40Png), Resource: resource40Png, Type: Consonant, Rune: "j"},
	{ID: "40", Img: canvas.NewImageFromResource(resource41Png), Resource: resource41Png, Type: Consonant, Rune: "w"},
	{ID: "41", Img: canvas.NewImageFromResource(resource42Png), Resource: resource42Png, Type: Consonant, Rune: "l"},
}

var specialChars = []alphabetItem{
	{ID: "42", Img: canvas.NewImageFromResource(resourceSpacePng), Resource: resourceSpacePng, Type: Special, Rune: " "},
	{ID: "43", Img: canvas.NewImageFromResource(resourceExclamationMarkPng), Resource: resourceExclamationMarkPng, Type: Special, Rune: "!"},
	{ID: "44", Img: canvas.NewImageFromResource(resourceQuestionMarkPng), Resource: resourceQuestionMarkPng, Type: Special, Rune: "?"},
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
