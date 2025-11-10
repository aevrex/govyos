// soren/options.go
package soren

import (
	"crypto/tls"
	"net/http"
)

type Option func(*Client)

func WithInsecureTLS() Option {
	return func(c *Client) {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}

		c.httpClient.Transport = tr
	}
}
