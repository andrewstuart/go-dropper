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
	"log"
	"net/http"
	"strings"
	"time"
)

const BASE_URL = "https://api.digitalocean.com/v2/"

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
	ResponseTimes []ResponseTime
}

//Get a client based on a token
func NewClient(token Token) Client {
	return Client{
		token: token,
	}
}

//Do a request
func (c *Client) doReq(r *http.Request) (*json.Decoder, error) {

	r.Header["Authorization"] = []string{fmt.Sprintf("Bearer %s", c.token)}

	t := time.Now()
	res, err := client.Do(r)

	if err != nil {
		log.Fatal(err)
	}

	if 400 <= res.StatusCode && res.StatusCode < 500 {
		return nil, errors.New(fmt.Sprintf("Client error %d", res.StatusCode))
	} else if 500 <= res.StatusCode && res.StatusCode < 600 {
		return nil, errors.New(fmt.Sprintf("Server error: %d", res.StatusCode))
	}

	c.ResponseTimes = append(c.ResponseTimes, ResponseTime{
		Time: time.Now().Sub(t),
		Path: r.URL.String(),
	})

	return json.NewDecoder(res.Body), nil
	// return json.NewDecoder(io.TeeReader(res.Body, os.Stdout))
}

func (c *Client) doDelete(path string) (*json.Decoder, error) {
	req, err := http.NewRequest("DELETE", BASE_URL+path, nil)

	if err != nil {
		log.Fatal(err)
	}

	return c.doReq(req)
}

//Do a post
func (c *Client) doPost(path string, r io.Reader) (*json.Decoder, error) {
	req, err := http.NewRequest("POST", BASE_URL+path, r)

	if err != nil {
		log.Fatal(err)
	}

	req.Header["Content-Type"] = []string{"application/json"}

	return c.doReq(req)
}

//Do a get
func (c *Client) doGet(path string) (*json.Decoder, error) {
	req, err := http.NewRequest("GET", BASE_URL+path, nil)

	if err != nil {
		log.Fatal(err)
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

	// dec2, err := c.doGet("images?type=distribution")

	// if err != nil {
	// 	return is.Images, err
	// }

	// is2 := &ImageResp{}

	// dec2.Decode(is2)

	// for _, i := range is2.Images {
	// 	is.Images = append(is.Images, i)
	// }

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
		d.Droplets[i].Client = c
	}

	return d.Droplets, nil
}

func (c *Client) CreateDroplet(d *Droplet) error {
	b, err := json.Marshal(d)

	if err != nil {
		return err
	}

	r := strings.NewReader(string(b))
	dec, err := c.doPost("droplets", r)

	if err != nil {
		log.Fatal(err)
	}

	dec.Decode(d)

	d.Client = c

	return nil
}
