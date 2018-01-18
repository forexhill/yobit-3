package yobit

import "context"

type Info struct {
	ServerTime int `json:"server_time"`
	Pairs      map[string]struct {
		DecimalPlaces float64 `json:"decimal_places"`
		MinPrice      float64 `json:"min_price"`
		MaxPrice      float64 `json:"max_price"`
		MinPmount     float64 `json:"min_amount"`
		MinTotal      float64 `json:"min_total"`
		Hidden        int     `json:"hidden"`
		Fee           float64 `json:"fee"`
		FeeBuyer      float64 `json:"fee_buyer"`
		FeeSeller     float64 `json:"fee_seller"`
	} `json:"pairs"`
}

func (c *Client) Info(ctx context.Context) (Info, error) {
	req, err := c.newRequest(ctx, "GET", "api/3/info", nil)
	if err != nil {
		return Info{}, err
	}

	var ret = &Info{}

	_, err = c.do(req, ret)
	if err != nil {
		return *ret, err
	}

	return *ret, nil
}
