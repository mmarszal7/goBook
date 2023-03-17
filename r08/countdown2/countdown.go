// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Countdown implementuje odliczanie do odpalenia rakiety.
package main

import (
	"fmt"
	"os"
	"time"
)

//!+

func main() {
	// ...tworzenie kanału abort...

	//!-

	//!+abort
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // odczyt pojedynczego bajtu
		abort <- struct{}{}
	}()
	//!-abort

	//!+
	fmt.Println("Rozpoczynam odliczanie. Aby przerwać, wciśnij Enter.")
	select {
	case <-time.After(10 * time.Second):
		// Nic nie rów.
	case <-abort:
		fmt.Println("Odpalanie przerwane!")
		return
	}
	launch()
}

//!-

func launch() {
	fmt.Println("Wystartowała!")
}
