package gopixel

import (
	"errors"
	"io/ioutil"

	"net/http"
)

type Client struct {
	Key     string
	Retries uint
}

// Returns a client object
func NewClient(key string, retries uint) *Client {
	return &Client{Key: key, Retries: retries}
}

// Internal function to handle http GET requests
func (client *Client) get(url string) ([]byte, error) {
	for i := 0; i < int(client.Retries)+1; i++ {
		resp, err := http.Get("https://" + url)

		if err != nil {
			return nil, nil
		}

		defer resp.Body.Close()

		if string(resp.Status)[0] == 5 {
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)

		return body, err
	}
	return make([]byte, 0), errors.New("server side error")
}
