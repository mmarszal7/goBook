// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


//!+

package tempconv

// CToF konwertuje temperaturę w stopniach Celsjusza na stopnie Fahrenheita.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC konwertuje temperaturę w stopniach Fahrenheita na stopnie Celsjusza.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//!-
