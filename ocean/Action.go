package ocean

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type ActionResp struct {
	Id           int        `json:"id"`
	Status       string     `json:"status"`
	Type         string     `json:"type"`
	StartedAt    time.Time  `json:"started_at"`
	CompletedAt  time.Time  `json:"completed_at"`
	ResourceId   int        `json:"resource_id"`
	ResourceType int        `json:"resource_type"`
	Region       RegionSlug `json:"region"`
}

func (d *Droplet) Perform(a *Action) (*ActionResp, error) {
	if d == nil {
		return nil, errors.New("Cannot perform action on nil Droplet pointer")
	}

	if d.Id == 0 {
		return nil, errors.New("Cannot perform an action on a Droplet with no ID")
	}

	b, err := json.Marshal(a)

	if err != nil {
		return nil, err
	}

	r := strings.NewReader(string(b))

	url := fmt.Sprintf("droplets/%d/actions", d.Id)

	dec, err := d.doPost(url, r)

	ar := &ActionResp{}

	err = dec.Decode(ar)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error decoding: %v", err))
	}

	return ar, nil
}

type Action map[string]string

func NewAction(t string) *Action {
	a := make(Action)
	a["type"] = t

	return &a
}

//Delete a droplet
func (d *Droplet) Delete() error {
	url := fmt.Sprintf("droplets/%d", d.Id)
	_, err := d.doDelete(url)

	if err != nil {
		return err
	}

	return nil
}

//Reboot a Droplet
func (d *Droplet) Reboot() (*ActionResp, error) {
	return d.Perform(NewAction("reboot"))
}

//Shutdown a Droplet
func (d *Droplet) Shutdown() (*ActionResp, error) {
	return d.Perform(NewAction("shutdown"))
}

//Force off a Droplet
func (d *Droplet) PowerOff() (*ActionResp, error) {
	return d.Perform(NewAction("power_off"))
}

//Boot a Droplet
func (d *Droplet) Boot() (*ActionResp, error) {
	return d.Perform(NewAction("power_on"))
}

//Rename a Droplet
func (d *Droplet) Rename(name string) (*ActionResp, error) {
	a := *NewAction("rename")
	a["name"] = name

	return d.Perform(&a)
}
