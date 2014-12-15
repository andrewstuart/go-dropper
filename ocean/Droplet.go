package ocean

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
	*Client
}

type DropCreateResp struct {
	Droplet Droplet `json:"droplet"`
}

type DropletResp struct {
	Droplets []Droplet `json:"droplets"`
}
