package main

import (
	"fmt"
	"log"
	"os"

	"github.com/andrewstuart/dropper/ocean"
)

const NUM_TO_CREATE = 8

var c *ocean.Client

func init() {
	s := os.ExpandEnv("$HOME/.do-token")
	t, err := ReadToken(s)

	if err != nil {
		log.Fatal(err)
	}

	c = ocean.NewClient(t)
}

func main() {

	switch cmd {
	case "who":
		acct, err := c.GetAccount()

		if err != nil {
			log.Fatal(err)
		}

		log.Println(acct)
		break
	case "ls":

		if len(os.Args) > 2 {
			switch os.Args[2] {
			case "images":
				imgs, err := c.GetImages()

				if err != nil {
					log.Fatal(err)
				}

				for i := range imgs {
					img := &imgs[i]
					fmt.Printf("%d.\t%s (%s) - [%v]\n", i+1, img.Name, img.Slug, img.Regions)
				}
				break
			}
		} else {

			drops, err := c.GetDroplets()

			if err != nil {
				log.Fatal(err)
			}

			for i := range drops {
				d := &drops[i]
				fmt.Printf("%d.\t%s (%s) - %v\n", i+1, d.Name, d.Size, d.Networks)
			}
		}
		break
	}
}
