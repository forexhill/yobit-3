package yobit

const (
	key    = "KEY HERE"
	secret = "SECRET HERE"
)

func newAuthClient() *Client {
	return NewClient().Auth(key, secret)
}
