// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Netcat jest prostym klientem odczytu/zapisu dla serwerów TCP.
package main

import (
	"io"
	"log"
	"net"
	"os"
)

//!+
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // UWAGA: ignorowanie błędów
		log.Println("zrobione")
		done <- struct{}{} // sygnalizowanie głównej funkcji goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // oczekiwanie na zakończenie funkcji goroutine działającej w tle
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
