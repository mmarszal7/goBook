// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


//!+

// Bzipper odczytuje dane wejściowe, kompresuje je za pomocą bzip2 i wypisuje je.
package main

import (
	"io"
	"log"
	"os"

	"code/r13/bzip"
)

func main() {
	w := bzip.NewWriter(os.Stdout)
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatalf("bzipper: %v\n", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("bzipper: zamykanie: %v\n", err)
	}
}

//!-
