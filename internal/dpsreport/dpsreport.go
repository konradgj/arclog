package dpsreport

import (
	"net/http"
	"time"
)

const baseUrl = "https://dps.report/"

type Client struct {
	Client http.Client
}

func (c *Client) NewClient(timeout time.Duration) {
	c.Client = http.Client{
		Timeout: timeout,
	}
}
