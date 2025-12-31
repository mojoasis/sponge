package xpay

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws" // 这里使用AWS的是因为在Go中不能直接对字面量取地址
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
)

type WxPayClient struct {
	Service *native.NativeApiService
	AppID   string
	MchID   string
}

func (w *WxPayClient) Prepay(ctx context.Context, order PayOrder) (string, error) {
	_, _, err := w.Service.Prepay(ctx, native.PrepayRequest{
		Appid:       aws.String(w.AppID),
		Mchid:       aws.String(w.MchID),
		Description: aws.String(order.Subject),
		OutTradeNo:  aws.String(order.OutTradeNo),
		Amount: &native.Amount{
			Total: aws.Int64(int64(order.Amount * 100)), // 微信以分为单位
		},
	})
	if err != nil {
		return "", err
	}
	return w.AppID, nil
}
