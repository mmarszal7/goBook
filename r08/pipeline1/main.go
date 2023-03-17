// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Pipeline1 demonstruje nieskończony potok trzyetapowy.
package main

import "fmt"

//!+
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Licznik.
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Potęga kwadratowa.
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// Wyświetlacz (w głównej funkcji goroutine).
	for {
		fmt.Println(<-squares)
	}
}

//!-
