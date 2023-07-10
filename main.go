package main

import (
	"fmt"
	"strings"
)

var alphabet = "abcdefghijklmnopqrstuvwxyz"
var ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func decrypt(n int, ciphertext string) string {
	var result strings.Builder
	for _, l := range ciphertext {
		if strings.ContainsRune(ALPHABET, l) {
			index := strings.IndexRune(ALPHABET, l)
			i := (index - n + 26) % 26
			result.WriteRune(rune(ALPHABET[i]))
		} else if strings.ContainsRune(alphabet, l) {
			index := strings.IndexRune(alphabet, l)
			i := (index - n + 26) % 26
			result.WriteRune(rune(alphabet[i]))
		} else {
			result.WriteRune(l)
		}
	}
	return result.String()
}

func main() {
	key := 0
	message := "Ugew gnwj zwjw Oslkgf"
	for key < 26 {
		dec := decrypt(key, message)
		fmt.Printf("Shift %d . %s\n", key, dec)
		key++
	}
}
