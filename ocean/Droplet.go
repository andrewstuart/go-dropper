package ocean

import (
	"encoding/json"
	"strings"
)

type Slug string

// type Httper interface {
// 	DoReq(*http.Request) *json.Decoder
// 	DoPost(string, *io.Reader) *json.Decoder
// 	DoDelete(string) *json.Decoder
// 	DoGet(string) *json.Decoder
// }

type Network struct {
	IP      string `json:"ip_address"`
	Netmask string `json:"netmask"`
	Gateway string `json:"gateway"`
	Type    string `json:"type"`
}

type Droplet struct {
	Id                int                  `json:"id,omitempty"`
	Name              string               `json:"name"`
	Region            Slug                 `json:"region"`
	Size              Slug                 `json:"size"`
	Image             Slug                 `json:"image"`
	Backups           bool                 `json:"backups"`
	IPv6              bool                 `json:"ipv6,omitempty"`
	SshKeys           []Slug               `json:"ssh_keys,omitempty"`
	PrivateNetworking bool                 `json:"private_networking"`
	UserData          string               `json:"user_data,omitempty"`
	Locked            bool                 `json:"locked,omitempty"`
	Networks          map[string][]Network `json:"networks,omitempty"`
	Status            string               `json:"status,omitempty"`
	Client
}

type dropletResp struct {
	Droplets []Droplet `json:"droplets"`
	Droplet  *Droplet  `json:"droplet"`
}

//GetDroplets gets a list of the user's droplets
//Returns an error if anything went wrong.
func (c *Client) GetDroplets() ([]Droplet, error) {
	dec, err := c.doGet("droplets")

	if err != nil {
		return []Droplet{}, err
	}

	d := &dropletResp{}
	dec.Decode(d)

	for i := range d.Droplets {
		dr := &d.Droplets[i]
		dr.Client = *c
	}

	return d.Droplets, nil
}

//Pass in an adequately configured droplet type (minimum of 'Name', 'Image'(slug),
//'Size'(slug), and 'Region'(slug) must be populated
func (c *Client) CreateDroplet(d *Droplet) error {
	b, err := json.Marshal(d)

	if err != nil {
		return err
	}

	r := strings.NewReader(string(b))
	dec, err := c.doPost("droplets", r)

	if err != nil {
		return err
	}

	dr := &dropletResp{}

	dec.Decode(dr)

	*d = *dr.Droplet

	d.Client = *c

	return nil
}
