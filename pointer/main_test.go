package main

import (
	"fmt"
	"testing"
)

func TestBitcoinString(t *testing.T) {
	btc := Bitcoin(1)

	got := fmt.Sprintf("%s", btc)
	want := "1 BTC"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))
		assertBalance(t, wallet, Bitcoin(20))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(50)}
		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(40))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(50))

		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, ErrInsufficientBalance)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("%#v got %s, want %s", wallet, got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t *testing.T, err error, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if err != want {
		t.Errorf("got %q, want %q", err.Error(), want)
	}
}
