package gopixel

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/valyala/fasthttp"
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
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)
		req.SetRequestURI("https://" + url)

		req.Header.Set("Accept-Encoding", "gzip")

		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		err := fasthttp.Do(req, resp)
		if err != nil {
			return make([]byte, 0), err
		}

		if fmt.Sprintf("%v", resp.StatusCode())[0] == '5' {
			continue
		}
		contentEncoding := resp.Header.Peek("Content-Encoding")
		var body []byte
		if bytes.EqualFold(contentEncoding, []byte("gzip")) {
			body, err = resp.BodyGunzip()
		} else {
			body = resp.Body()
		}

		return body, err
	}
	return make([]byte, 0), errors.New("server side error")
}
