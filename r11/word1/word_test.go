// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+test
package word

import "testing"

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("owocowo") {
		t.Error(`IsPalindrome("owocowo") = false`)
	}
	if !IsPalindrome("kajak") {
		t.Error(`IsPalindrome("kajak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrom") {
		t.Error(`IsPalindrome("palindrom") = true`)
	}
}

//!-test

// Dla poniższych testów spodziewane są wyniki negatywne.
// Poprawką znajdziesz w code/r11/word2.

//!+more
func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}

//!-more
