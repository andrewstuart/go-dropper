package main

type SizeSlug string

type Size struct {
	Slug         SizeSlug     `json:"slug"`
	Memory       int64        `json:"memory"`
	VCpus        int          `json:"vcpus"`
	Disk         int64        `json:"disk"`
	Transfer     int          `json:"transfer"`
	PriceMonthly float64      `json:"price_monthly"`
	PriceHourly  float64      `json:"price_hourly"`
	Regions      []RegionSlug `json:"regions"`
}

type SizeResp struct {
	Sizes []Size `json:"sizes"`
}
