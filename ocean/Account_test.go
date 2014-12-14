package ocean

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAccounts(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"account": {"droplet_limit": 12, "uuid": "abc123", "email": "admin@example.com", "email_verified": true}}`)
	}))
	defer ts.Close()

	c := NewClient("foo")
	c.BaseUrl = ts.URL + "/"

	a, err := c.GetAccount()

	if err != nil {
		t.Errorf("Error getting accounts:\n\t%v", err)
	}

	if a.UUID != "abc123" || a.Email != "admin@example.com" || !a.EmailVerified || a.DropletLimit != 12 {
		t.Errorf("Parsing failed on %v", a)
	}

}
