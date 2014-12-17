package ocean

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRegions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expUrl := "/regions"
		if r.URL.Path != expUrl {
			t.Error("Error: Request expected at %s, went to %s", expUrl, r.URL.Path)
		}

		fmt.Fprintln(w, `{"regions": [{"slug": "foo", "sizes": ["512mb"], "name": "FooLand"},{"slug":"bar","sizes":["1gb"],"name":"BarLand"}]}`)

	}))
	defer ts.Close()

	c := NewClient("foo")
	c.BaseUrl = ts.URL + "/"

	rs, err := c.GetRegions()

	if err != nil {
		t.Errorf("Error getting regions:\n\t%v", err)
	}

	if len(rs) != 2 {
		t.Error("Error: wrong number of regions.")
	}
}
