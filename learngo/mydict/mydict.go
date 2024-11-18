package mydict

import (
	"errors"
)

// Dictionary type
type Dictionary map[string]string

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]

	if exists {
		return value, nil
	}
	return "", errors.New("Not Found")
}

func (d Dictionary) Add(word, value string) error {

	_, exists := d[word]
	if exists {
		return errors.New("added values is existed in dictionary")
	}

	d[word] = value
	return nil

}
func (d Dictionary) Update(word, value string) error {
	_, exists := d[word]
	if !(exists) {
		return errors.New("updated values is not existed in dictionary please add it")
	}
	d[word] = value
	return nil
}
func (d Dictionary) Delete(word string) error {

	_, exists := d[word]
	if !(exists) {
		return errors.New("deleted values is not existed in dictionary")
	}
	delete(d, word)
	return nil
}
