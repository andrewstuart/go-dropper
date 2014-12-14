package ocean

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAction(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "actions") {
			if r.URL.Path != "droplets/123/actions" {
				t.Errorf("Request went to %s, not 'actions.'", r.URL.Path)
			}
		}
	}))
	defer ts.Close()

	c := NewClient("foo")
	c.BaseUrl = ts.URL

	d := Droplet{
		Id: 123,
	}

	c.CreateDroplet(&d)

	a, err := d.Shutdown()

	if err != nil {
		t.Errorf("Error returned by shutdown: %v", err)
	}

	t.Log(a)

}
