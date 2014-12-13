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

	wg.Add(3)

	go func() {
		rs := c.GetRegions()
		log.Println(rs)
		wg.Done()
	}()

	go func() {
		is := c.GetImages()

		for _, i := range is {
			log.Println(i.Name)
		}

		wg.Done()
	}()

	go func() {
		sz := c.GetSizes()

		log.Println(sz)
		wg.Done()
	}()

	wg.Wait()

}
