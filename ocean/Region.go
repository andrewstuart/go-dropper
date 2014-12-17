package ocean

import "fmt"

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
		return nil, fmt.Errorf("Error retreiving regions from DO:\n\t%v", err)
	}

	//Struct literal for brevity
	rs := &struct{ Regions []Region }{}

	err = dec.Decode(rs)

	if err != nil {
		return nil, fmt.Errorf("Error decoding regions:\n\t%v", err)
	}

	return rs.Regions, nil
}
