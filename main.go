package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"

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
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 1, 4, 1, ' ', 0)

	defer w.Flush()

	switch cmd {
	case "rm":
		dropMap := make(map[string]*ocean.Droplet)

		drops, err := c.GetDroplets()

		if err != nil {
			log.Fatal(err)
		}

		for i := range drops {
			d := &drops[i]

			dropMap[d.Name] = d
			dropMap[strconv.Itoa(d.Id)] = d
		}

		dropIdent := flag.Arg(1)

		if chosenDrop, exists := dropMap[dropIdent]; exists {
			chosenDrop.Delete()
		} else {
			log.Println("Droplet with that ID/name doesn't exist.")
		}

		break
	case "create":
		d := &ocean.Droplet{
			Name:   name,
			Region: ocean.RegionSlug(*region),
			Size:   ocean.SizeSlug(*size),
			Image:  ocean.ImageSlug(*image),
		}

		log.Println(d)

		err := c.CreateDroplet(d)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Created droplet %s with id %d", d.Name, d.Id)
		break
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

					_, err := fmt.Fprintf(w, "%d.\t%s\t%s\t%v\n", i+1, img.Name, img.Slug, img.Regions)

					if err != nil {
						log.Fatal(err)
					}
				}
				break
			case "regions":
				regs, err := c.GetRegions()
				if err != nil {
					log.Fatal(err)
				}

				for i := range regs {
					r := &regs[i]
					fmt.Fprintf(w, "%d.\t%s\t%v\t%v\t%v\n", i+1, r.Name, r.Images, r.Sizes, r.Features)
				}
				break
			case "sizes":
				sizes, err := c.GetSizes()
				if err != nil {
					log.Fatal(err)
				}
				for i := range sizes {
					s := &sizes[i]
					fmt.Fprintf(w, "%d\t%s\t%v\t%v\t%v\n", i+1, s.Slug, s.Memory, s.VCpus, s.PriceHourly)
				}
			}
		} else {

			drops, err := c.GetDroplets()

			if err != nil {
				log.Fatal(err)
			}

			for i := range drops {
				d := &drops[i]
				fmt.Fprintf(w, "%d.\t%s\t%s\t%v\n", i+1, d.Name, d.Size, d.Networks)
			}
		}
		break
	}
}
