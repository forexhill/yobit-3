package yobit

import (
	"context"
	"fmt"
	"testing"
)

func TestInfo(t *testing.T) {
	yobit := NewClient()
	ctx := context.Background()
	ret, err := yobit.Info(ctx)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(ret)
}
