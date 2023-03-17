// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Trace wykorzystuje instrukcję defer, aby dodać do funkcji diagnostykę wejścia/wyjścia.
package main

import (
	"log"
	"time"
)

//!+main
func bigSlowOperation() {
	defer trace("bigSlowOperation")() // nie zapomnij o dodatkowych nawiasach
	// ...dużo pracy...
	time.Sleep(10 * time.Second) // symulowanie powolnej operacji za pomocą funkcji Sleep
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("punkt wejści %s", msg)
	return func() { log.Printf("punkt wyjścia %s (%s)", msg, time.Since(start)) }
}

//!-main

func main() {
	bigSlowOperation()
}

/*
!+output
$ go build code/r05/trace
$ ./trace
2015/11/18 09:53:26 punkt wejścia bigSlowOperation
2015/11/18 09:53:36 punkt wyjścia bigSlowOperation (10.000589217s)
!-output
*/
