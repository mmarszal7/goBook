// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package bank implementuje współbieżnie bezpieczny bank z jednym kontem.
package bank

var deposits = make(chan int) // wysyłanie kwoty do wpłaty
var balances = make(chan int) // odbieranie salda

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int // zmienna balance jest zamknięta w funkcji goroutine teller
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // uruchomienie monitorującej funkcji goroutine
}

//!-
