package ocean

import "fmt"

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
	Account *Account `json:"account"`
}

//Get account information for the current user
func (c *Client) GetAccount() (*Account, error) {
	dec, err := c.doGet("account")

	if err != nil {
		return nil, fmt.Errorf("Error retreiving acount info:\n\t%v", err)
	}

	a := &accountResp{}

	err = dec.Decode(a)

	if err != nil {
		return nil, fmt.Errorf("Error decoding Account response:\n\t%v", err)
	}

	return a.Account, nil
}
