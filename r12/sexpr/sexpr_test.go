// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"reflect"
	"testing"
)

// Test weryfikuje, czy kodowanie i dekodowanie złożonych wartości danych
// generują równoważny wynik.
//
// Test nie przyjmuje bezpośrednich założeń na temat zakodowanych danych wyjściowych,
// ponieważ zależą one od kolejności iterowania mapy, która jest niedeterministyczna
// Dane wyjściowe z instrukcji the t.Log można sprawdzić
// uruchamiając test z flagą -v:
//
// 	$ go test -v code/r12/sexpr
//
func Test(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr Strangelove",
		Subtitle: "Czyli jak przestałem się martwić i pokochałem bombę",
		Year:     1964,
		Actor: map[string]string{
			"Dr Strangelove":                  "Peter Sellers",
			"Kapitan Lionel Mandrake":         "Peter Sellers",
			"Prezydent Merkin Muffley":        "Peter Sellers",
			"Generał Buck Turgidson":          "George C. Scott",
			"Generał brygady Jack D. Ripper":  "Sterling Hayden",
			`Major T.J. "King" Kong`:          "Slim Pickens",
		},
		Oscars: []string{
			"Najlepszy aktor pierwszoplanowy (nominacja)",
			"Najlepszy scenariusz adaptopwany (nominacja)",
			"Najlepszy reżyser (nominacja)",
			"Najlepszy film (nominacja)",
		},
	}

	// Kodowanie
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Wykonanie Marshal nie powiodło się: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Dekodowanie
	var movie Movie
	if err := Unmarshal(data, &movie); err != nil {
		t.Fatalf("Wykonanie Unmarshal nie powiodło się: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Sprawdzanie równoważności.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("nie są równoważne")
	}

	// Pretty-print:
	data, err = MarshalIndent(strangelove)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = %s\n", data)
}
