package api

import (
	"errors"

	"github.com/mattrx/poc-package-requirements/internal/requirements"
)

func init() {
	requirements.Register(func() error {
		if !configured {
			return errors.New("Api client not configured")
		}

		return nil
	})
}

type Client struct{}

var configured = false

func New(user string, pass string) (*Client, error) {

	// do some configuration and maybe return an error

	configured = true

	return &Client{}, nil
}

type Request struct{}

type Response struct{}

func (b *Client) Do(req Request) (*Response, error) {
	return &Response{}, nil
}
