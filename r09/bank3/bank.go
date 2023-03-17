// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Package bank implementuje współbieżnie bezpieczny bank z jednym kontem.
package bank

//!+
import "sync"

var (
	mu      sync.Mutex // strzeże zmiennej balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

//!-
