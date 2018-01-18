# Yobit v3

Unofficial Yobit exchange platform wrapper for Golang.

## Usage

> Info

```
package main

import (
	"context"
	"fmt"
)

func main() {
	yobit := yobit.NewClient()
	ctx := context.Background()
	info, err := yobit.Info(ctx)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(info)
}
```

> getInfo (needs authentication)

```
package main

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

func main() {
	now := time.Now().Unix()
	params := url.Values{}

	params.Set("method", "getInfo")
	params.Add("nonce", fmt.Sprintf("%v", now))

	yobit := yobit.NewClient().Auth("YOUR API KEY", "YOUR API SECRET")
	ctx := context.Background()
	getInfo, err := yobit.GetInfo(ctx, params)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(getInfo)
}
```

### Wrapped APIs

- [x]info
- [ ]ticker
- [ ]depth
- [ ]trades
- [x]getinfo
- [ ]Trade
- [ ]ActiveOrders
- [ ]OrderInfo
- [ ]CancelOrder
- [ ]TradeHistory
- [ ]GetDepositAddress
- [ ]WithdrawCoinsToAddress
- [ ]CreateYobicode
- [ ]RedeemYobicode

#### Donations

>BTC
>
>1R4Sp6GenViQQpeakD7yeThLGpdxmCKUE

>XMR
>
>42acWg6Fvjte73rG3CXFhVarjEjsKbd1tPAkPvt3nhkh5QW1Gv1Bixy9GS2eZteTFXCrjWXeAs8YdgYBXWAB7xsbUAx267u