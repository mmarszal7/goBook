// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package bzip zapewnia interfejs zapisujący, który używa kompresji bzip2 (bzip.org).
package bzip

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -L/usr/lib -lbz2
#include <bzlib.h>
int bz2compress(bz_stream *s, int action,
                char *in, unsigned *inlen, char *out, unsigned *outlen);
*/
import "C"

import (
	"io"
	"unsafe"
)

type writer struct {
	w      io.Writer // bazowy strumień wyjściowy
	stream *C.bz_stream
	outbuf [64 * 1024]byte
}

// NewWriter zwraca interfejs zapisujący dla strumieni skompresowanych za pomocą bzip2.
func NewWriter(out io.Writer) io.WriteCloser {
	const (
		blockSize  = 9
		verbosity  = 0
		workFactor = 30
	)
	// UWAGA: ten kod może nie działać w przyszłych wersjach Go.
	w := &writer{w: out, stream: new(C.bz_stream)}
	C.BZ2_bzCompressInit(w.stream, blockSize, verbosity, workFactor)
	return w
}

//!-

//!+write
func (w *writer) Write(data []byte) (int, error) {
	if w.stream == nil {
		panic("zamknięty")
	}
	var total int // zapisywanie nieskompresowanych bajtów

	for len(data) > 0 {
		inlen, outlen := C.uint(len(data)), C.uint(cap(w.outbuf))
		C.bz2compress(w.stream, C.BZ_RUN,
			(*C.char)(unsafe.Pointer(&data[0])), &inlen,
			(*C.char)(unsafe.Pointer(&w.outbuf)), &outlen)
		total += int(inlen)
		data = data[inlen:]
		if _, err := w.w.Write(w.outbuf[:outlen]); err != nil {
			return total, err
		}
	}
	return total, nil
}

//!-write

//!+close
// Close spłukuje skompresowane dane i zamyka strumień.
// Nie zamyka bazowego io.Writer.
func (w *writer) Close() error {
	if w.stream == nil {
		panic("zamknięty")
	}
	defer func() {
		C.BZ2_bzCompressEnd(w.stream)
		w.stream = nil
	}()
	for {
		inlen, outlen := C.uint(0), C.uint(cap(w.outbuf))
		r := C.bz2compress(w.stream, C.BZ_FINISH, nil, &inlen,
			(*C.char)(unsafe.Pointer(&w.outbuf)), &outlen)
		if _, err := w.w.Write(w.outbuf[:outlen]); err != nil {
			return err
		}
		if r == C.BZ_STREAM_END {
			return nil
		}
	}
}

//!-close
