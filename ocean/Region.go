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

//A region
type Region struct {
	Name         string   `json:"name"`
	Slug         Slug     `json:"slug"`
	Features     []string `json:"features"`
	Available    bool     `json:"available"`
	Sizes        []Slug   `json:"sizes"`
	ImagesBySlug map[Slug]*Image
}

//RegionResp is a wrapper for the region responses
type RegionResp struct {
	Regions []Region `json:"regions"`
	Region  *Region  `json:"region"`
}

//Get a list of regions for the user
func (c *Client) GetRegions() ([]Region, error) {
	dec, err := c.doGet("regions")

	if err != nil {
		return []Region{}, err
	}

	rs := &RegionResp{}

	dec.Decode(rs)

	return rs.Regions, nil
}
