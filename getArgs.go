package main

import (
	"flag"
	"os"
)

var cmd string

var defaultName string

var region *string = flag.String("r", "sfo1", "The region you would like to use for your droplet")
var image *string = flag.String("i", "ubuntu-14-04-x64", "The image you would like to use for your droplet")
var size *string = flag.String("s", "512mb", "The size you would like to use for your droplet")
var name *string = flag.String("n", "foo", "The name you would like to use for your droplet")

var force *bool = flag.Bool("f", false, "Force the action to be performed")

var key *string = flag.String("k", "", "The name of the ssh key you want to use for your droplet")

func init() {
	flag.Parse()

	if len(os.Args) > 1 {
		cmd = flag.Arg(0)
	}
}
