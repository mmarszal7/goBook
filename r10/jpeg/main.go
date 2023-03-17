// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


//!+main

// Polecenie jpeg odczytuje obraz PNG ze standardowego strumienia wejściowego
// i zapisuje go jako obraz JPEG do standardowego strumienia wyjściowego.
package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png" // rejestrowanie dekodera PNG
	"io"
	"os"
)

func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Format wejściowy =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

//!-main

/*
//!+with
$ go build code/r03/mandelbrot
$ go build code/r10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
Format wejściowy = png
//!-with

//!+without
$ go build code/r10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
jpeg: image: unknown format
//!-without
*/
