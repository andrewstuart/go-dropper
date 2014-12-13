package main

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
)

type RegionSlug string

//A region
type Region struct {
	Name      string     `json:"name"`
	Slug      RegionSlug `json:"slug"`
	Features  []string   `json:"features"`
	Available bool       `json:"available"`
	Sizes     []string   `json:"sizes"`
}

//RegionResp is a wrapper for the region responses
type RegionResp struct {
	Regions []Region `json:"regions"`
}
