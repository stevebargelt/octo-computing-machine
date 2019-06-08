package ransomletter

import "strings"

func IsRansomLetterReproducible(letter string, newspaper string) bool {

	var availableChars [256]int
	//error checking
	// no nil arrays, etc.
	// if the letter is longer than the newspaper then return false

	letter = strings.Replace(letter, " ", "", -1)

	for _, char := range newspaper {
		availableChars[char]++
	}

	for _, char := range letter {
		availableChars[char]--
		if availableChars[char] < 0 {
			return false
		}
	}

	return true
}
