// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Fetchall pobiera równolegle zawartości kilku adresów URL i raportuje czasy pobierania oraz rozmiary odpowiedzi.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // rozpoczęcie funkcji goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // odbieranie z kanału ch
	}
	fmt.Printf("%.2fs upłynęło\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // wysyłanie do kanału ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // aby nie wyciekały zasoby
	if err != nil {
		ch <- fmt.Sprintf("podczas odczytywania %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
