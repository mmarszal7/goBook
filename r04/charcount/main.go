// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Charcount liczy wystąpienia znaków Unicode.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // zliczanie wystąpień znaków Unicode
	var utflen [utf8.UTFMax + 1]int // zliczanie długości kodowań UTF-8
	invalid := 0                    // liczba nieprawidłowych znaków UTF-8

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // zwraca runę, liczbę bajtów i błąd
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("runa\tliczba wystąpień\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\ndługość\tliczba wystąpień\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d nniewłaściwych znaków UTF-8\n", invalid)
	}
}

//!-
