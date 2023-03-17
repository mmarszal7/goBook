// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Basename1 odczytuje nazwy plików z stdin i wyświetla bazową nazwę każdego z nich.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
	// UWAGA: ignorowanie potencjalnych błędów z funkcji input.Err()
}

//!+
// basename usuwa komponenty ścieżki katalogu oraz przyrostek po kropce,
// np., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	// Porzuca ostatni znak '/' i wszystko, co znajduje się przed nim.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Zachowuje wszystko przed ostatnim znakiem '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

//!-
