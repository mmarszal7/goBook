// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Chat jest serwerem, który pozwala klientom rozmawiać ze sobą.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//!+broadcaster
type client chan<- string // kanał komunikatów wychodzących

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // wszystkie przychodzące komunikaty klientów
)

func broadcaster() {
	clients := make(map[client]bool) // wszystkie podłączone klienty
	for {
		select {
		case msg := <-messages:
			// Rozgłasza przychodzące komunikaty do kanałów
			// komunikatów wychodzących wszystkich klientów.
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // wychodzące komunikaty klienta
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "Jesteś " + who
	messages <- who + " przybył"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// UWAGA: ignorowanie potencjalnych błędów z input.Err().

	leaving <- ch
	messages <- who + " odszedł"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // UWAGA: ignorowanie błędów sieciowych
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
