package http

import (
	"context"
	"net/http"
)

type Client struct {
	*http.Client
	Domain string
}

func (c *Client) Dial(ctx context.Context, cfg ClientConfig) error {
	c.Client = &http.Client{}
	c.Domain = cfg.Domain

	return nil
}
