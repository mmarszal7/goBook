// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Du2 oblicza wykorzystanie dysku przez pliki w katalogu.
package main

// Wariant du2 wykorzystuje select i time.Ticker,
// aby wyświetlać podsumowania periodyczniem, jeśli ustawiona jest flaga -v.

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

//!+
var verbose = flag.Bool("v", false, "pokazuje rozszerzone komunikaty postępu")

func main() {
	// ...uruchamianie funkcji goroutine działającej w tle...

	//!-
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

	//!+
	// Wyświetla wyniki periodycznie.
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // kanał fileSizes został zamknięty
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes) // końcowe sumy
}

//!-

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d plików  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

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
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
