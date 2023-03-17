// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

// ---- lexer ----

// Ten typ lexer jest podobny do opisanego w rozdziale 12.
type lexer struct {
	scan  scanner.Scanner
	token rune // bieżący token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

type lexPanic string

// describe zwraca łańcuch znaków opisujący bieżący token, do użycia w błędach.
func (lex *lexer) describe() string {
	switch lex.token {
	case scanner.EOF:
		return "end of file"
	case scanner.Ident:
		return fmt.Sprintf("identyfikator %s", lex.text())
	case scanner.Int, scanner.Float:
		return fmt.Sprintf("liczba %s", lex.text())
	}
	return fmt.Sprintf("%q", rune(lex.token)) // każda inna runa
}

func precedence(op rune) int {
	switch op {
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	}
	return 0
}

// ---- parser ----

// Parse parsuje wejściowy łańcuch znaków jako wyrażenie arytmetyczne.
//
//   expr = num                         literał liczbowy, np., 3.14159
//        | id                          nazwa zmiennej, np., x
//        | id '(' expr ',' ... ')'     wywołanie funkcji
//        | '-' expr                    operator jednoargumentowy (+-)
//        | expr '+' expr               operator binarny (+-*/)
//
func Parse(input string) (_ Expr, err error) {
	defer func() {
		switch x := recover().(type) {
		case nil:
			// nie ma paniki
		case lexPanic:
			err = fmt.Errorf("%s", x)
		default:
			// nieoczekiwana panika: przywrócenie stanu paniki.
			panic(x)
		}
	}()
	lex := new(lexer)
	lex.scan.Init(strings.NewReader(input))
	lex.scan.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats
	lex.next() // initial lookahead
	e := parseExpr(lex)
	if lex.token != scanner.EOF {
		return nil, fmt.Errorf("nieoczekiwane %s", lex.describe())
	}
	return e, nil
}

func parseExpr(lex *lexer) Expr { return parseBinary(lex, 1) }

// binary = unary ('+' binary)*
// parseBinary zatrzymuje się, gdy napotka
// operator o niższym priorytecie niż prec1.
func parseBinary(lex *lexer, prec1 int) Expr {
	lhs := parseUnary(lex)
	for prec := precedence(lex.token); prec >= prec1; prec-- {
		for precedence(lex.token) == prec {
			op := lex.token
			lex.next() // konsumuje operator
			rhs := parseBinary(lex, prec+1)
			lhs = binary{op, lhs, rhs}
		}
	}
	return lhs
}

// unary = '+' expr | primary
func parseUnary(lex *lexer) Expr {
	if lex.token == '+' || lex.token == '-' {
		op := lex.token
		lex.next() // konsumuje '+' or '-'
		return unary{op, parseUnary(lex)}
	}
	return parsePrimary(lex)
}

// primary = id
//         | id '(' expr ',' ... ',' expr ')'
//         | num
//         | '(' expr ')'
func parsePrimary(lex *lexer) Expr {
	switch lex.token {
	case scanner.Ident:
		id := lex.text()
		lex.next() // konsumuje Ident
		if lex.token != '(' {
			return Var(id)
		}
		lex.next() // konsumuje '('
		var args []Expr
		if lex.token != ')' {
			for {
				args = append(args, parseExpr(lex))
				if lex.token != ',' {
					break
				}
				lex.next() // konsumuje ','
			}
			if lex.token != ')' {
				msg := fmt.Sprintf("otrzymane %s, oczekiwane ')'", lex.token)
				panic(lexPanic(msg))
			}
		}
		lex.next() // konsumuje ')'
		return call{id, args}

	case scanner.Int, scanner.Float:
		f, err := strconv.ParseFloat(lex.text(), 64)
		if err != nil {
			panic(lexPanic(err.Error()))
		}
		lex.next() // konsumuje liczbę
		return literal(f)

	case '(':
		lex.next() // konsumuje ')'
		e := parseExpr(lex)
		if lex.token != ')' {
			msg := fmt.Sprintf("otrzymane %s, oczekiwane ')'", lex.describe())
			panic(lexPanic(msg))
		}
		lex.next() // konsumuje ')'
		return e
	}
	msg := fmt.Sprintf("nieoczekiwane %s", lex.describe())
	panic(lexPanic(msg))
}
