package pointersErrors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit test", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		fmt.Printf("address of balance in test is %v \n", &wallet.balance)
		want := Bitcoin(10)
		assertEqual(t, wallet, want)
	})

	t.Run("withdraw test", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		got := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)
		assertNoError(t, got)
		assertEqual(t, wallet, want)
	})

	t.Run("insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(100)
		assertError(t, err, ErrInsufficientFunds)
		assertEqual(t, wallet, startingBalance)

	})
}

func assertEqual(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
