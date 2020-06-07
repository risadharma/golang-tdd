package main

// Err constant
const (
	ErrNotFound         = DictionaryErr("word unknown")
	ErrWordDoesNotExist = DictionaryErr("word does not exist")
	ErrWordExist        = DictionaryErr("word already exist")
)

type Dictionary map[string]string
type DictionaryErr string

func (de DictionaryErr) Error() string {
	return string(de)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExist
	}

	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case ErrNotFound:
		return ErrWordDoesNotExist
	}

	return nil
}

func main() {
}
