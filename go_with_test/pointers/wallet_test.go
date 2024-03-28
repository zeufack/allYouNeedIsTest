package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.balance
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		// got := wallet.Balance()
		// fmt.Printf("address of balance in test is %p \n", &wallet.balance)
		// want := Bitcoin(10)

		// if got != want {
		// 	t.Errorf("got %s want %s", got, want)
		// }
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		// got := wallet.Balance()

		// want := Bitcoin(10)

		// if got != want {
		// 	t.Errorf("got %s want %s", got, want)
		// }
		assertBalance(t, wallet, Bitcoin(10))

	})

}
