// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Package cake zapewnia symulację
// współbieżnej cukierni z wieloma parametrami.
//
// Użyj poniższego polecenia, aby uruchomić benchmarki:
// 	$ go test -bench=. code/r08/cake
package cake

import (
	"fmt"
	"math/rand"
	"time"
)

type Shop struct {
	Verbose        bool
	Cakes          int           // liczba ciast do upieczenia
	BakeTime       time.Duration // czas pieczenia jednego ciasta
	BakeStdDev     time.Duration // standardowe odchylenie czasu pieczenia
	BakeBuf        int           // przerwy buforowe pomiędzy pieczeniem i lukrowaniem 
	NumIcers       int           // liczba cukierników wykonujących lukrowanie
	IceTime        time.Duration // czas lukrowania jednego ciasta
	IceStdDev      time.Duration // standardowe odchylenie czasu lukrowania
	IceBuf         int           // przerwy buforowe pomiędzy lukrowaniem i dekorowaniem napisami
	InscribeTime   time.Duration // czas dekorowania jednego ciasta
	InscribeStdDev time.Duration // standardowe odchylenie czasu dekorowania
}

type cake int

func (s *Shop) baker(baked chan<- cake) {
	for i := 0; i < s.Cakes; i++ {
		c := cake(i)
		if s.Verbose {
			fmt.Println("pieczenie", c)
		}
		work(s.BakeTime, s.BakeStdDev)
		baked <- c
	}
	close(baked)
}

func (s *Shop) icer(iced chan<- cake, baked <-chan cake) {
	for c := range baked {
		if s.Verbose {
			fmt.Println("lukrowanie", c)
		}
		work(s.IceTime, s.IceStdDev)
		iced <- c
	}
}

func (s *Shop) inscriber(iced <-chan cake) {
	for i := 0; i < s.Cakes; i++ {
		c := <-iced
		if s.Verbose {
			fmt.Println("dekorowanie", c)
		}
		work(s.InscribeTime, s.InscribeStdDev)
		if s.Verbose {
			fmt.Println("zakońćzone", c)
		}
	}
}

// Work uruchamia symulację 'runs' razy.
func (s *Shop) Work(runs int) {
	for run := 0; run < runs; run++ {
		baked := make(chan cake, s.BakeBuf)
		iced := make(chan cake, s.IceBuf)
		go s.baker(baked)
		for i := 0; i < s.NumIcers; i++ {
			go s.icer(iced, baked)
		}
		s.inscriber(iced)
	}
}

// work blokuje wywołujacą funkcję goroutine przez przedział czssu,
// zwykle rozkładajacy się wokół wartości d,
// która jest standardowym odchyleniem stddev.
func work(d, stddev time.Duration) {
	delay := d + time.Duration(rand.NormFloat64()*float64(stddev))
	time.Sleep(delay)
}
