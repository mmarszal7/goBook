// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Bytecounter demonstruje implementację interfejsu io.Writer, która liczy bajty.
package main

import (
	"fmt"
)

//!+bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // konwersja int na ByteCounter
	return len(p), nil
}

//!-bytecounter

func main() {
	//!+main
	var c ByteCounter
	c.Write([]byte("witaj"))
	fmt.Println(c) // "5", = len("witaj")

	c = 0 // resetowanie licznika
	var name = "Marta"
	fmt.Fprintf(&c, "witaj, %s", name)
	fmt.Println(c) // "12", = len("witaj, Marta")
	//!-main
}
