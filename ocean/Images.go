package ocean

type Image struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Distro  string `json:"distribution"`
	Slug    Slug   `json:"slug"`
	Public  bool   `json:"public"`
	Regions []Slug `json:"regions"`
}

type ImageResp struct {
	Images []Image `json:"images"`
}
