package ocean

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type fakeDoer struct{}

func (d fakeDoer) doReq(r *http.Request) (*json.Decoder, error) {
	return nil, errors.New("Foo")
}

func TestHeaders(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Test auth header
		if r.Header["Authorization"][0] != "Bearer abc" {
			t.Error("Authorization header did not equal Bearer abc")
		}
	}))

	testCreate(ts)
}

func TestCreateDroplet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"id": 123, "image": "foo"}`)

		dec := json.NewDecoder(r.Body)

		d := &Droplet{}

		dec.Decode(d)

		if d.Id != 0 {
			t.Error("Sent create droplet with an id.")
		}
	}))
	defer ts.Close()
}

func TestErrors(t *testing.T) {
	is := []int{404, 500, 400, 499, 599, 501, 503}

	for _, i := range is {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(i)

			//Test auth header
			if r.Header["Authorization"][0] != "Bearer abc" {
				t.Error("No Authorization header present")
			}
		}))

		err := testCreate(ts)

		if err == nil {
			t.Error("Should have passed back the client error as an error")
		}
	}
}

func TestBaseDefault(t *testing.T) {
	c := NewClient("abc")

	if c.BaseUrl != "https://api.digitalocean.com/v2/" {
		t.Errorf("Base URL was not set as expected: %s", c.BaseUrl)
	}
}

func TestErr(t *testing.T) {
	c := NewClient("abc")
	c.doer = fakeDoer{}

	_, err := c.GetDroplets()

	if err == nil {
		t.Error("Did not pas error back")
		log.Printf("err %+v\n", err)
	}
}

func TestGetDrops(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"droplets": [{"id": 123, "name": "foo", "image": "ruby", "size": "512mb"},{"id": 123, "name": "foo", "image": "ruby", "size": "512mb"}]}`)
	}))
	defer ts.Close()

	c := NewClient("abc")
	c.BaseUrl = ts.URL + "/"
	drops, err := c.GetDroplets()

	if err != nil {
		t.Errorf("Droplets parsing failed:\t%v", err)
	}

	l := 2
	if len(drops) != l {
		t.Errorf("Droplet length should have been %d", l)
	}

	if drops[0].Size != "512mb" {
		t.Errorf("Size was not 512mb but was %s instead", drops[0].Size)
	}
}

func TestRegion(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"regions": [{"name": "123", "slug": "thaBest", "sizes": ["12", "23", "34"]}]}`)
	}))
	defer ts.Close()

	c := setup(ts)

	rs, err := c.GetRegions()

	if err != nil {
		t.Error("Get Regions error")
	}

	if len(rs) != 1 {
		t.Errorf("Wrong number of regions (%d)", len(rs))
	} else {
		if rs[0].Slug != "thaBest" {
			t.Errorf("Didn't parse slug")
		}
	}
}

func TestSizeError(t *testing.T) {
	c := NewClient("abc")
	c.doer = fakeDoer{}

	_, err := c.GetSizes()

	if err == nil {
		t.Error("GetSizes did not properly return error")
	}
}

func TestSize(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"sizes": [{"slug": "512mb", "vcpus": 1, "disk": 20000, "price_monthly": 1.2, "price_hourly": 0.12, "regions": ["sfo1"], "memory": 1000, "transfer": 1234}]}`)
	}))
	defer ts.Close()

	c := NewClient("abc")
	c.BaseUrl = ts.URL + "/"

	ss, err := c.GetSizes()

	log.Println(err)

	if err != nil {
		t.Error("GetSizes returned an error")
	}

	if len(ss) != 1 {
		t.Errorf("Wrong number of sizes returned")
	}

	s := &ss[0]
	if s.Slug != "512mb" || s.Disk != 20000 || s.VCpus != 1 || len(s.Regions) != 1 || s.PriceHourly != 0.12 || s.PriceMonthly != 1.2 || s.Transfer != 1234 || s.Memory != 1000 {
		t.Errorf("Improperly parsed size: %v", s)
	}
}

func setup(ts *httptest.Server) *Client {
	c := NewClient("abc")
	c.BaseUrl = ts.URL + "/"
	return c
}

func testCreate(ts *httptest.Server) error {
	defer ts.Close()

	c := NewClient("abc")
	c.BaseUrl = ts.URL + "/"
	err := c.CreateDroplet(&Droplet{})

	return err
}
