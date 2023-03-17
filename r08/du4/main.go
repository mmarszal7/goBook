// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Du4 command computes the disk usage of the files in a directory.
package main

// Wariant du4 obejmuje anulowanie:
// po wciśnieciu przycisku Enter program zostaje szybko zakończony.

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

//!+1
var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

//!-1

func main() {
	// Określa początkowe katalogi.
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//!+2
	// Anulowanie trawersacji, gdy wykryte zostaną dane wyjściowe.
	go func() {
		os.Stdin.Read(make([]byte, 1)) // odczyt pojedynczego bajtu
		close(done)
	}()
	//!-2

	// Trawersuje każdy korzeń drzewa plików równolegle.
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// Wyświetla wyniki periodycznie.
	tick := time.Tick(500 * time.Millisecond)
	var nfiles, nbytes int64
loop:
	//!+3
	for {
		select {
		case <-done:
			// Osuszanie kanału fileSizes, aby umożliwić dokończenie wykonywania istniejącym funkcjom goroutine.
			for range fileSizes {
				// Nic nie rób.
			}
			return
		case size, ok := <-fileSizes:
			// ...
			//!-3
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

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d plików  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir rekurencyjnie przechodzi drzewo plików zakorzenione w dir
// i wysyła rozmiar każdego znalezionego pliku przez kanał fileSizes.
//!+4
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		// ...
		//!-4
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
		//!+4
	}
}

//!-4

var sema = make(chan struct{}, 20) // semafor zliczajacy, który ogranicza współbieżność

// dirents zwraca wpisy katalogu dir.
//!+5
func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // nabycie żetonu
	case <-done:
		return nil // anulowane
	}
	defer func() { <-sema }() // zwolnienie żetonu

	// ...odczyt katalogu...
	//!-5

	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	defer f.Close()

	entries, err := f.Readdir(0) // 0 => brak ograniczeń; odczytywanie wszystkich wpisów
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		// Nie zwracaj: Readdir może zwrócić częściowe wyniki.
	}
	return entries
}
