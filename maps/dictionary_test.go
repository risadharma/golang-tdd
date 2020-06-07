package main

import (
	"testing"
)

func TestDictionaryErr(t *testing.T) {
	sampleDictionaryErr := DictionaryErr("this is sample error")
	assertStrings(t, sampleDictionaryErr.Error(), "this is sample error")
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("baba")
		assertError(t, err, ErrNotFound)

	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is just a test",
	}

	assertTest := map[string]struct {
		word       string
		definition string
		err        error
	}{
		"Add success":       {"baba", "father", nil},
		"Add existing word": {"test", "this is just a test", ErrWordExist},
	}

	for name, tt := range assertTest {
		t.Run(name, func(t *testing.T) {
			err := dictionary.Add(tt.word, tt.definition)
			assertError(t, err, tt.err)
			assertDefinition(t, dictionary, tt.word, tt.definition)
		})
	}

}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)
		assertError(t, err, nil)

		newDefinition := "this is just a test that updated as new definition"

		err = dictionary.Update(word, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "baba"
		definition := "father"

		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("success delete", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary.Add(word, definition)

		_, err := dictionary.Search(word)
		assertError(t, err, nil)

		err = dictionary.Delete(word)
		assertError(t, err, nil)

		_, err = dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})

	t.Run("word not exist", func(t *testing.T) {
		word := "test"
		err := dictionary.Delete(word)
		assertError(t, err, ErrWordDoesNotExist)
	})

}

func assertDefinition(t *testing.T, dicitionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dicitionary.Search(word)
	if err != nil {
		t.Errorf("got %v, want %v", err, nil)
	}

	assertStrings(t, got, definition)
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
