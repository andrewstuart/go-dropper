package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/andrewstuart/dropper/ocean"
)

const NUM_TO_CREATE = 8

func main() {
	s := os.ExpandEnv("$HOME/.do-token")
	t, err := ReadToken(s)

	if err != nil {
		log.Fatal(err)
	}

	c := ocean.NewClient(t)

	// d := &ocean.Droplet{
	// 	Name:   "foo",
	// 	Region: "sfo1",
	// 	Size:   "512mb",
	// 	Image:  "lamp",
	// }

	// err = c.CreateDroplet(d)

	drops, err := c.GetDroplets()

	if err != nil {
		log.Fatal(err)
	}

	if len(drops) > 0 {
		resp, err := drops[0].Rename("the-droplet")

		if err != nil {
			log.Fatal(err)
		}

		enc := json.NewEncoder(os.Stdout)

		enc.Encode(resp)

	}

	// imgs, err := c.GetImages()

	// for _, i := range imgs {
	// 	enc.Encode(i)
	// }
}
