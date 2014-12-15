package ocean

type Size struct {
	Slug         Slug    `json:"slug"`
	Memory       int64   `json:"memory"`
	VCpus        int     `json:"vcpus"`
	Disk         int64   `json:"disk"`
	Transfer     int     `json:"transfer"`
	PriceMonthly float64 `json:"price_monthly"`
	PriceHourly  float64 `json:"price_hourly"`
	Regions      []Slug  `json:"regions"`
}

type SizeResp struct {
	Sizes []Size `json:"sizes"`
}
