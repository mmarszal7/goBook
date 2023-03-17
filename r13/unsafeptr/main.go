// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Package unsafeptr demonstruje podstawowe użycie unsafe.Pointer.
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//!+main
	var x struct {
		a bool
		b int16
		c []int
	}

	// Równoważne z pb := &x.b.
	pb := (*int16)(unsafe.Pointer(
		uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42

	fmt.Println(x.b) // "42"
	//!-main
}

/*
//!+wrong
	// UWAGA: subtelnie nieprawidłowe!
	tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	pb := (*int16)(unsafe.Pointer(tmp))
	*pb = 42
//!-wrong
*/
