package dictmaps

import "testing"

func TestSearch(t *testing.T) {
  dictionary := Dictionary{"test": "this is just a test"}

  t.Run("known word", func(t *testing.T) {

    got, _ := dictionary.Search("test")
    want := "this is just a test"

    assertStrings(t, got, want)
  })

  t.Run("unknown word", func(t *testing.T) {
    _, got := dictionary.Search("unknown")

    assertErrors(t, got, ErrNotFound)
  })
}

func TestAdd(t *testing.T) {
  dictionary := Dictionary{}

  want := "this is a new field"

  got, err := dictionary.Add("newfield", "this is a new field")
  if err != nil {
		t.Fatal("should find added word:", err)
	}

  assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
  t.Helper()

  if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertErrors(t testing.TB, got, want error) {
  t.Helper()

  if got != want {
    t.Errorf("got error %q want %q", got, want)
  }
}
