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
    _, got := dictionary.Search("mehh")

    assertErrors(t, got, ErrNotFound)
  })
}

func TestAdd(t *testing.T) {
  t.Run("new word", func(t *testing.T) {
    dictionary := Dictionary{}
    word := "new"
    definition := "this is a new element"

    err := dictionary.Add(word, definition)
    assertErrors(t, err, nil)
    assertDefinition(t, dictionary, word, definition)
  })

  t.Run("existing word", func(t *testing.T) {
    word := "new"
    definition := "this is a new element"
    dictionary := Dictionary{word: definition}

    err := dictionary.Add(word, "this is just a test")

    assertErrors(t, err, ErrWordExists)
    assertDefinition(t, dictionary, word, definition)
  })
}

func TestUpdate(t *testing.T) {

  t.Run("existing word", func(t *testing.T) {
    word := "testing"
    definition := "this is just a test"
    dictionary := Dictionary{word: definition}
    newDefinition := "new definition"

    err := dictionary.Update(word, newDefinition)

    assertErrors(t, err, nil)
    assertDefinition(t, dictionary, word, newDefinition)
  })

  t.Run("new word", func(t *testing.T) {
    word := "test"
    definition := "this is just a test"
    dictionary := Dictionary{}

    err := dictionary.Update(word, definition)
    assertErrors(t, err, ErrWordDoesNotExist)
  })
}

func TestDelete(t *testing.T) {
  t.Run("existing word", func(t *testing.T) {
    word := "test"
    definition := "this is a new word to be deleted"
    dictionary := Dictionary{word: definition}
  
    err := dictionary.Delete(word)
    assertErrors(t, err, nil)
  })

}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
  t.Helper()

  got, err := dict.Search(word)

  if err != nil {
		t.Fatal("should find added word:", err)
	}

  assertStrings(t, got, definition)
}

func assertErrors(t testing.TB, got, want error) {
  t.Helper()

  if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got string, want string) {
  t.Helper()

  if got != want {
    t.Errorf("got %q expected %q given %q", got, want, "test")
  }
}
