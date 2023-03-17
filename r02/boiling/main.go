// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Boiling wyświetla temperaturę wrzenia wody.
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("temperatura wrzenia = %g°F or %g°C\n", f, c)
	// Output:
	// temperatura wrzenia = 212°F or 100°C
}

//!-
