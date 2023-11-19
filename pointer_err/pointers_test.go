package pointererr

import "testing"

func TestWallet(t *testing.T) {

  assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin)  {
    t.Helper()
    got := wallet.Balance()
    if got != want {
      t.Errorf("got %s want %s", got, want)
    }
  }

  assertError := func(t testing.TB, got error, want error)  {
    t.Helper()

    if got == nil {
      t.Fatal("didn't get an error but wanted one")
    }

    if got != want {
      t.Errorf("got %q, want %q", got, want)
    }
  }

  t.Run("deposit", func(t *testing.T) {
    wallet := Wallet{}
    wallet.Deposit(Bitcoin(10))
    want := Bitcoin(10)
    assertBalance(t, wallet, want)
  })

  t.Run("withdraw", func(t *testing.T) {
    wallet := Wallet{balance: Bitcoin(20)}
    err := wallet.Withdraw(Bitcoin(10))
    want := Bitcoin(10)
    assertError(t, err, ErrInsufficientFunds)
    assertBalance(t, wallet, want)
  })

  t.Run("withdraw insufficient balance", func(t *testing.T) {
    initialBalance := Bitcoin(20)
    wallet := Wallet{balance: initialBalance}
    err := wallet.Withdraw(Bitcoin(100))
    assertError(t, err, ErrInsufficientFunds)
    assertBalance(t, wallet, initialBalance)
  })
}
