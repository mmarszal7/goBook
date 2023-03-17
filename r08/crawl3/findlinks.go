// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Crawl3 indeksuje strony internetowe rozpoczynając od argumentów wiersza poleceń.
//
// Ta wersja używa ograniczonej współbieżności.
// Dla uproszczenia nie adresuje problemu przerwania działania.
//
package main

import (
	"fmt"
	"log"
	"os"

	"code/r05/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!+
func main() {
	worklist := make(chan []string)  // listy adresów URL; mogą mieć duplikaty
	unseenLinks := make(chan string) // adresy URL bez duplikatów

	// Dodaje do worklist argumenty wiersza poleceń.
	go func() { worklist <- os.Args[1:] }()

	// Tworzy 20 funkcji goroutine indeksowania, aby pobrać każdy niewidziany jeszcze link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// Główna funkcja goroutine usuwa duplikaty elementów worklist
	// i wysyła niewidziane jeszcze elementy do funkcji indeksowania.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

//!-
