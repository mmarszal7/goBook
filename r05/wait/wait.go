// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Wait czeka, aż serwer HTTP zacznie odpowiadać.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//!+
// WaitForServer próbuje się połączyć z serwerem adresu URL.
// Próby są ponawiane przez minutę z wykorzystaniem algorytmu exponential back-off.
// Jeśli wszystkie próby zawiodą, raportowany jest błąd.
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // powodzenie
		}
		log.Printf("serwer nie odpowiada (%s); ponawianie...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("serwer %s nie odpowiedział po %s", url, timeout)
}

//!-

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "sposób użycia: wait url\n")
		os.Exit(1)
	}
	url := os.Args[1]
	//!+main
	// (W funkcji main.)
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Strona nie działa: %v\n", err)
		os.Exit(1)
	}
	//!-main
}
