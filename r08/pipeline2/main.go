// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Pipeline2 demonstruje skończony potok trzyetapowy.
package main

import "fmt"

//!+
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Licznik.
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Potęga kwadratowa.
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Wyświetlacz (w głównej funkcji goroutine).
	for x := range squares {
		fmt.Println(x)
	}
}

//!-
