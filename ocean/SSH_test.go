package ocean

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateSSH(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)

		if err != nil {
			t.Errorf("Error reading request body:\n\t%v", err)
		}
		body := string(b)

		if !strings.Contains(body, "public_key") {
			t.Errorf("Request body did not contain public_key")
		}

		log.Println(body)
	}))
	defer ts.Close()

	c := &Client{}
	c.BaseUrl = ts.URL + "/"

	key := &SSHKey{
		PublicKey: "ssh-rsa foo a@example.com",
		Name:      "bar",
	}

	err := c.CreateSSHKey(key)

	if err != nil {
		t.Errorf("Error creating ssh key:\r\t%v", err)
	}
}
