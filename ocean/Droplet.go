package ocean

import (
	"fmt"
	"log"
)

type SSHKey string

// type Httper interface {
// 	DoReq(*http.Request) *json.Decoder
// 	DoPost(string, *io.Reader) *json.Decoder
// 	DoDelete(string) *json.Decoder
// 	DoGet(string) *json.Decoder
// }

type Droplet struct {
	Id                int        `json:"id"`
	Name              string     `json:"name"`
	Region            RegionSlug `json:"region"`
	Size              SizeSlug   `json:"size"`
	Image             ImageSlug  `json:"image"`
	Backups           bool       `json:"backups"`
	IPv6              bool       `json:"ipv6"`
	SshKeys           []SSHKey   `json:"ssh_keys,omitempty"`
	PrivateNetworking bool       `json:"private_networking"`
	UserData          string     `json:"user_data,omitempty"`
	*Client
}

type DropletResp struct {
	Droplets []Droplet `json:"droplets"`
}

func (d *Droplet) Delete() {
	url := fmt.Sprintf("droplets/%d", d.Id)
	log.Println(d.ResponseTimes)
	d.doDelete(url)
}
