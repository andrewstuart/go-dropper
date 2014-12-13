package main

type DropletSize string

type Droplet struct {
	Name   string      `json:"name"`
	Region RegionSlug  `json:"region"`
	Size   DropletSize `json:"size"`
}
