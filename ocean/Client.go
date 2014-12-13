package ocean

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
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

//Used for measuring client response times
type ResponseTime struct {
	Time time.Duration
	Path string
}

//API Client type
type Client struct {
	token         Token
	responseTimes []ResponseTime
}

//Get a client based on a token
func NewClient(token Token) Client {
	return Client{
		token: token,
	}
}

//Do a request
func (c *Client) doReq(r *http.Request) *json.Decoder {
	t := time.Now()
	res, err := cli.Do(r)

	if err != nil {
		log.Fatal(err)
	}

	c.responseTimes = append(c.responseTimes, ResponseTime{
		Time: time.Now().Sub(t),
		Path: r.RequestURI,
	})

	return json.NewDecoder(io.TeeReader(res.Body, os.Stdout))
}

//Do a post
func (c *Client) doPost(path string, r io.Reader) *json.Decoder {
	req, err := http.NewRequest("POST", BASE_URL+path, r)

	if err != nil {
		log.Fatal(err)
	}

	req.Header["Authorization"] = []string{fmt.Sprintf("Bearer %s", c.token)}

	return c.doReq(req)
}

//Do a get
func (c *Client) doGet(path string) *json.Decoder {
	req, err := http.NewRequest("GET", BASE_URL+path, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header["Authorization"] = []string{fmt.Sprintf("Bearer %s", c.token)}

	return c.doReq(req)
}

//Get a list of regions for the user
func (c *Client) GetRegions() []Region {
	dec := c.doGet("regions")

	rs := &RegionResp{}

	dec.Decode(rs)

	return rs.Regions
}

//Get a list of images for the user
func (c *Client) GetImages() []Image {
	dec := c.doGet("images")

	is := &ImageResp{}

	dec.Decode(is)

	dec2 := c.doGet("images?type=distribution")

	is2 := &ImageResp{}

	dec2.Decode(is2)

	for _, i := range is2.Images {
		is.Images = append(is.Images, i)
	}

	return is.Images
}

//Get a list of sizes for the user
func (c *Client) GetSizes() []Size {
	dec := c.doGet("sizes")

	sr := &SizeResp{}

	dec.Decode(sr)

	return sr.Sizes
}

func (c *Client) CreateDroplet(d *Droplet) {
	b, err := json.Marshal(d)

	if err != nil {
		log.Fatal(err)
	}

	r := strings.NewReader(string(b))
	dec := c.doPost("droplets", r)
	dec.Decode(d)
}
