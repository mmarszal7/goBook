// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Surface oblicza renderowanie SVG funkcji powierzchniowej 3D.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // rozmiar płótna w pikselach
	cells         = 100                 // liczba komórek siatki
	xyrange       = 30.0                // zakresy osi (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // liczba pikseli na jednostkę x lub y
	zscale        = height * 0.4        // liczba pikseli na jednostkę z 
	angle         = math.Pi / 6         // kąt nachylenia osi x, y (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Znajdowanie punktu (x, y) w rogu komórki (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Obliczenie wysokości z powierzchni.
	z := f(x, y)

	// Rzutowanie (x, y, z) izometrycznie na płótno 2D SVG (sx, sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // odległość od punktu (0,0)
	return math.Sin(r) / r
}

//!-
