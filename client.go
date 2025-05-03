package ionq

import (
	"fmt"
	"net/http"
)

type Client struct {
	endpoint string // example: https://api.ionq.co/v0.3
	apiKey   string
	client   *http.Client
}

func NewClient(endpoint string, apiKey string) *Client {
	return &Client{
		endpoint: endpoint,
		apiKey:   apiKey,
		client:   &http.Client{},
	}
}

func (c *Client) makeURL(path string) string {
	return fmt.Sprintf("%s/%s", c.endpoint, path)
}
