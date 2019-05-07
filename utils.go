package main

import (
	"errors"
)

func SToB(a string) (b bool, err error) {
	if a == "true" {
		b = true
	} else if a == "false" {
		b = false
	} else {
		err = errors.New("string is not boolean")
	}

	return
}
