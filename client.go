package yobit

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const endpoint = "https://yobit.net/"

type Client struct {
	BaseURL   *url.URL
	apiKey    string
	apiSecret string
}

type Response struct {
	Response *http.Response
	Body     []byte
}

func NewClient() *Client {
	baseURL, _ := url.Parse(endpoint)

	return &Client{BaseURL: baseURL}
}

func (c *Client) Auth(key, secret string) *Client {
	c.apiKey = key
	c.apiSecret = secret

	return c
}

func (c *Client) newAuthenticatedRequest(ctx context.Context, method, refURL string, params url.Values) (*http.Request, error) {
	rel, err := url.Parse(refURL)
	if err != nil {
		return nil, err
	}

	if params != nil {
		rel.RawQuery = params.Encode()
	}

	var req *http.Request
	u := c.BaseURL.ResolveReference(rel)
	param := strings.NewReader(params.Encode())

	req, err = http.NewRequest(method, u.String(), param)
	if err != nil {
		return nil, err
	}

	sign := generateSig(params.Encode(), c.apiSecret)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Key", c.apiKey)
	req.Header.Add("Sign", sign)

	req = req.WithContext(ctx)

	return req, nil
}

func (c *Client) newRequest(ctx context.Context, method, refURL string, params url.Values) (*http.Request, error) {
	rel, err := url.Parse(refURL)
	if err != nil {
		return nil, err
	}

	if params != nil {
		rel.RawQuery = params.Encode()
	}

	var req *http.Request
	u := c.BaseURL.ResolveReference(rel)

	req, err = http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := newResponse(resp)

	if v == nil {
		return response, nil
	}

	err = json.Unmarshal(response.Body, &v)
	if err != nil {
		return response, err
	}

	return response, nil
}

func generateSig(body, secret string) string {
	hasher := hmac.New(sha512.New, []byte(secret))
	hasher.Write([]byte(body))

	return hex.EncodeToString(hasher.Sum(nil))
}

func newResponse(r *http.Response) *Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		body = []byte(`Error reading body:` + err.Error())
	}

	return &Response{r, body}
}
