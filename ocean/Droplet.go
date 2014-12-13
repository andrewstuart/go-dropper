package ocean

type SSHKey string

type Droplet struct {
	Name              string     `json:"name"`
	Region            RegionSlug `json:"region"`
	Size              SizeSlug   `json:"size"`
	Image             ImageSlug  `json:"image"`
	Backups           bool       `json:"backups"`
	IPv6              bool       `json:"ipv6"`
	SshKeys           []SSHKey   `json:"ssh_keys,omitempty"`
	PrivateNetworking bool       `json:"private_networking"`
	UserData          string     `json:"user_data,omitempty"`
}
