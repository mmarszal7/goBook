// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

import (
	"fmt"
	"math"
	"testing"
)

//!+Eval
func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
		//!-Eval
		// Dodatkowe testy, które nie pojawiają się w książce.
		{"-1 + -x", Env{"x": 1}, "-2"},
		{"-1 - x", Env{"x": 1}, "-2"},
		//!+Eval
	}
	var prevExpr string
	for _, test := range tests {
		// Wyswietla expr tylko, gdy się zmienia.
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // błąd parsowania
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() w %s = %q, oczekiwane %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}

//!-Eval

/*
//!+output
sqrt(A / pi)
	map[A:87616 pi:3.141592653589793] => 167

pow(x, 3) + pow(y, 3)
	map[x:12 y:1] => 1729
	map[x:9 y:10] => 1729

5 / 9 * (F - 32)
	map[F:-40] => -40
	map[F:32] => 0
	map[F:212] => 100
//!-output

// Dodatkowe dane wyjściowe, które nie pojawiają się w książce.

-1 - x
	map[x:1] => -2

-1 + -x
	map[x:1] => -2
*/

func TestErrors(t *testing.T) {
	for _, test := range []struct{ expr, wantErr string }{
		{"x % 2", "nieoczekiwane '%'"},
		{"math.Pi", "nieoczekiwane '.'"},
		{"!true", "nieoczekiwane '!'"},
		{`"witaj"`, "nieoczekiwane '\"'"},
		{"log(10)", `nieznana funkcja "log"`},
		{"sqrt(1, 2)", "wywołanie sqrt ma argumentów 2 , wymaga 1"},
	} {
		expr, err := Parse(test.expr)
		if err == nil {
			vars := make(map[Var]bool)
			err = expr.Check(vars)
			if err == nil {
				t.Errorf("nieoczekiwane powodzenie: %s", test.expr)
				continue
			}
		}
		fmt.Printf("%-20s%v\n", test.expr, err) // (for book)
		if err.Error() != test.wantErr {
			t.Errorf("ma error %s, oczekiwane %s", err, test.wantErr)
		}
	}
}

/*
//!+errors
x % 2               nieoczekiwane '%'
math.Pi             nieoczekiwane '.'
!true               nieoczekiwane '!'
"hello"             nieoczekiwane '"'

log(10)             nieznana funkcja "log"
sqrt(1, 2)          wywołanie sqrt ma argumentów 2 , wymaga 1
//!-errors
*/
