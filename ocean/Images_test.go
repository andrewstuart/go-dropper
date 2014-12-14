package ocean

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestImages(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"images": [{"name": "Foo", "slug": "foo", "regions": ["nyc3", "sfo1"], "public": true},{"name": "Bar Image", "slug": "bar-image", "regions": ["ams1"], "public": true}]}`)
	}))
	defer ts.Close()

	c := NewClient("foo")
	c.BaseUrl = ts.URL + "/"

	imgs, err := c.GetImages()

	if err != nil {
		t.Error(err)
	}

	if len(imgs) != 2 {
		t.Errorf("Not enough images returned: %d", len(imgs))
	}

	i := &imgs[0]

	if i.Name != "Foo" || i.Slug != "foo" || len(i.Regions) != 2 || !i.Public {
		t.Errorf("Image parsed incorrectly: %v", i)
	}
}
