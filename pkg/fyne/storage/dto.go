package storage

// AlphabetItem is a struct that holds the image URI and the rune of an alphabet item
type alphabetItem struct {
	Type  int `json:"type"`
	Index int `json:"index"`
}
