// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Crawl1 indeksuje strony internetowe rozpoczynając od argumentów wiersza poleceń.
//
// Ta wersja szybko wyczerpuje dostępne deskryptory plików
// z uwagi na nadmierne współbieżne wywołania links.Extract.
//
// Ponadto nigdy nie kończy działania, ponieważ worklist nigdy nie jest zamykana.
package main

import (
	"fmt"
	"log"
	"os"

	"code/r05/links"
)

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
	worklist := make(chan []string)

	// Rozpoczynanie od argumentów wiersza poleceń.
	go func() { worklist <- os.Args[1:] }()

	// Współbieżne indeksowanie stron internetowych.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

//!-main

/*
//!+output
$ go build code/r08/crawl1
$ ./crawl1 http://gopl.io/
http://gopl.io/
https://golang.org/help/

https://golang.org/doc/
https://golang.org/blog/
...
2015/07/15 18:22:12 Get ...: dial tcp: lookup blog.golang.org: no such host
2015/07/15 18:22:12 Get ...: dial tcp 23.21.222.120:443: socket:
                                                        too many open files
...
//!-output
*/
