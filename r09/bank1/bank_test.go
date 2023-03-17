// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank_test

import (
	"fmt"
	"testing"

	"code/r09/bank1"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alicja
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Robert
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Oczekiwanie na obie transakcje.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
