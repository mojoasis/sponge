package xpay

import (
	"context"
	"fmt"

	"github.com/smartwalle/alipay/v3"
)

type AliPayClient struct {
	Client *alipay.Client
}

func (a *AliPayClient) Prepay(ctx context.Context, order PayOrder) (string, error) {
	p := alipay.TradePagePay{}
	p.OutTradeNo = order.OutTradeNo
	p.TotalAmount = fmt.Sprintf("%.2f", order.Amount)
	p.Subject = order.Subject
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err := a.Client.TradePagePay(p)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}
