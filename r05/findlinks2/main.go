// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Findlinks2 wykonuje żądanie HTTP GET na każdym adresie URL, 
// parsuje wynik jako HTML, a następnie wyświetla znajdujące się w nim linki.
//
// Sposób użycia:
//	findlinks url ...
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// visit dołącza do zmiennej links każdy link znaleziony w n, a następnie zwraca wynik.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

//!+
func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

// findLinks wysyła żądanie HTTP GET dla adresu URL,
// parsuje odpowiedź jako HTML, a następnie wyodrębnia i zwraca linki.
func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("pobieranie %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsowanie %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

//!-
