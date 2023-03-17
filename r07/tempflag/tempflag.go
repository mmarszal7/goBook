// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Tempflag wyświetla wartość swojej flagi -temp (temperatura).
package main

import (
	"flag"
	"fmt"

	"code/r07/tempconv"
)

//!+
var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperatura")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}

//!-
