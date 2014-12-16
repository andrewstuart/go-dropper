package ocean

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

//Type ActionResult is returned by performing an action (most methods on a Droplet or Image)
type ActionResult struct {
	Id           int       `json:"id"`
	Status       string    `json:"status"`
	Type         string    `json:"type"`
	StartedAt    time.Time `json:"started_at"`
	CompletedAt  time.Time `json:"completed_at"`
	ResourceId   int       `json:"resource_id"`
	ResourceType int       `json:"resource_type"`
	Region       Slug      `json:"region"`
}

type ActionResp struct {
	Action  *Action   `json:"action"`
	Actions []*Action `json:"actions"`
}

func (c *Client) GetActionLog() ([]*ActionResult, error) {
	dec, err := c.doGet("actions")

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error retreiving actions:\n\t%v", err))
	}

	ar := []*ActionResult{}

	err = dec.Decode(ar)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error decoding response:\n\t%v", err))
	}

	return ar, nil
}

//Have the droplet perform an action.
func (d *Droplet) Perform(a *Action) (*ActionResult, error) {
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
