// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Fetch zapisuje zawartość adresu URL w lokalnym pliku.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

//!+
// Fetch pobiera zawartość adresu URL 
// i zwraca nazwę oraz wielkość lokalnego pliku.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Zamykanie pliku, przy czym preferowane jest zgłaszanie błędu z funkcji Copy, 
	// jeśli w ogóle wystąpi jakiś błąd.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}

//!-

func main() {
	for _, url := range os.Args[1:] {
		local, n, err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
			continue
		}
		fmt.Fprintf(os.Stderr, "%s => %s (%d bytów).\n", url, local, n)
	}
}
