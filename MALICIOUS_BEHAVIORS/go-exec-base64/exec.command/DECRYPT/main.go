package main

import "fmt"

func main() {
	
	key := map[rune]rune{
		'a': 'z',
		'b': 'y',
		'c': 'x',
		'd': 'w',
		'e': 'v',
		'f': 'u',
		'g': 't',
		'h': 's',
		'i': 'r',
		'j': 'q',
		'k': 'p',
		'l': 'o',
		'm': 'n',
		'n': 'm',
		'o': 'l',
		'p': 'k',
		'q': 'j',
		'r': 'i',
		's': 'h',
		't': 'g',
		'u': 'f',
		'v': 'e',
		'w': 'd',
		'x': 'c',
		'y': 'b',
		'z': 'a',
	}

	encryptedMessage := "svool"

	var decryptedMessage string
	for _, c := range encryptedMessage {
		decryptedMessage += string(key[c])
	}

	fmt.Println(decryptedMessage)
}
