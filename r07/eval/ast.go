// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

// Expr jest wyrażeniem arytmetycznym.
type Expr interface {
	// Eval zwraca wartość tego wyrażenia Expr w środowisku env.
	Eval(env Env) float64
	// Check zgłasza błędy w tym wyrażeniu Expr i dodaje do zbioru swoje wartości Var.
	Check(vars map[Var]bool) error
}

//!+ast

// Var identyfikuje zmienną, np. x.
type Var string

// literal jest stałą liczbową, np. 3.141.
type literal float64

// unary reprezentuje wyrażenia operatora jednoargumentowego, np. –x.
type unary struct {
	op rune // one of '+', '-'
	x  Expr
}

// binary reprezentuje wyrażenie operatora binarnego, np. x+y.
type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr
}

// call reprezentuje wyrażenie wywołania funkcji, np. sin(x).
type call struct {
	fn   string // możliwe wartości: "pow", "sin", "sqrt"
	args []Expr
}

//!-ast
