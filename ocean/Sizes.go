package ocean

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
)

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
	Size  *Size  `json:"size"`
}

//Get a list of sizes for the user
func (c *Client) GetSizes() ([]Size, error) {
	dec, err := c.doGet("sizes")

	if err != nil {
		return []Size{}, err
	}

	sr := &SizeResp{}

	dec.Decode(sr)

	return sr.Sizes, nil
}
