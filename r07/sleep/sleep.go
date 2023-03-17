// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Sleep zostaje uśpiony na określony przedział czasu.
package main

import (
	"flag"
	"fmt"
	"time"
)

//!+sleep
var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("śpi przez %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}

//!-sleep
