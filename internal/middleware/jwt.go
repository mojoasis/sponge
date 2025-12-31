package middleware

import (
	"sponge/pkg/res"
	"strings"
	"time"

	"sponge/pkg/global"
	"sponge/pkg/utils"
	"sponge/pkg/xerror"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取 Header 中的 Token
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			res.FailCode(xerror.ERROR_AUTH_TOKEN, "未携带Token", c)
			c.Abort()
			return
		}

		// 格式通常是 "Bearer {token}"
		parts := strings.SplitN(tokenHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			res.FailCode(xerror.ERROR_AUTH_TOKEN, "Token格式错误", c)
			c.Abort()
			return
		}
		tokenString := parts[1]

		// 2. 解析 Token
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			res.FailCode(xerror.ERROR_AUTH_TOKEN, "Token无效", c)
			c.Abort()
			return
		}

		// 3. Redis 双重校验 & 自动续期
		// 只有 Redis 里存在的 Token 才是有效的（支持服务端强制登出）
		redisKey := "login:token:" + claims.UserID
		redisToken, err := global.Redis.Client.Get(c, redisKey).Result()

		if err != nil || redisToken != tokenString {
			// Redis 中不存在，或与当前 Token 不匹配（可能已在别处登录）
			res.FailCode(xerror.ERROR_AUTH_TIMEOUT, "登录已过期或失效", c)
			c.Abort()
			return
		}

		// 4. 自动续期逻辑 (Sliding Window)
		// 如果 Redis key 的剩余生存时间 (TTL) 小于 12 小时，则重置为 24 小时
		ttl := global.Redis.Client.TTL(c, redisKey).Val()
		if ttl < 12*time.Hour {
			// 仅更新 Redis 的过期时间，不重新生成 Token（前端无感）
			global.Redis.Client.Expire(c, redisKey, 24*time.Hour)
		}

		// 5. 将用户信息存入上下文，供后续 Handler 使用
		c.Set("userID", claims.UserID)
		c.Set("username", claims.UserName)
		c.Next()
	}
}
