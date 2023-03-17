// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Issuesreport wyświetla raport na temat tematów odpowiadajacych kryteriom wyszukiwania.
package main

import (
	"log"
	"os"
	"text/template"
	"time"

	"code/r04/github"
)

//!+template
const templ = `Liczba znalezionych tematów {{.TotalCount}}:
{{range .Items}}----------------------------------------
Numer:        {{.Number}}
Użytkownik:   {{.User.Login}}
Tytuł:        {{.Title | printf "%.64s"}}
Utworzony:    {{.CreatedAt | daysAgo}} dni temu
{{end}}`

//!-template

//!+daysAgo
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

//!-daysAgo

//!+exec
var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

//!-exec

func noMust() {
	//!+parse
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}
	//!-parse
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

/*
//!+output
$ go build code/r04/issuesreport
$ ./issuesreport repo:golang/go is:open json decoder
Liczba znalezionych tematów 13:
----------------------------------------
Numer:        5680
Użytkownik:   eaigner
Tytuł:        encoding/json: set key converter on en/decoder
Utworzony:    750 dni temu 
----------------------------------------
Numer:        6050
Użytkownik:   gopherbot
Tytuł:        encoding/json: provide tokenizer
Utworzony:    695 days
----------------------------------------
...
//!-output
*/
