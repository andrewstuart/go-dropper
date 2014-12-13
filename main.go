package main

import (
	"log"
	"sync"
)

func main() {
	t, err := ReadToken("./.token")

	if err != nil {
		log.Fatal(err)
	}

	c := NewClient(t)

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		rs := c.GetRegions()
		log.Println(rs)
		wg.Done()
	}()

	go func() {
		is := c.GetImages()
		log.Println(is)
		wg.Done()
	}()

	wg.Wait()

}
