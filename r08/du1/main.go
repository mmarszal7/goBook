// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


//!+main

// Du1 oblicza wykorzystanie dysku przez pliki w katalogu.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	// Określa początkowe katalogi.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Trawersuje drzewo plików.
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// Wyświetla wyniki.
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d plików  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

//!-main

//!+walkDir

// walkDir rekurencyjnie przechodzi drzewo plików zakorzenione w dir
// i wysyła rozmiar każdego znalezionego pliku przez kanał fileSizes.
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents zwraca wpisy katalogu dir.
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

//!-walkDir

// Wariant du1 wykorzystuje dwie funkcje goroutines
// i wyświetla podsumowanie po każdym znalezionym pliku.
