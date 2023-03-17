// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Package memo zapewnia współbieżnie bezpieczną, nieblokującą memoizację funkcji.
// Żądania dla różnych kluczy są wykonywane równolegle.
// Współbieżne żądania dla tego samego klucza blokują, dopóki pierwsze nie zostanie zakończone.
// Ta implementacja wykorzystuje monitorującą funkcję goroutine.
package memo

//!+Func

// Func jest typem funkcji, która ma być zmemoizowana.
type Func func(key string) (interface{}, error)

// result jest wynikiem wywołania Func.
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // zamykany, gdy res jest gotowy
}

//!-Func

//!+get

// request jest komunikatem żądania, aby funkcja Func została zastosowana do klucza.
type request struct {
	key      string
	response chan<- result // klient chce pojedynczy wynik
}

type Memo struct{ requests chan request }

// New zwraca memoizację funkcji f. Klienty muszą kolejno wywoływać Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

//!-get

//!+monitor

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// To jest pierwsze żądanie dla tego klucza.
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // wywołuje f(key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// Ewaluuje funkcję.
	e.res.value, e.res.err = f(key)
	// Rozgłasza stan gotowości.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Czeka na stan gotowości.
	<-e.ready
	// Wysyła wynik do klienta.
	response <- e.res
}

//!-monitor
