package main

import "log"

func main() {
	t, err := ReadToken("./.token")

	if err != nil {
		log.Fatal(err)
	}

	c := NewClient(t)

	c.GetRegions()
}
