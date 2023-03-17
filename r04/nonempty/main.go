// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


//!+nonempty

// Nonempty jest przykładem algorytmu in situ wycinka.
package main

import "fmt"

// nonempty zwraca wycinek przechowujący tylko niepuste łańcuchy.
// Podczas wywoływania funkcji modyfikowana jest tablica bazowa.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

//!-nonempty

func main() {
	//!+main
	data := []string{"jeden", "", "trzy"}
	fmt.Printf("%q\n", nonempty(data)) // `["jeden" "trzy"]`
	fmt.Printf("%q\n", data)           // `["jeden" "trzy" "trzy"]`
	//!-main
}

//!+alt
func nonempty2(strings []string) []string {
	out := strings[:0] // zerowej długości wycinek oryginału
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

//!-alt
