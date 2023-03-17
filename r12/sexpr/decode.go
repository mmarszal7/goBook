// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Package sexpr zapewnia środki do konwertowania obiektów Go na
// S-wyrażenia i odwrotnie.
package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"text/scanner"
)

//!+Unmarshal
// Unmarshal parsuje dane S-wyrażenia i zapełnia zmienną,
// której adres znajduje się w różnym od nil wskaźniku out.
func Unmarshal(data []byte, out interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(bytes.NewReader(data))
	lex.next() // pobiera pierwszy token
	defer func() {
		// UWAGA: to nie jest przykład idealnej obsługi błędów.
		if x := recover(); x != nil {
			err = fmt.Errorf("błąd w %s: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
}

//!-Unmarshal

//!+lexer
type lexer struct {
	scan  scanner.Scanner
	token rune // bieżący token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want { // UWAGA: nie jest to przykład dobrej obsługi błędów
		panic(fmt.Sprintf("otrzymane %q, oczekiwane %q", lex.text(), want))
	}
	lex.next()
}

//!-lexer

// Funkcja read jest decoderem dla niewielkiego podzbioru poprawnych składniowo
// S-wyrażeń.  Dla zwięzłości przykładu wykorzystuje wiele wątpliwych skrótów.
//
// parser zakłada, że
// - dane wejsciowe S-wyrażenia są poprawne składniowo; nie ma kontroli błędów.
// - dane wejsciowe S-wyrażenia odpowiadają typowi zmiennej.
// - wszystkie liczby w danych wejściowych są nieujemnymi, dziesiętnymi liczbami całkowitymi.
// - wszystkie klucze w składni struktury ((klucz wartość) ...) są symbolami niecytowanymi.
// - dane wejściowe nie zawieraja list kropkowanych, takich jak (1 2 . 3).
// - dane wejściowe nie zawierają makr czytnika Lisp, takich jak 'x and #'x.
//
// logika refleksji zakłada, że 
// - v jest zawsze zmienną właściwego typu dla wartości
//   S-wyrażenia.  Przykładowo: v nie może być wartością logiczną,
//   interfejsem, kanałem lub funkcją, a jeśli v jest tablicą, dane
//   wejściowe muszą mieć prawidłową liczbę elementów.
// - v w najwyższego poziomu wywołaniu read ma wartość zerową
//   swojego typu i nie wymaga czyszczenia.
// - jeśli v jest zmienną liczbową, jest liczbą całkowitą ze znakiem.

//!+read
func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		// Jedynymi prawidłowymi identyfikatorami są 
		// "nil" oraz nazwy pól struktury.
		if lex.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text()) // UWAGA: ignorowanie błędów
		v.SetString(s)
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text()) // UWAGA: ignorowanie błędów
		v.SetInt(int64(i))
		lex.next()
		return
	case '(':
		lex.next()
		readList(lex, v)
		lex.next() // konsumuje ')'
		return
	}
	panic(fmt.Sprintf("nieoczekiwany token %q", lex.text()))
}

//!-read

//!+readlist
func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array: // (pozycja ...)
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}

	case reflect.Slice: // (pozycja ...)
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}

	case reflect.Struct: // ((nazwa wartość) ...)
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("otrzymano token %q, oczekiwana nazwa pola", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name))
			lex.consume(')')
		}

	case reflect.Map: // ((klucz wartość) ...)
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}

	default:
		panic(fmt.Sprintf("nie można zdekodować listy na %v", v.Type()))
	}
}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("koniec pliku")
	case ')':
		return true
	}
	return false
}

//!-readlist
