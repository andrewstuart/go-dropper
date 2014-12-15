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
	Fingerprint string `json:"fingerprint,omitempty"`
	PublicKey   string `json:"public_key"`
	Name        string `json:"name"`
}

type sshResp struct {
	Key  SSHKey   `json:"ssh_key"`
	Keys []SSHKey `json:"ssh_keys"`
}

func ReadSSHKey(path, name string) (*SSHKey, error) {
	b, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	s := string(b)

	s = strings.Replace(s, "\n", "", -1)

	k := &SSHKey{
		PublicKey: s,
		Name:      name,
	}

	return k, nil
}

func (c *Client) CreateSSHKey(s *SSHKey) error {
	b, err := json.Marshal(s)

	if err != nil {
		return err
	}

	rdr := strings.NewReader(string(b))

	dec, err := c.doPost("account/keys", rdr)

	if err != nil {
		return errors.New(fmt.Sprintf("Error sending ssh key to DO:\n\t%v", err))
	}

	resp := &sshResp{}

	dec.Decode(resp)

	*s = resp.Key

	return nil
}

func (c *Client) GetSSHKeys() ([]SSHKey, error) {
	dec, err := c.doGet("account/keys")

	if err != nil {
		return []SSHKey{}, errors.New(fmt.Sprintf("Error getting keys:\n\t%v", err))
	}

	sr := &sshResp{}

	err = dec.Decode(sr)

	if err != nil {
		return []SSHKey{}, err
	}

	return sr.Keys, nil
}
