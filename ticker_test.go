package yobit

import (
	"context"
	"fmt"
	"testing"
)

func TestTicker(t *testing.T) {
	ticker := "doge_usd"
	yobit := NewClient()
	ctx := context.Background()
	ret, err := yobit.Ticker(ctx, ticker)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(ret)
}
