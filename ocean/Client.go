//Package ocean encapsulates the DigitalOcean API into an easy-to-use, idiomatic
//golang package.
//
//It's designed to easily be able to inspect, create, update, and destroy droplets
//with relative ease, with the implementation of digitalocean tooling in mind,
//eventually expanding to VPS-provider-agnostic tooling with pluggable interfaces
package ocean

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type doer interface {
	doReq(*http.Request) (*json.Decoder, error)
}

const DEFAULT_BASE_URL = "https://api.digitalocean.com/v2/"

//Basic abstraction over string in case token type changes
type Token string

//The client to use for all requests.
var client *http.Client

//Start out by creating a client
func init() {
	client = &http.Client{}
}

//Used for measuring client response times
type ResponseTime struct {
	Time time.Duration
	Path string
}

//API Client type
type Client struct {
	token         Token
	BaseUrl       string
	ResponseTimes []ResponseTime
	Account       *Account
	// WhenReady     <-chan bool
	doer
}

//Get a client based on a token
func NewClient(token Token) *Client {
	c := &Client{
		token:   token,
		BaseUrl: DEFAULT_BASE_URL,
		// WhenReady: make(chan bool),
	}

	return c
}

//Do a request
func (c *Client) doReq(r *http.Request) (*json.Decoder, error) {

	r.Header["Authorization"] = []string{fmt.Sprintf("Bearer %s", c.token)}

	t := time.Now()
	res, err := client.Do(r)

	if err != nil {
		return nil, err
	}

	if 400 <= res.StatusCode && res.StatusCode < 500 {
		return nil, errors.New(fmt.Sprintf("Client error: %s", res.Status))
	} else if 500 <= res.StatusCode && res.StatusCode < 600 {
		return nil, errors.New(fmt.Sprintf("Server error: %s", res.Status))
	}

	c.ResponseTimes = append(c.ResponseTimes, ResponseTime{
		Time: time.Now().Sub(t),
		Path: r.URL.String(),
	})

	return json.NewDecoder(res.Body), nil
	// return json.NewDecoder(io.TeeReader(res.Body, os.Stdout)), nil
}

func (c *Client) doDelete(path string) (*json.Decoder, error) {
	req, err := http.NewRequest("DELETE", c.BaseUrl+path, nil)

	if err != nil {
		return nil, err
	}

	return c.doReq(req)
}

//Do a post
func (c *Client) doPost(path string, r io.Reader) (*json.Decoder, error) {
	// req, err := http.NewRequest("POST", c.BaseUrl+path, io.TeeReader(r, os.Stdout))
	req, err := http.NewRequest("POST", c.BaseUrl+path, r)

	if err != nil {
		return nil, err
	}

	req.Header["Content-Type"] = []string{"application/json"}

	return c.doReq(req)
}

//Do a get
func (c *Client) doGet(path string) (*json.Decoder, error) {
	req, err := http.NewRequest("GET", c.BaseUrl+path, nil)

	if err != nil {
		return nil, err
	}

	return c.doReq(req)
}

//Get a list of regions for the user
func (c *Client) GetRegions() ([]Region, error) {
	dec, err := c.doGet("regions")

	if err != nil {
		return []Region{}, err
	}

	rs := &RegionResp{}

	dec.Decode(rs)

	return rs.Regions, nil
}

//Get a list of images for the user
//May return both an error and a list of images in the case that a request fails
func (c *Client) GetImages() ([]Image, error) {
	dec, err := c.doGet("images?per_page=100")

	if err != nil {
		return []Image{}, err
	}

	is := &ImageResp{}

	dec.Decode(is)

	return is.Images, nil
}

//Get a list of sizes for the user
func (c *Client) GetSizes() ([]Size, error) {
	dec, err := c.doGet("sizes")

	if err != nil {
		return []Size{}, err
	}

	sr := &SizeResp{}

	dec.Decode(sr)

	return sr.Sizes, nil
}

//GetDroplets gets a list of the user's droplets
//Returns an error if anything went wrong.
func (c *Client) GetDroplets() ([]Droplet, error) {
	dec, err := c.doGet("droplets")

	if err != nil {
		return []Droplet{}, err
	}

	d := &DropletResp{}
	dec.Decode(d)

	for i := range d.Droplets {
		dr := &d.Droplets[i]
		dr.Client = c
	}

	return d.Droplets, nil
}

type dropletCreateResp struct {
	Droplet *Droplet `json:"droplet"`
}

//Pass in an adequately configured droplet type (minimum of 'Name', 'Image'(slug),
//'Size'(slug), and 'Region'(slug) must be populated
func (c *Client) CreateDroplet(d *Droplet) error {
	b, err := json.Marshal(d)

	if err != nil {
		return err
	}

	r := strings.NewReader(string(b))
	dec, err := c.doPost("droplets", r)

	if err != nil {
		return err
	}

	dr := &dropletCreateResp{}

	dec.Decode(dr)

	*d = *dr.Droplet

	d.Client = c

	return nil
}

type AccountResp struct {
	Account *Account `json:"account"`
}

//Get account information for the current user
func (c *Client) GetAccount() (*Account, error) {
	dec, err := c.doGet("account")

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error retreiving acount info:\n\t%v", err))
	}

	a := &AccountResp{}

	dec.Decode(a)

	return a.Account, nil
}
