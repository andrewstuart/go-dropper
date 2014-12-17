package ocean

//A region
type Region struct {
	Name         string   `json:"name"`
	Slug         Slug     `json:"slug"`
	Features     []string `json:"features"`
	Available    bool     `json:"available"`
	Sizes        []Slug   `json:"sizes"`
	ImagesBySlug map[Slug]*Image
}

//Get a list of regions for the user
func (c *Client) GetRegions() ([]Region, error) {
	dec, err := c.doGet("regions")

	if err != nil {
		return []Region{}, err
	}

	//Struct literal for brevity
	rs := &struct{ Regions []Region }{}

	dec.Decode(rs)

	return rs.Regions, nil
}
