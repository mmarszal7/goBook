// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Cross wyświetla wartości GOOS i GOARCH dla docelowego systemu operacyjnego.
package main

import (
	"fmt"
	"runtime"
)

//!+
func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}

//!-
