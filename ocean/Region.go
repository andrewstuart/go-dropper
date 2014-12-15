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
	Name      string            `json:"name"`
	Slug      Slug              `json:"slug"`
	Features  []string          `json:"features"`
	Available bool              `json:"available"`
	Sizes     []Slug            `json:"sizes"`
	Images    map[string]*Image `json:"images,omitempty"`
}

//RegionResp is a wrapper for the region responses
type RegionResp struct {
	Regions []Region `json:"regions"`
}
