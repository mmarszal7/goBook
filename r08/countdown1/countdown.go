// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Countdown implementuje odliczanie do odpalenia rakiety.
package main

import (
	"fmt"
	"time"
)

//!+
func main() {
	fmt.Println("Rozpoczynam odliczanie.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}

//!-

func launch() {
	fmt.Println("Wystartowała!")
}
