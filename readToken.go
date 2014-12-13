package main

import (
	"io/ioutil"

	"github.com/andrewstuart/linedropper/ocean"
)

func ReadToken(fname string) (ocean.Token, error) {
	b, err := ioutil.ReadFile(fname)

	if err != nil {
		return "", err
	}

	return Token(string(b)), nil
}
