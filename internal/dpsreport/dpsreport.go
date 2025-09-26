package dpsreport

import (
	"net/http"
	"time"
)

const baseUrl = "https://dps.report/"

type Client struct {
	Client http.Client
}

func NewClient(timeout time.Duration) *Client {
	c := http.Client{
		Timeout: timeout,
	}

	return &Client{
		Client: c,
	}
}
