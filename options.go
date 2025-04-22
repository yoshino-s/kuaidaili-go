package kuaidailigo

import "net/http"

type WithOption func(*BaseClient)

func WithHttpClient(client *http.Client) WithOption {
	return func(c *BaseClient) {
		c.http = client
	}
}

func WithSignType(signType SignType) WithOption {
	return func(c *BaseClient) {
		c.sign = signType
	}
}
