package ocean

type Account struct {
	DropletLimit  int    `json:"droplet_limit"`
	Email         string `json:"email"`
	UUID          string `json:"uuid"`
	EmailVerified bool   `json:"email_verified"`
}

type AccountResp struct {
	Account Account `json:"account"`
}
