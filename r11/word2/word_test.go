// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package word

import (
	"fmt"
	"math/rand"
	"time"
)

//!+bench

import "testing"

//!-bench

//!+test
func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kajak", true},
		{"owocowo", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrom", false}, // to nie jest palindrom
		{"żartem", false},   // to jest półpalindrom
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

//!-test

//!+bench
func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

//!-bench

//!+example

func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("palindrom"))
	// Output:
	// true
	// false
}

//!-example

/*
//!+random
import "math/rand"

//!-random
*/

//!+random
// randomPalindrome zwraca palindrom, którego długość i treść
// są uzyskiwane na podstawie wskazań pseudolosowego generatora liczb rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // losowa długość do 24 znaków
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // losowa runa z zakresu do '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Inicjowanie pseudolosowego generatora liczb.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

//!-random

/*
// Odpowiedż do ćwiczenia 11.1: zmodyfikuj randomPalindrome, aby sprawdzała
// obsługę IsPalindrome dla interpunkcji i spacji.

// UWAGA: konwersja r -> "wielka litera" -> "mała litera" nie zachowuje
// wartości r w niektórych przypadkach, np., µ Μ, ſ S, ı I

// randomPalindrome zwraca palindrom, którego długość i treść
// są uzyskiwane na podstawie wskazań pseudolosowego generatora liczb rng.
func randomNoisyPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // losowa długość do 24 znaków
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x200)) // losowa runa z zakresu do \u99
		runes[i] = r
		r1 := r
		if unicode.IsLetter(r) && unicode.IsLower(r) {
			r = unicode.ToUpper(r)
			if unicode.ToLower(r) != r1 {
				fmt.Printf("cap? %c %c\n", r1, r)
			}
		}
		runes[n-1-i] = r
	}
	return "?" + string(runes) + "!"
}

func TestRandomNoisyPalindromes(t *testing.T) {
	// Inicjowanie pseudolosowego generatora liczb.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	n := 0
	for i := 0; i < 1000; i++ {
		p := randomNoisyPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsNoisyPalindrome(%q) = false", p)
			n++
		}
	}
	fmt.Fprintf(os.Stderr, "fail = %d\n", n)
}
*/
