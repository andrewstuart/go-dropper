package ocean

type DropletSize string

type SSHKey string

type Droplet struct {
	Name              string      `json:"name"`
	Region            RegionSlug  `json:"region"`
	Size              DropletSize `json:"size"`
	Backups           bool        `json:"backups"`
	IPv6              bool        `json:"ipv6"`
	SshKeys           []SSHKey    `json:"ssh_keys,omitempty"`
	PrivateNetworking bool        `json:"private_networking"`
	UserData          string      `json:"user_data,omitempty"`
}
