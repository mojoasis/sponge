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

// Set 存入原始值 如果 value 是 string, int, bool 等基础类型，go-redis 会直接处理
func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(ctx, key, value, expiration).Err()
}

// Get 获取原始字符串
func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

// SetObject 自动处理 JSON 序列化
func (r *RedisClient) SetObject(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	var data interface{}

	switch v := value.(type) {
	case string, []byte, int, int64, float64, bool:
		data = v
	default:
		// 只有复杂结构体才走 JSON 序列化
		b, err := sonic.Marshal(value)
		if err != nil {
			return err
		}
		data = b
	}
	return r.Client.Set(ctx, key, data, expiration).Err()
}

// GetObject 增强版：泛型支持，直接返回目标类型
func (r *RedisClient) GetObject(ctx context.Context, key string, obj interface{}) error {
	data, err := r.Client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return sonic.Unmarshal(data, obj)
}

// GetCache 泛型，T 代表你要获取的数据类型，例如 *model.User
func GetCache[T any](ctx context.Context, r *RedisClient, key string, expiration time.Duration, dbQuery func() (T, error)) (T, error) {
	var obj T

	// 1. 尝试从缓存获取
	err := r.GetObject(ctx, key, &obj)
	if err == nil {
		return obj, nil // 命中缓存
	}

	// 2. 缓存未命中（处理 redis.Nil）
	if errors.Is(err, redis.Nil) {
		// 3. 执行 DB 查询
		res, dbErr := dbQuery()
		if dbErr != nil {
			return obj, dbErr
		}

		// 4. 异步回写缓存（不阻塞主流程，提升响应速度）
		// 注意：如果对一致性要求极高，请改为同步
		go func() {
			_ = r.SetObject(context.Background(), key, res, expiration)
		}()

		return res, nil
	}

	return obj, err
}
