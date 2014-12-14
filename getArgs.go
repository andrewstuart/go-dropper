package main

import "os"

var cmd string

func init() {
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}
}
