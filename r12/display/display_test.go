// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package display

import (
	"io"
	"net"
	"os"
	"reflect"
	"sync"
	"testing"

	"code/r07/eval"
)

// UWAGA: nie możemy użyć komentarzy !+..!-, żeby powybierać te fragmenty testów do książki,
// ponieważ niszczą one mechanizm funkcji Example,
// co wymaga, aby komentarze // Output znajdowały się
// na końcu funkcji.

func ExampleExpr() {
	e, _ := eval.Parse("sqrt(A / pi)")
	Display("e", e)
	// Output:
	// Display e (eval.call):
	// e.fn = "sqrt"
	// e.args[0].type = eval.binary
	// e.args[0].value.op = 47
	// e.args[0].value.x.type = eval.Var
	// e.args[0].value.x.value = "A"
	// e.args[0].value.y.type = eval.Var
	// e.args[0].value.y.value = "pi"
}

func ExampleSlice() {
	Display("slice", []*int{new(int), nil})
	// Output:
	// Display slice ([]*int):
	// (*slice[0]) = 0
	// slice[1] = nil
}

func ExampleNilInterface() {
	var w io.Writer
	Display("w", w)
	// Output:
	// Display w (<nil>):
	// w = invalid
}

func ExamplePtrToInterface() {
	var w io.Writer
	Display("&w", &w)
	// Output:
	// Display &w (*io.Writer):
	// (*&w) = nil
}

func ExampleStruct() {
	Display("x", struct{ x interface{} }{3})
	// Output:
	// Display x (struct { x interface {} }):
	// x.x.type = int
	// x.x.value = 3
}

func ExampleInterface() {
	var i interface{} = 3
	Display("i", i)
	// Output:
	// Display i (int):
	// i = 3
}

func ExamplePtrToInterface2() {
	var i interface{} = 3
	Display("&i", &i)
	// Output:
	// Display &i (*interface {}):
	// (*&i).type = int
	// (*&i).value = 3
}

func ExampleArray() {
	Display("x", [1]interface{}{3})
	// Output:
	// Display x ([1]interface {}):
	// x[0].type = int
	// x[0].value = 3
}

func ExampleMovie() {
	//!+movie
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	//!-movie
	//!+strangelove
	strangelove := Movie{
		Title:    "Dr Strangelove",
		Subtitle: "Czyli jak przestałem się martwić i pokochałem bombę",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr Strangelove":                  "Peter Sellers",
			"Kapitan Lionel Mandrake":         "Peter Sellers",
			"Prezydent Merkin Muffley":        "Peter Sellers",
			"Generał Buck Turgidson":          "George C. Scott",
			"Generał brygady Jack D. Ripper":  "Sterling Hayden",
			`Major T.J. "King" Kong`:          "Slim Pickens",
		},

		Oscars: []string{
			"Najlepszy aktor pierwszoplanowy (nominacja)",
			"Najlepszy scenariusz adaptopwany (nominacja)",
			"Najlepszy reżyser (nominacja)",
			"Najlepszy film (nominacja)",
		},
	}
	//!-strangelove
	Display("strangelove", strangelove)

	// Nie używamy komentarza Output:, ponieważ
	// wyświetlanie mapy jest niedeterministyczne.
	/*
		//!+output
		Display strangelove (display.Movie):
		strangelove.Title = "Dr Strangelove"
		strangelove.Subtitle = "Czyli jak przestałem się martwić i pokochałem bombę"
		strangelove.Year = 1964
		strangelove.Color = false
		strangelove.Actor["Generał Buck Turgidson"] = "George C. Scott"
		strangelove.Actor["Generał brygady Jack D. Ripper"] = "Sterling Hayden"
		strangelove.Actor["Major T.J. \"King\" Kong"] = "Slim Pickens"
		strangelove.Actor["Dr Strangelove"] = "Peter Sellers"
		strangelove.Actor["Kapitan Lionel Mandrake"] = "Peter Sellers"
		strangelove.Actor["Prezydent Merkin Muffley"] = "Peter Sellers"
		strangelove.Oscars[0] = "Najlepszy aktor pierwszoplanowy (nominacja)"
		strangelove.Oscars[1] = "Najlepszy scenariusz adaptopwany (nominacja)"
		strangelove.Oscars[2] = "Najlepszy reżyser (nominacja)"
		strangelove.Oscars[3] = "Najlepszy film (nominacja)"
		strangelove.Sequel = nil
		//!-output
	*/
}

// Ten test zapewnia, że program będzie zamykany bez awarii.
func Test(t *testing.T) {
	// Some other values (YMMV)
	Display("os.Stderr", os.Stderr)
	// Output:
	// Display os.Stderr (*os.File):
	// (*(*os.Stderr).file).fd = 2
	// (*(*os.Stderr).file).name = "/dev/stderr"
	// (*(*os.Stderr).file).nepipe = 0

	var w io.Writer = os.Stderr
	Display("&w", &w)
	// Output:
	// Display &w (*io.Writer):
	// (*&w).type = *os.File
	// (*(*(*&w).value).file).fd = 2
	// (*(*(*&w).value).file).name = "/dev/stderr"
	// (*(*(*&w).value).file).nepipe = 0

	var locker sync.Locker = new(sync.Mutex)
	Display("(&locker)", &locker)
	// Output:
	// Display (&locker) (*sync.Locker):
	// (*(&locker)).type = *sync.Mutex
	// (*(*(&locker)).value).state = 0
	// (*(*(&locker)).value).sema = 0

	Display("locker", locker)
	// Output:
	// Display locker (*sync.Mutex):
	// (*locker).state = 0
	// (*locker).sema = 0
	// (*(&locker)) = nil

	locker = nil
	Display("(&locker)", &locker)
	// Output:
	// Display (&locker) (*sync.Locker):
	// (*(&locker)) = nil

	ips, _ := net.LookupHost("golang.org")
	Display("ips", ips)
	// Output:
	// Display ips ([]string):
	// ips[0] = "173.194.68.141"
	// ips[1] = "2607:f8b0:400d:c06::8d"

	// Even metarecursion!  (YMMV)
	Display("rV", reflect.ValueOf(os.Stderr))
	// Output:
	// Display rV (reflect.Value):
	// (*rV.typ).size = 8
	// (*rV.typ).ptrdata = 8
	// (*rV.typ).hash = 871609668
	// (*rV.typ)._ = 0
	// ...

	// wskaźnik, który wskazuje samego siebie
	type P *P
	var p P
	p = &p
	if false {
		Display("p", p)
		// Output:
		// Display p (display.P):
		// ...stuck, no output...
	}

	// mapa, która zawiera samą siebie
	type M map[string]M
	m := make(M)
	m[""] = m
	if false {
		Display("m", m)
		// Output:
		// Display m (display.M):
		// ...stuck, no output...
	}

	// wycinek, który zawiera samego siebie
	type S []S
	s := make(S, 1)
	s[0] = s
	if false {
		Display("s", s)
		// Output:
		// Display s (display.S):
		// ...stuck, no output...
	}

	// lista powiązana, która zjada własny ogon
	type Cycle struct {
		Value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}
	if false {
		Display("c", c)
		// Output:
		// Display c (display.Cycle):
		// c.Value = 42
		// (*c.Tail).Value = 42
		// (*(*c.Tail).Tail).Value = 42
		// ...ad infinitum...
	}
}
