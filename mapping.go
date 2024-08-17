package main

type rune struct {
	phoneme   string
	confident bool
}

var alphabetMap = map[int]rune{
	0:  {phoneme: "ɔː", confident: true},
	1:  {phoneme: "ɑː", confident: false},
	2:  {phoneme: "ɪ", confident: true},
	3:  {phoneme: "e", confident: false},
	4:  {phoneme: "ʊ", confident: true},
	5:  {phoneme: "ʌ", confident: true},
	6:  {phoneme: "i:", confident: true},
	7:  {phoneme: "uː", confident: false},
	8:  {phoneme: "ə", confident: true},
	9:  {phoneme: "ɔː", confident: false},
	10: {phoneme: "ɑː", confident: true},
	11: {phoneme: "ɪə", confident: false},
	12: {phoneme: "eɪ", confident: true},
	13: {phoneme: "aɪ", confident: false},
	14: {phoneme: " ", confident: true},
	15: {phoneme: "aʊ", confident: true},
	16: {phoneme: "əʊ", confident: true},
	17: {phoneme: "eə", confident: false},
	18: {phoneme: "m", confident: true},
	19: {phoneme: "n", confident: true},
	20: {phoneme: "ŋ", confident: false},
	21: {phoneme: "p", confident: true},
	22: {phoneme: "b", confident: true},
	23: {phoneme: "t", confident: true},
	24: {phoneme: "d", confident: true},
	25: {phoneme: "k", confident: true},
	26: {phoneme: "g", confident: false},
	27: {phoneme: "ʤ", confident: true},
	28: {phoneme: "tʃ", confident: true},
	29: {phoneme: "f", confident: true},
	30: {phoneme: "v", confident: true},
	31: {phoneme: "θ", confident: true},
	32: {phoneme: "ð", confident: false},
	33: {phoneme: "s", confident: true},
	34: {phoneme: "z", confident: true},
	35: {phoneme: "ʃ", confident: true},
	36: {phoneme: " ", confident: true},
	37: {phoneme: "ˈh", confident: false},
	38: {phoneme: "r", confident: true},
	39: {phoneme: "j", confident: true},
	40: {phoneme: "w", confident: true},
	41: {phoneme: "l", confident: true},
	-1: {phoneme: " ", confident: true}, // space
}
