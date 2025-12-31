package xpay

import (
	"context"
)

// PayOrder 统一入参
type PayOrder struct {
	OutTradeNo string  // 商户订单号
	Amount     float64 // 金额
	Subject    string  // 标题
}

// Payer 统一支付接口
type Payer interface {
	Prepay(ctx context.Context, order PayOrder) (string, error) // 预支付，返回跳转链接或参数
	VerifyNotify(ctx context.Context, data interface{}) error   // 回调验签
}
