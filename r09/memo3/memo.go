// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Package memo zapewnia współbieżnie bezpieczną memoizację funkcji typu Func.
// Żądania dla różnych kluczy są uruchamiane współbieżnie.
// Współbieżne żądania dla tego samego klucza powodują duplikowanie pracy.
package memo

import "sync"

type Memo struct {
	f     Func
	mu    sync.Mutex // strzeże cache
	cache map[string]result
}

type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

//!+

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)

		// Pomiędzy tymi dwiema sekcjami krytycznymi kilka funkcji goroutine może się ścigać,
		// aby obliczyć f(key) i zaktualizować mapę.
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}

//!-
