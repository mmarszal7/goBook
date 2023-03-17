// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Echo4 wyświetla swoje argumenty wiersza poleceń.
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "pominięcie na końcu znaku nowej linii")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

//!-
