// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Surface wyświetla wykres powierzchniowy 3D funkcji dostarczonej przez użytkownika.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

//!+parseAndCheck
import "code/r07/eval"

//!-parseAndCheck

// -- skopiowane z code/r03/surface --

const (
	width, height = 600, 320            // rozmiar płótna w pikselach
	cells         = 100                 // liczba komórek siatki
	xyrange       = 30.0                // zakresy osi (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // liczba pikseli na jednostkę x lub y
	zscale        = height * 0.4        // liczba pikseli na jednostkę z 
)

var sin30, cos30 = 0.5, math.Sqrt(3.0 / 4.0) // sin(30°), cos(30°)

func corner(f func(x, y float64) float64, i, j int) (float64, float64) {
	// Znajdowanie punktu (x, y) w rogu komórki (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y) // obliczenie wysokości z powierzchni.

	// Rzutowanie (x, y, z) izometrycznie na płótno 2D SVG (sx, sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func surface(w io.Writer, f func(x, y float64) float64) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(f, i+1, j)
			bx, by := corner(f, i, j)
			cx, cy := corner(f, i, j+1)
			dx, dy := corner(f, i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

// -- główny kod dla code/r07/surface --

//!+parseAndCheck
func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("puste wyrażenie")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("niezdefiniowana zmienna: %s", v)
		}
	}
	return expr, nil
}

//!-parseAndCheck

//!+plot
func plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, func(x, y float64) float64 {
		r := math.Hypot(x, y) // odległość od punktu (0,0)
		return expr.Eval(eval.Env{"x": x, "y": y, "r": r})
	})
}

//!-plot

//!+main
func main() {
	http.HandleFunc("/plot", plot)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main
