// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Toposort wyświetla węzły skierowanego grafu acyklicznego w porządku topologicznym.
package main

import (
	"fmt"
	"sort"
)

//!+table
// prereqs mapuje kursy informatyczne na ich wymagania wstępne.
var prereqs = map[string][]string{
	"algorytmy":                        {"struktury danych"},
	"rachunek różniczkowy i całkowy":   {"algebra liniowa"},

	"kompilatory": {
		"struktury danych",
		"języki formalne",
		"organizacja procesora",
	},

	"struktury danych":       {"matematyka dyskretna"},
	"bazy danych":            {"struktury danych"},
	"matematyka dyskretna":   {"wstęp do programowania"},
	"języki formalne":        {"matematyka dyskretna"},
	"sieci":                  {"systemy operacyjne"},
	"systemy operacyjne":     {"struktury danych", "organizacja procesora"},
	"języki programowania":   {"struktury danych", "organizacja procesora"},
}

//!-table

//!+main
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}

//!-main
