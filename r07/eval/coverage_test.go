// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

import (
	"fmt"
	"math"
	"testing"
)

//!+TestCoverage
func TestCoverage(t *testing.T) {
	var tests = []struct {
		input string
		env   Env
		want  string // oczekiwany błąd z Parse/Check lub wynik z Eval
	}{
		{"x % 2", nil, "nieoczekiwane '%'"},
		{"!true", nil, "nieoczekiwane '!'"},
		{"log(10)", nil, `nieznana funkcja "log"`},
		{"sqrt(1, 2)", nil, "wywołanie sqrt ma argumentów 2, wymaga 1"},
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
	}

	for _, test := range tests {
		expr, err := Parse(test.input)
		if err == nil {
			err = expr.Check(map[Var]bool{})
		}
		if err != nil {
			if err.Error() != test.want {
				t.Errorf("%s: ma %q, oczekiwane %q", test.input, err, test.want)
			}
			continue
		}

		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		if got != test.want {
			t.Errorf("%s: %v => %s, oczekiwane %s",
				test.input, test.env, got, test.want)
		}
	}
}

//!-TestCoverage
