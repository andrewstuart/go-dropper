package ocean

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestAction(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		io.Copy(os.Stdout, r.Body)

		if strings.Contains(r.URL.Path, "actions") {
			if r.URL.Path != "droplets/123/actions" {
				t.Errorf("Request went to %s, not 'actions.'", r.URL.Path)
			}
			w.WriteHeader(201)
			fmt.Fprintln(w)
		} else {
			w.WriteHeader(200)
		}

	}))
	defer ts.Close()

	c := NewClient("foo")
	c.BaseUrl = ts.URL

	d := &Droplet{
		Id: 123,
	}

	d.Client = c

	// _, err := d.Shutdown()

	// if err != nil {
	// 	t.Errorf("Error returned by shutdown: %v", err)
	// }

}
