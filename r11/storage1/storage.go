// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/


// Package storage jest częścią hipotetycznego serwera przechowywania danych w chmurze.
//!+main
package storage

import (
	"fmt"
	"log"
	"net/smtp"
)

var usage = make(map[string]int64)

func bytesInUse(username string) int64 { return usage[username] }

// Konfiguracja nadawcy poczty e-mail.
// UWAGA: nigdy nie umieszczaj hasła w kodzie źródłowym!
const sender = "notifications@example.com"
const password = "correcthorsebatterystaple"
const hostname = "smtp.example.com"

const template = `Uwaga: wykorzystujesz %d bajtów przestrzeni,
%d%% Twojego limitu.`

func CheckQuota(username string) {
	used := bytesInUse(username)
	const quota = 1000000000 // 1 GB
	percent := 100 * used / quota
	if percent < 90 {
		return // OK
	}
	msg := fmt.Sprintf(template, used, percent)
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender,
		[]string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendMail(%s) nie powiodło się: %s", username, err)
	}
}

//!-main
