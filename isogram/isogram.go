package isogram

import "unicode"

// IsIsogram returns true if the given word is an isogram, and false if it is not
// An Isogram is a word or phrase without a repeating letter, with the exception of a space
// or hyphen
func IsIsogram(word string) bool {
	duplicateLetter := map[string]int{}

	for _, v := range word {
		v = unicode.ToLower(v)
		_, exist := duplicateLetter[string(v)]

		if exist {
			if !unicode.IsLetter(v) {
				continue
			}
			duplicateLetter[string(v)]++
		} else {
			duplicateLetter[string(v)] = 1
		}

		if duplicateLetter[string(v)] > 1 {
			return false
		}
	}
	return true
}
