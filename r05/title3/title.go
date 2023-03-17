// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Title3 wyświetla tytuł dokumentu HTML określonego poprzez adres URL.
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// Skopiowane z code/r05/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

//!+
// soleTitle zwraca tekst pierwszego niepustego elementu title w dokumencie doc 
// oraz error, jeśli nie było dokładnie jednego.
func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			// nie ma paniki
		case bailout{}:
			// "oczekiwana" panika
			err = fmt.Errorf("wiele elementów title")
		default:
			panic(p) // nieoczekiwana panika; kontynuowanie paniki
		}
	}()

	// Wyjście z rekurencji, jeśli znaleziony zostanie więcej niż jeden niepusty element title.
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // wiele elementów title 
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("brak elementu title")
	}
	return title, nil
}

//!-

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	// Sprawdza, czy Content-Type to HTML (np. "text/html; charset=utf-8").
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s ma typ %s, a nie text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsowanie %s jako HTML: %v", url, err)
	}
	title, err := soleTitle(doc)
	if err != nil {
		return err
	}
	fmt.Println(title)
	return nil
}

func main() {
	for _, arg := range os.Args[1:] {
		if err := title(arg); err != nil {
			fmt.Fprintf(os.Stderr, "tytuł: %v\n", err)
		}
	}
}
