package main

import (
	"log"

	"github.com/andrewstuart/linedropper/ocean"
)

func main() {
	t, err := ReadToken("./.token")

	if err != nil {
		log.Fatal(err)
	}

	c := ocean.NewClient(t)

	drops := c.GetDroplets()

	drops[0].Delete()
}
