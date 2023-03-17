// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bzip_test

import (
	"bytes"
	"compress/bzip2" // reader
	"io"
	"testing"

	"code/r13/bzip" // writer
)

func TestBzip2(t *testing.T) {
	var compressed, uncompressed bytes.Buffer
	w := bzip.NewWriter(&compressed)

	// Zapisuje powtarzający się komunikat w milionie kawałków,
	// kompresując jedną kopię, ale nie drugą.
	tee := io.MultiWriter(w, &uncompressed)
	for i := 0; i < 1000000; i++ {
		io.WriteString(tee, "witaj")
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	// Sprawdzanie rozmiaru skompresowanego strumienia.
	if got, want := compressed.Len(), 255; got != want {
		t.Errorf("1 million komunikatów witaj skompresowany do %d bajtów, oczekiwano %d", got, want)
	}

	// Dekompresja i porównanie z oryginałem.
	var decompressed bytes.Buffer
	io.Copy(&decompressed, bzip2.NewReader(&compressed))
	if !bytes.Equal(uncompressed.Bytes(), decompressed.Bytes()) {
		t.Error("dekompresja dała inny komunikat")
	}
}
