package ocean

//Account holds all relevant account info
type Account struct {
	DropletLimit  int    `json:"droplet_limit"`
	Email         string `json:"email"`
	UUID          string `json:"uuid"`
	EmailVerified bool   `json:"email_verified"`
	SSHKeys       []*SSHKey
	SSHByName     map[string]SSHKey
}

type accountResp struct {
	Account Account `json:"account"`
}
