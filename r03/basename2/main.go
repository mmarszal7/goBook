// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Basename2 odczytuje nazwy plików z stdin i wyświetla bazową nazwę każdego z nich.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
	// UWAGA: ignorowanie potencjalnych błędów z funkcji input.Err()
}

// basename usuwa komponenty ścieżki katalogu oraz przyrostek po kropce,
// np., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
//!+
func basename(s string) string {
	slash := strings.LastIndex(s, "/") // –1, jeśli nie znaleziono "/"
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

//!-
