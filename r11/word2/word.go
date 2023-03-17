// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package word zapewnia narzędzia do gier słownych.
package word

import "unicode"

// IsPalindrome raportuje, czy s czyta się tak samo od lewej do prawej i od prawej do lewej.
// Ignorowane są wielkości liter oraz znaki niebędące literami.
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

//!-
