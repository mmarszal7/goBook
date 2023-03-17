// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package word zapewnia narzędzia do gier słownych.
package word

// IsPalindrome raportuje, czy s czyta się tak samo od lewej do prawej i od prawej do lewej.
// (Nasza pierwsza próba).
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

//!-
