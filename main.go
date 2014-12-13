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

	d := drops[0]

	log.Println(d.Id)

	d.Delete()
}
