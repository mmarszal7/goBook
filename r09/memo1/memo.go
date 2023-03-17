// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


//!+

// Package memo zapewnia memoizację funkcji typu Func.
// Ta memoizacja nie jest współbieżnie bezpieczna.
package memo

// Memo buforuje wyniki wywołania Func.
type Memo struct {
	f     Func
	cache map[string]result
}

// Func jest typem funkcji, która ma być zmemoizowana.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// UWAGA: to nie jest współbieżnie bezpieczne!
func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}

//!-
