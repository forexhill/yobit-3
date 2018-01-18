package yobit

import (
	"context"
	"fmt"
	"net/url"
	"testing"
	"time"
)

func TestGetInfo(t *testing.T) {
	now := time.Now().Unix()
	params := url.Values{}

	params.Set("method", "getInfo")
	params.Add("nonce", fmt.Sprintf("%v", now))

	yobit := newAuthClient()
	ctx := context.Background()
	ret, err := yobit.GetInfo(ctx, params)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(ret)
}
