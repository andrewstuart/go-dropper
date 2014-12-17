package ocean

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

//Type ActionResult is returned by performing an action (most methods on a Droplet or Image)
type ActionResult struct {
	Id           int       `json:"id"`
	Status       string    `json:"status"`
	Type         string    `json:"type"`
	StartedAt    time.Time `json:"started_at,omitempty"`
	CompletedAt  time.Time `json:"completed_at,omitempty"`
	ResourceId   int       `json:"resource_id"`
	ResourceType string    `json:"resource_type"`
	Region       Slug      `json:"region"`
}

type actionResp struct {
	Action  *ActionResult   `json:"action"`
	Actions []*ActionResult `json:"actions"`
}

func (c *Client) GetActionLog() ([]*ActionResult, error) {
	dec, err := c.doGet("actions?per_page=200")

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error retreiving actions:\n\t%v", err))
	}

	ar := &actionResp{}

	err = dec.Decode(ar)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error decoding response:\n\t%v", err))
	}

	return ar.Actions, nil
}

//Have the droplet perform an action.
func (d *Droplet) Perform(a *Action) (*ActionResult, error) {
	if d.Id == 0 {
		return nil, errors.New("Cannot perform an action on a Droplet with no ID")
	}

	if a == nil {
		return nil, errors.New("Action provided was nil")
	}

	b, err := json.Marshal(a)

	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(b)

	url := fmt.Sprintf("droplets/%d/actions", d.Id)

	dec, err := d.doPost(url, r)

	if err != nil {
		return nil, err
	}

	ar := &ActionResult{}

	err = dec.Decode(ar)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error decoding: %v", err))
	}

	return ar, nil
}

//An action type is required to perform an action
type Action map[string]string

//NewAction returns an action of type 't'
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
func (d *Droplet) Reboot() (*ActionResult, error) {
	return d.Perform(NewAction("reboot"))
}

//Shutdown a Droplet
func (d *Droplet) Shutdown() (*ActionResult, error) {
	return d.Perform(NewAction("shutdown"))
}

//Force off a Droplet
func (d *Droplet) PowerOff() (*ActionResult, error) {
	return d.Perform(NewAction("power_off"))
}

//Boot a Droplet
func (d *Droplet) Boot() (*ActionResult, error) {
	return d.Perform(NewAction("power_on"))
}

//Rename a Droplet
func (d *Droplet) Rename(name string) (*ActionResult, error) {
	a := *NewAction("rename")

	a["name"] = name

	return d.Perform(&a)
}

func (d *Droplet) Snapshot(name string) (*ActionResult, error) {
	a := *NewAction("snapshot")
	a["name"] = name
	return d.Perform(&a)
}

func (d *Droplet) EnableIPv6() (*ActionResult, error) {
	return d.Perform(NewAction("enable_ipv6"))
}

func (d *Droplet) Rebuild(image string) (*ActionResult, error) {
	a := *NewAction("rebuild")
	a["image"] = image
	return d.Perform(&a)
}
