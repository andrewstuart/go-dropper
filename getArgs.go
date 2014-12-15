package main

import (
	"flag"
	"os"
	"time"
)

var cmd string

var defaultName string

var name string

var nameDef = time.Now().Format(time.RFC3339)

var region *string = flag.String("r", "sfo1", "The region you would like to use for your droplet")
var image *string = flag.String("i", "ubuntu-14-04-x64", "The image you would like to use for your droplet")
var size *string = flag.String("s", "512mb", "The size you would like to use for your droplet")

func init() {
	flag.StringVar(&name, "n", nameDef, "The image you would like to use for your droplet")

	flag.Parse()

	if len(os.Args) > 1 {
		cmd = flag.Arg(0)
	}
}
