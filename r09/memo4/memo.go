// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Package memo zapewnia współbieżnie bezpieczną memoizację funkcji typu Func.
// Żądania dla różnych kluczy są wykonywane równolegle.
// Współbieżne żądania dla tego samego klucza blokują, dopóki pierwsze nie zostanie zakończone.
// Ta implementacja wykorzystuje Mutex.
package memo

import "sync"

// Func jest typem funkcji, która ma być zmemoizowana.
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

//!+
type entry struct {
	res   result
	ready chan struct{} // zamykany, gdy res jest gotowy
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex // strzeże cache
	cache map[string]*entry
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// To jest pierwsze żądanie dla tego klucza.
		// Ta funkcja goroutine staje się odpowiedzialna za obliczanie wartości i rozgłaszanie stanu gotowości.
		// wartości i rozgłaszanie stanu gotowości.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready) // rozgłasza stan gotowości
	} else {
		// To jest powtórzone żądanie dla tego klucza.
		memo.mu.Unlock()

		<-e.ready // czeka na stan gotowości
	}
	return e.res.value, e.res.err
}

//!-
