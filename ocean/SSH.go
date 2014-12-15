package ocean

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type SSHKey struct {
	Id          int    `json:"id,omitempty"`
	Fingerprint string `json:"fingerprint, omitempty"`
	PublicKey   string `json:"public_key"`
	Name        string `json:"name"`
}

func ReadSSHKey(path, name string) (*SSHKey, error) {
	b, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	k := &SSHKey{
		PublicKey: string(b),
		Name:      name,
	}

	return k, nil
}

func (c *Client) CreateSSHKey(s *SSHKey) (*SSHKey, error) {
	b, err := json.Marshal(s)

	if err != nil {
		return nil, err
	}

	rdr := strings.NewReader(string(b))

	dec, err := c.doPost("account/keys", rdr)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error sending ssh key to DO:\n\t%v", err))
	}

	k := &SSHKey{}

	dec.Decode(k)

	return k, nil
}
