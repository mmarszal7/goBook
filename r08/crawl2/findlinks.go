// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Crawl2 indeksuje strony internetowe rozpoczynając od argumentów wiersza poleceń.
//
// Ta wersja używa kanału buforowanego jako semafora zliczającego,
// aby ograniczyć liczbę współbieżnych wywołań links.Extract.
package main

import (
	"fmt"
	"log"
	"os"

	"code/r05/links"
)

//!+sema
// tokens jest semaforem zliczającym wykorzystywanym 
// do wyegzekwowania limitu 20 współbieżnych żądań.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // nabycie żetonu
	list, err := links.Extract(url)
	<-tokens // zwolnienie żetonu

	if err != nil {
		log.Print(err)
	}
	return list
}

//!-sema

//!+
func main() {
	worklist := make(chan []string)
	var n int // liczba oczekujących operacji wysłania do worklist

	// Rozpoczynanie od argumentów wiersza poleceń.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Współbieżne indeksowanie stron internetowych.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

//!-
