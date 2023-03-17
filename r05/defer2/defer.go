// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Defer2 demonstruje odroczone wywołanie runtime.Stack podczas paniki.
package main

import (
	"fmt"
	"os"
	"runtime"
)

//!+
func main() {
	defer printStack()
	f(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

//!-

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panika, jeśli x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

/*
//!+printstack
goroutine 1 [running]:
main.printStack()
	src/code/r05/defer2/defer.go:20
main.f(0)
	src/code/r05/defer2/defer.go:27
main.f(1)
	src/code/r05/defer2/defer.go:29
main.f(2)
	src/code/r05/defer2/defer.go:29
main.f(3)
	src/code/r05/defer2/defer.go:29
main.main()
	src/code/r05/defer2/defer.go:15
//!-printstack
*/
