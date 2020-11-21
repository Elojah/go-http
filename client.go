package http

import (
	"context"
)

type Client struct {
}

func (c *Client) Dial(ctx context.Context, cfg ConfigClient) error {
	return nil
}

func (c *Client) Close(ctx context.Context) error {
	return nil
}
