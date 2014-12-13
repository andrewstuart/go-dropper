package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func (c *Client) GetRegions() {
	req, err := http.NewRequest("GET", BASE_URL+"regions", nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header["Authorization"] = []string{fmt.Sprintf("Bearer %s", c.token)}

	res, err := cli.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(res.Body)

	rs := &RegionResp{}

	dec.Decode(rs)

	fmt.Println(rs)
}
