// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Package geometry definiuje proste typy dla geometrii płaskiej.
//!+point
package geometry

import "math"

type Point struct{ X, Y float64 }

// Tradycyjna funkcja.
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// To samo, ale jako metoda typu Point.
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//!-point

//!+path

// Path to ścieżka łącząca punkty za pomocą linii prostych.
type Path []Point

// Distance zwraca odległość pokonaną wzdłuż ścieżki.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

//!-path
