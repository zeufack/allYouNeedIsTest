package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
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

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds.Error())
		assertBalance(t, wallet, startingBalance)

	})

}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.balance
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()
	// if err == nil {
	// 	t.Error("wanted an error but didn't get one")
	// }
	if got == nil {
		t.Fatal("didn't get an error but want one")
	}

	if got.Error() != want {
		t.Errorf("got %q want %q", got, want)
	}
}
