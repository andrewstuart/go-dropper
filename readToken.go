package main

import (
	"io/ioutil"

	"github.com/andrewstuart/dropper/ocean"
)

func ReadToken(fname string) (ocean.Token, error) {
	b, err := ioutil.ReadFile(fname)

	if err != nil {
		return "", err
	}

	return ocean.Token(string(b)), nil
}
