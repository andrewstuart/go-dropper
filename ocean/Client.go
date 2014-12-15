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
