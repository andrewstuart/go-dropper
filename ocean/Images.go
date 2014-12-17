package ocean

import (
	"errors"
	"fmt"
)

type Image struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Distro  string `json:"distribution"`
	Slug    Slug   `json:"slug"`
	Public  bool   `json:"public"`
	Regions []Slug `json:"regions"`
}

type imageResp struct {
	Images []Image `json:"images"`
	Image  *Image  `json:"image"`
}

//Get a list of images for the user
//May return both an error and a list of images in the case that a request fails
func (c *Client) GetImages() ([]Image, error) {
	dec, err := c.doGet("images?per_page=100")

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error retrieving images:\n\t%v", err))
	}

	is := &imageResp{}

	err = dec.Decode(is)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error decoding response from DO:\n\t%v", err))
	}

	return is.Images, nil
}
