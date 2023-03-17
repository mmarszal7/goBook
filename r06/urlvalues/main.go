// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Polecenie urlvalues demonstruje typ mapy z metodami.
package main

/*
//!+values
package url

// Values mapuje klucz będący łańcuchem znaków na listę wartości.
type Values map[string][]string

// Get zwraca pierwszą wartość powiązaną z podanym kluczem 
// lub pusty łańcuch "", jeśli nie ma żadnej wartości.
func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

// Add dodaje wartość do klucza.
// Dołącza do wszystkich istniejących wartości powiązanych z danym kluczem.
func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}
//!-values
*/

import (
	"fmt"
	"net/url"
)

func main() {
	//!+main
	m := url.Values{"lang": {"en"}} // konstrukcja bezpośrednia
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang")) // "en"
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // "1"      (pierwsza wartość)
	fmt.Println(m["item"])     // "[1 2]"  (bezpośredni dostęp do mapy)

	m = nil
	fmt.Println(m.Get("item")) // ""
	m.Add("item", "3")         // panic: przypisanie do wpisu w mapie nil
	//!-main
}
