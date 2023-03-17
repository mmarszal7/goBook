// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Package tempconv wykonuje obliczenia temperatur w skalach Celsjusza i Fahrenheita.
package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

/*
//!+flagvalue
package flag

// Value jest interfejsem dla wartości przechowywanej we fladze.
type Value interface {
	String() string
	Set(string) error
}
//!-flagvalue
*/

//!+celsiusFlag
// *celsiusFlag spełnia warunki interfejsu flag.Value.
type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // nie potrzeba kontroli błędów
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("nieprawidłowa temperatura %q", s)
}

//!-celsiusFlag

//!+CelsiusFlag

// CelsiusFlag definiuje flagę Celsius z określoną nazwą, domyślną wartością
// oraz informacją o sposobie wykorzystania i zwraca adres zmiennej flagi.
// Argument flagi musi podawać ilość i jednostkę, np. "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

//!-CelsiusFlag
