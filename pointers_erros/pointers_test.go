package pointerserros

import "testing"

func TestWallet(t *testing.T) {

  assertBalance := func (t testing.TB, wallet Wallet, want Bitcoin)  {
    t.Helper()

    got := wallet.balance

    if got != want {
			t.Errorf("got %s want %s", got, want)
		}

  }

  assertErr := func ( t testing.TB, err error)  {
    t.Helper()

    if err != nil {
      t.Error(err)
    }
  }

  t.Run("deposit", func(t *testing.T) {
    wallet := Wallet{}

    wallet.Deposit(Bitcoin(10))
    assertBalance(t, wallet, Bitcoin(10))
  })

  t.Run("withdraw", func(t *testing.T) {
    wallet := Wallet{balance: Bitcoin(20)}

    wallet.Withdraw(Bitcoin(10))
    assertBalance(t, wallet, Bitcoin(10))
  })

  t.Run("withdraw insuffficient funds", func(t *testing.T) {
    startingBalance := Bitcoin(20)

    wallet := Wallet{startingBalance}

    err := wallet.Withdraw(Bitcoin(100))

    assertErr(t, err)
    assertBalance(t, wallet, startingBalance)
  })
}
