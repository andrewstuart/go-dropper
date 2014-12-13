package main

import (
	"encoding/json"
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

	imgs, err := c.GetImages()

	enc := json.NewEncoder(os.Stdout)

	for _, i := range imgs {
		enc.Encode(i)
	}
}
