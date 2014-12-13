package main

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
)

//A region
type Region struct {
	Name      string   `json:"name"`
	Slug      string   `json:"slug"`
	Features  []string `json:"features"`
	Available bool     `json:"available"`
	Sizes     []string `json:"sizes"`
}

// func (r Region) String() string {
// 	avail := ""
// 	if r.Available {
// 		avail = "Not available"
// 	} else {
// 		avail = "Available"
// 	}
// 	return fmt.Sprintf("%s\t(%s):\t%s", r.Name, r.Slug, avail)
// }

//RegionResp is a wrapper for the region responses
type RegionResp struct {
	Regions []Region `json:"regions"`
}

// func (r RegionResp) String() string {
// 	s := ""
// 	for _, r := range r.Regions {
// 		s += fmt.Sprintf("%s\n", r)
// 	}

// 	return s
// }
