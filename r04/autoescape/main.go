// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Autoescape demonstruje automatyczną ucieczkę HTML w html/template.
package main

import (
	"html/template"
	"log"
	"os"
)

//!+
func main() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))
	var data struct {
		A string        // niezaufany zwykły tekst
		B template.HTML // zaufany HTML
	}
	data.A = "<b>Witaj!</b>"
	data.B = "<b>Witaj!</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}

//!-
