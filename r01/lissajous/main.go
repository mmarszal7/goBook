// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Uruchamianie z argumentem "web" wiersza poleceń dla serwera WWW.
//!+main

// Lissajous generuje animacje GIF losowych figur Lissajous.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

//!-main
// Pakiety niewymgane przez wersję z książki.
import (
	"log"
	"net/http"
	"time"
)

//!+main

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // pierwszy kolor w zmiennej palette
	blackIndex = 1 // następny kolor w zmiennej palette
)

func main() {
	//!-main
	// SEkwencja obrazów jest deterministyczna, chyba że inicjujemy
	// generator liczb pseudolosowych wykorzystując bieżący czas.
	// Dziękujemy Randallowi McPhersonowi za wskazanie tego pominięcia.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // liczba pełnych obiegów oscylatora x
		res     = 0.001 // rozdzielczość kątowa
		size    = 100   // rozmiar płótna obrazu [-size..+size]
		nframes = 64    // liczba klatek animacji
		delay   = 8     // opóźnienie między klatkami w jednostkach 10 ms
	)
	freq := rand.Float64() * 3.0 // częstotliwość względna oscylatora y 
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // przesunięcie fazowe
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // UWAGA: ignorowanie błędów kodowania
}

//!-main
