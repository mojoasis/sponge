package xredis

import (
	"context"
	"errors"
	"time"

	"github.com/bytedance/sonic"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

// SetObject 自动将对象序列化为 JSON 存入 Redis
func (r *RedisClient) SetObject(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	// 使用 sonic 序列化
	data, err := sonic.Marshal(value)
	if err != nil {
		return err
	}
	return r.Client.Set(ctx, key, data, expiration).Err()
}

// GetObject 自动将 Redis 中的 JSON 反序列化到指定对象
func (r *RedisClient) GetObject(ctx context.Context, key string, obj interface{}) error {
	data, err := r.Client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	// 使用 sonic 反序列化
	return sonic.Unmarshal(data, obj)
}

// GetWithCache 通用缓存逻辑：先查缓存，没有则查 DB 并回写
func (r *RedisClient) GetWithCache(ctx context.Context, key string, obj interface{}, expiration time.Duration, dbQuery func() (interface{}, error)) error {
	err := r.GetObject(ctx, key, obj)
	if err == nil {
		return nil // 命中缓存
	}

	if errors.Is(err, redis.Nil) {
		// 缓存未命中，执行 DB 查询
		res, dbErr := dbQuery()
		if dbErr != nil {
			return dbErr
		}
		// 回写缓存
		_ = r.SetObject(ctx, key, res, expiration)

		// 将结果赋值给 obj (这部分通常需要反射或在 dbQuery 中处理)
		// 简单处理：重新 Marshal/Unmarshal 一次以填充指针
		b, _ := sonic.Marshal(res)
		return sonic.Unmarshal(b, obj)
	}

	return err
}
