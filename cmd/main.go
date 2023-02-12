package main

import (
	//"fmt"
	"strings"
)

var MORSE_CODE = map[string]string{
	"a": ".—",
	"b": "—...",
	"c": "—.—.",
	"d": "—..",
	"e": ".",
	"f": "..—.",
	"g": "——.",
	"h": "....",
	"i": "..",
	"j": ".———",
	"k": "—.—",
	"l": ".—..",
	"m": "——",
	"n": "—.",
	"o": "———",
	"p": ".——.",
	"q": "——.—",
	"r": ".—.",
	"s": "...",
	"t": "—",
	"u": "..—",
	"v": "...—",
	"w": ".——",
	"x": "—..—",
	"y": "—.——",
	"z": "——..",
}

func main() {
	return
}

func DecodeMorse(morseCode string) (res string) {
	cleanPhrase := strings.TrimSpace(morseCode)
	words := strings.Split(cleanPhrase, "   ")
	var word []string
	for _, w := range words {
		word = strings.Split(w, " ")
		for _, letter := range word {
			res += MORSE_CODE[letter]
		}
		res += " "
	}
	res = strings.TrimSpace(res)
	return
}
