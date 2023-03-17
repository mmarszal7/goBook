// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Package bank implementuje współbieżnie bezpieczny bank z jednym kontem.
package bank

//!+
var (
	sema    = make(chan struct{}, 1) // binarny semafor strzegący zmiennej balance
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // nabycie żetonu
	balance = balance + amount
	<-sema // zwolnienie żetonu
}

func Balance() int {
	sema <- struct{}{} // nabycie żetonu
	b := balance
	<-sema // zwolnienie żetonu
	return b
}

//!-
