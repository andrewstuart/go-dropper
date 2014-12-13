package main

type ImageSlug string

type Image struct {
	Id      int          `json:"id"`
	Name    string       `json:"name"`
	Distro  string       `json:"distribution"`
	Slug    ImageSlug    `json:"slug"`
	Public  bool         `json:"public"`
	Regions []RegionSlug `json:"regions"`
}

type ImageResp struct {
	Images []Image `json:"images"`
}
