package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/andrewstuart/dropper/ocean"
)

var c *ocean.Client

var dropMap = make(map[string]*ocean.Droplet)

func init() {

	//Don't worry if there aren't any args.
	if len(flag.Args()) > 0 {

		s := os.ExpandEnv("$HOME/.do-token")
		t, err := ReadToken(s)

		if err != nil {
			log.Fatal(err)
		}

		c = ocean.NewClient(t)

		drops, err := c.GetDroplets()

		if err != nil {
			log.Fatal(err)
		}

		if err == nil && len(drops) > 0 {
			for i := range drops {
				drop := &drops[i]

				idSt := strconv.Itoa(drop.Id)

				dropMap[drop.Name] = drop
				dropMap[idSt] = drop
			}
		}
	}
}

func main() {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 1, 4, 1, ' ', 0)

	defer w.Flush()

	switch strings.ToLower(cmd) {
	case "log":
		actions, err := c.GetActionLog()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintln(w, "#\tType\tStatus\tStarted\tCompleted")
		for i := range actions {
			action := actions[i]

			fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n", i+1, action.Type, action.Status, action.StartedAt.Format(time.RFC3339Nano), action.CompletedAt.Format(time.RFC3339Nano))
		}

	case "rm":
		dropIdent := flag.Arg(1)

		if chosenDrop, exists := dropMap[dropIdent]; exists {
			err := chosenDrop.Delete()

			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("Successfully deleted droplet %d.\n", chosenDrop.Id)
			}
		} else {
			fmt.Println("Droplet with that ID/name doesn't exist.")
		}

		break
	case "create":
		d := &ocean.Droplet{
			Name:   *name,
			Region: ocean.Slug(*region),
			Size:   ocean.Slug(*size),
			Image:  ocean.Slug(*image),
		}

		if *key != "" {
			kMap := make(map[string]ocean.SSHKey)

			ks, err := c.GetSSHKeys()

			if err != nil {
				log.Fatal(err)
			}

			keysToUse := []ocean.Slug{}
			if *key == "*" {
				for _, acctKey := range ks {
					keysToUse = append(keysToUse, ocean.Slug(acctKey.Fingerprint))
				}
			} else {

				for _, k := range ks {
					kMap[k.Name] = k
				}

				keysToUse = append(keysToUse, ocean.Slug(kMap[*key].Fingerprint))

			}

			d.SshKeys = keysToUse
		}

		err := c.CreateDroplet(d)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Created droplet %s with id %d\n", d.Name, d.Id)
		break
	case "key":
		if len(flag.Args()) > 2 {
			name := flag.Arg(1)
			path := os.ExpandEnv(flag.Arg(2))

			k, err := ocean.ReadSSHKey(path, name)

			fmt.Println(k.PublicKey, k.Name)

			if err != nil {
				log.Fatal(err)
			}

			err = c.CreateSSHKey(k)

			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal("SSH requires a keyname and filename as an argument.")
		}
		break
	case "who":
		acct, err := c.GetAccount()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Account for %s\tDroplet limit: %d\n", acct.Email, acct.DropletLimit)
		break
	case "ls":

		if len(flag.Args()) > 1 {
			switch flag.Arg(1) {
			case "images":
				imgs, err := c.GetImages()

				if err != nil {
					log.Fatal(err)
				}

				fmt.Fprintln(w, "#\tDistro\tName\tSlug\tRegions")
				for i := range imgs {
					img := &imgs[i]

					_, err := fmt.Fprintf(w, "%d.\t%s\t%s\t%s\t%v\n", i+1, img.Distro, img.Name, img.Slug, img.Regions)

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

				fmt.Fprintln(w, "#\tName\tSizes\tFeatures")
				for i := range regs {
					r := &regs[i]
					fmt.Fprintf(w, "%d.\t%s\t%v\t%v\n", i+1, r.Name, r.Sizes, r.Features)
				}
				break
			case "keys":
				keys, err := c.GetSSHKeys()

				if err != nil {
					log.Fatal(err)
				}
				fmt.Fprintln(w, "#\tName\tFingerprint")
				for i := range keys {
					k := &keys[i]
					fmt.Fprintf(w, "%d\t%s\t%s", i+1, k.Name, k.Fingerprint)
				}
				break
			case "sizes":
				sizes, err := c.GetSizes()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Fprintln(w, "#\tSlug\tMemory\tVcpus\tHourly")
				for i := range sizes {
					s := &sizes[i]
					fmt.Fprintf(w, "%d\t%s\t%v\t%v\t%v\n", i+1, s.Slug, s.Memory, s.VCpus, s.PriceHourly)
				}
				break
			default:
				usage()
				break
			}

		} else {
			if len(dropMap) > 0 {
				byId := make(map[int]*ocean.Droplet)
				dropSlice := []*ocean.Droplet{}

				//Dedupe
				for _, d := range dropMap {
					if _, exists := byId[d.Id]; !exists {
						byId[d.Id] = d
						dropSlice = append(dropSlice, d)
					}
				}

				fmt.Fprintln(w, "#\tId\tName\tStatus\tSize\tipv4")

				//Print
				for i, d := range dropSlice {
					fmt.Fprintf(w, "%d.\t%d\t%s\t%s\t%s\t%v\n", i+1, d.Id, d.Name, d.Status, d.Size, d.Networks["v4"])
				}
			} else {
				fmt.Println("You don't have any droplets! Use 'dropper create' to make one.")
			}
		}
		break
	case "restart":
		if enough := checkArgs(1, "Please provide a droplet name or id to restart"); enough {
			dropName := flag.Arg(1)

			d := dropMap[dropName]

			_, err := d.Reboot()

			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("Restarted droplet %s(%d)\n", d.Name, d.Id)
			}
		}
		break

	case "rename":
		if enough := checkArgs(2, "Please provide a current name and new name to rename"); enough {
			dropName := flag.Arg(1)
			newName := flag.Arg(2)

			if d, exists := dropMap[dropName]; exists {
				_, err := d.Rename(newName)

				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("Renamed droplet %s(%d) to %s\n", dropName, d.Id, newName)
			} else {
				log.Fatal("A droplet with that name or id does not exist.\n")
			}
		}
		break

	case "shutdown":
		if enough := checkArgs(1, "Please provide a droplet name or id to shutdown"); enough {
			dropName := flag.Arg(1)

			var err error

			//Are we being asked to force the droplet off?
			if len(flag.Args()) > 2 || *force {
				if flag.Arg(2) == "-f" || *force {
					_, err = dropMap[dropName].PowerOff()
				}
			} else {
				_, err = dropMap[dropName].Shutdown()
			}

			if err != nil {
				log.Fatal(err)
			}
		}

	case "boot":
		if enough := checkArgs(1, "Please provide a droplet name or id to boot"); enough {
			dropName := flag.Arg(1)

			_, err := dropMap[dropName].Boot()

			if err != nil {
				log.Fatal(err)
			}
		}
		break

	case "snapshot":
		if enough := checkArgs(2, "Please provide a droplet name"); enough {
			dropName := flag.Arg(1)
			snapName := flag.Arg(2)

			_, err := dropMap[dropName].Snapshot(snapName)

			if err != nil {
				log.Fatal(err)
			}
		}
		break
	default:
		usage()
		break
	}
}

func usage() {
	cmds := []string{"who", "log", "ls", "  ls regions", "  ls images", "  ls sizes", "  ls keys", "create", "rm", "rename", "restart", "shutdown", "boot", "snapshot"}

	fmt.Println("Please use a valid dropper command:")

	for _, cmd := range cmds {
		fmt.Printf("\t%s\n", cmd)
	}

	os.Exit(1)
}

func checkArgs(n int, msg string) bool {
	if len(flag.Args()) > n {
		return true
	} else {
		fmt.Println(msg)
		return false
	}
}
