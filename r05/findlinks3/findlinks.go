// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Findlinks3 indeksuje strony internetowe rozpoczynając od adresów URL podanych w wierszu poleceń.
package main

import (
	"fmt"
	"log"
	"os"

	"code/r05/links"
)

//!+breadthFirst
// breadthFirst wywołuje funkcję f dla każdej pozycji z listy worklist.
// Wszystkie pozycje zwrócone przez funkcję f są dodawane do worklist.
// Funkcja f jest wywoływana co najwyżej raz dla każdej pozycji.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

//!+crawl
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!-crawl

//!+main
func main() {
	// Indeksowanie stron internetowych za pomocą przechodzenia wszerz,
	// rozpoczynane od argumentów wiersza poleceń.
	breadthFirst(crawl, os.Args[1:])
}

//!-main
