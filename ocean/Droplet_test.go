package ocean

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteDroplet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Error("Method should be delete")
		}

		w.WriteHeader(204)
		w.Write([]byte{})
	}))
	defer ts.Close()

	c := NewClient("foo")
	c.BaseUrl = ts.URL + "/"

	d := Droplet{}
	d.Client = *c

	err := d.Delete()

	if err != nil {
		t.Errorf("Delete returned an error: %v", err)
	}
}

func TestDropErr(t *testing.T) {
	t.Log("Testing droplet errors")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(401)
	}))
	defer ts.Close()

	c := NewClient("foo")
	c.doer = fakeDoer{}

	d := Droplet{
		Client: *c,
	}

	err := d.Delete()

	if err == nil {
		t.Error("Error was not thrown")
	}
}
