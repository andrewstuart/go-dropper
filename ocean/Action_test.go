package ocean

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAction(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "actions") {
			if r.URL.Path != "/droplets/123/actions" {
				t.Errorf("Request went to %s, not '/droplets/123/actions'", r.URL.Path)
			}
			w.WriteHeader(201)
			fmt.Fprintln(w, `{"status": "foo", "id": 123}`)
		} else {
			w.WriteHeader(200)
		}

	}))
	defer ts.Close()

	c := NewClient("foo")
	c.BaseUrl = ts.URL + "/"

	d := &Droplet{
		Id: 123,
	}

	d.Client = c

	_, err := d.Shutdown()

	if err != nil {
		t.Errorf("Error returned by shutdown: %v", err)
	}

}

func TestSnapshot(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/droplets/123/actions" {
			t.Error("Error: Wrong path")
		}

		a := make(Action)

		dec := json.NewDecoder(r.Body)

		err := dec.Decode(&a)

		if err != nil {
			t.Errorf("Error decoding req body:\n\t%v", err)
		}

		if a["type"] != "snapshot" {
			t.Errorf("Error: Wrong action type for snapshot: %s", a["type"])
		}

		fmt.Fprintln(w, `{"status": "foo", "id": 123}`)
	}))
	defer ts.Close()

	c := NewClient("foo")
	c.BaseUrl = ts.URL + "/"

	d := &Droplet{
		Id:     123,
		Client: c,
	}

	_, err := d.Snapshot()

	if err != nil {
		t.Errorf("Error creating snapshot:\n\t%v", err)
	}
}
