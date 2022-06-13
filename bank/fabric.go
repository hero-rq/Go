package mydict

import (
	"errors"
	)

type Dictionary map[string] string 

var (
	NotFound = errors.New("Not found")
	errCantUpdate = errors.New("Can't update nonexistsing word")
	errWordExists = errors.New("That word already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	value, exits := d[word]
	if exits {
		return value, nil
	} 
	return "", NotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == NotFound {
		d[word] = def
	} else if err == nil{
		return errWordExists
	}
	return nil
}


func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch  err {
	case nil:
		d[word] = definition
	case NotFound : 
		return errCantUpdate	
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
