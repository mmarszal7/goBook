// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Server1 jest minimalnym serwerem "echo".
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // każde żądanie wywołuje funkcję handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler zwraca komponent the Path żądanego adresu URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

//!-
