package main

import (
	"log"
	"os"
	"runtime"

	"github.com/andrewstuart/dropper/ocean"
)

const NUM_TO_CREATE = 8

func main() {
	s := os.ExpandEnv("$HOME/.do-token")
	t, err := ReadToken(s)

	runtime.GOMAXPROCS(8)

	if err != nil {
		log.Fatal(err)
	}

	c := ocean.NewClient(t)

	a, err := c.GetAccount()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(a)

	// imgs, err := c.GetImages()

	// for _, i := range imgs {
	// 	enc.Encode(i)
	// }
}
