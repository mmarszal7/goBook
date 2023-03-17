// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Countdown implementuje odliczanie do odpalenia rakiety.
package main

// UWAGA: funkcja goroutine tickera nigdy nie kończy działania, jeśli odpalanie zostanie przerwane.
// Jest to "wyciek funkcji goroutine".

import (
	"fmt"
	"os"
	"time"
)

//!+

func main() {
	// ...tworzenie kanału abort...

	//!-

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // odczyt pojedynczego bajtu
		abort <- struct{}{}
	}()

	//!+
	fmt.Println("Rozpoczynam odliczanie. Aby przerwać, wciśnij Enter.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// Nic nie rób.
		case <-abort:
			fmt.Println("Odpalanie przerwane!")
			return
		}
	}
	launch()
}

//!-

func launch() {
	fmt.Println("Wystartowała!")
}
