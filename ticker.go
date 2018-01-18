package yobit

import "context"

type Ticker map[string]struct {
	High    float64 `json:"high"`
	Low     float64 `json:"low"`
	Avg     float64 `json:"avg"`
	Vol     float64 `json:"vol"`
	VolCur  float64 `json:"vol_cur"`
	Last    float64 `json:"last"`
	Buy     float64 `json:"buy"`
	Sell    float64 `json:"sell"`
	Updated int     `json":updated"`
}

func (c *Client) Ticker(ctx context.Context, ticker string) (Ticker, error) {
	path := "api/3/ticker/" + ticker
	req, err := c.newRequest(ctx, "GET", path, nil)
	if err != nil {
		return Ticker{}, err
	}

	var ret = &Ticker{}

	_, err = c.do(req, ret)
	if err != nil {
		return *ret, err
	}

	return *ret, nil
}
