package global

import (
	"sponge/internal/conf"
	"sponge/pkg/xoss"
	"sponge/pkg/xpay"
	"sponge/pkg/xredis"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config *conf.AppConfig     // 全局配置句柄
	Logger *zap.Logger         // 添加这一行
	DB     *gorm.DB            // 全局 MySQL 实例
	Redis  *xredis.RedisClient // 全ify后的 Redis 实例
	OSS    *xoss.OSSClient     // 全局 OSS 实例
	// AliPay WxPay  支付
	AliPay xpay.Payer
	WxPay  xpay.Payer
)
