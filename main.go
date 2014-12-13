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

	d := &ocean.Droplet{
		Name:   "foo",
		Region: "sfo1",
		Size:   "512mb",
		Image:  "ubuntu-12-04-x64",
	}

	c.CreateDroplet(d)

	log.Println(d)
}
