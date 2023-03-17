// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+test
package storage

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}

	const user = "joe@example.org"
	usage[user] = 980000000 // symulowanie sytuacji wykorzystania 980 MB

	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("funkcja notifyUser nie została wywołana")
	}
	if notifiedUser != user {
		t.Errorf("niewłaściwy użytkownik (%s) został powiadomiony, oczekiwany %s",
			notifiedUser, user)
	}
	const wantSubstring = "98% Twojego limitu"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("nieoczekiwany komunikat powiadomienia <<%s>>, "+
			"oczekiwany podłańcuch %q", notifiedMsg, wantSubstring)
	}
}

//!-test

/*
//!+defer
func TestCheckQuotaNotifiesUser(t *testing.T) {
	// Zapisanie i przywrócenie oryginalnej implementacji notifyUser.
	saved := notifyUser
	defer func() { notifyUser = saved }()

	// Instalowanie testowej atrapy implementacji notifyUser.
	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}
	// ...reszta testu...
}
//!-defer
*/
