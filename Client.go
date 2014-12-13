package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const BASE_URL = "https://api.digitalocean.com/v2/"

//Basic abstraction over string in case token type changes
type Token string

//The client to use for all requests.
var cli *http.Client

//Start out by creating a client
func init() {
	cli = &http.Client{}
}

//API Client type
type Client struct {
	token Token
}

func NewClient(token Token) Client {
	return Client{
		token: token,
	}
}

func (c *Client) doReq(r *http.Request) *json.Decoder {
	t := time.Now()
	res, err := cli.Do(r)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(time.Now().Sub(t))

	return json.NewDecoder(res.Body)
}

func (c *Client) doPost(path string, r io.Reader) *json.Decoder {
	req, err := http.NewRequest("POST", BASE_URL+path, r)

	if err != nil {
		log.Fatal(err)
	}

	req.Header["Authorization"] = []string{fmt.Sprintf("Bearer %s", c.token)}

	return c.doReq(req)
}

func (c *Client) doGet(path string) *json.Decoder {
	req, err := http.NewRequest("GET", BASE_URL+path, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header["Authorization"] = []string{fmt.Sprintf("Bearer %s", c.token)}

	return c.doReq(req)
}

func (c *Client) GetRegions() []Region {
	dec := c.doGet("regions")

	rs := &RegionResp{}

	dec.Decode(rs)

	return rs.Regions
}

func (c *Client) GetImages() []Image {
	dec := c.doGet("images")

	is := &ImageResp{}

	dec.Decode(is)

	return is.Images
}

func (c *Client) GetSizes() []Size {
	dec := c.doGet("sizes")

	sr := &SizeResp{}

	dec.Decode(sr)

	return sr.Sizes
}

func (c *Client) CreateDroplet(Droplet) {
}
