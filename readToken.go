package main

import "io/ioutil"

func ReadToken(fname string) (Token, error) {
	b, err := ioutil.ReadFile(fname)

	if err != nil {
		return "", err
	}

	return Token(string(b)), nil
}
