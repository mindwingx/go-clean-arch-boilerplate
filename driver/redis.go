package driver

import (
	"context"
	"encoding/json"
	"fmt"
	redisCore "github.com/go-redis/redis/v8"
	"github.com/mindwingx/go-clean-arch-boilerplate/helper"
	"time"
)

type (
	redis struct {
		config RedisConfig
		locale LocaleAbstraction
		redis  *redisCore.Client
	}

	RedisConfig struct {
		DB       int
		Host     string
		Port     string
		Password string
		Timeout  time.Duration
	}
)

type CacheAbstraction interface {
	InitCache()
	Store(key string, data interface{}, dur time.Duration) error
	Exists(key string) bool
	Get(key string) ([]byte, error)
	Delete(key string) error
}

func NewRedis(registry RegistryAbstraction, locale LocaleAbstraction) CacheAbstraction {
	cache := new(redis)
	registry.Parse(&cache.config)
	cache.locale = locale

	cache.redis = redisCore.NewClient(&redisCore.Options{
		Addr:     fmt.Sprintf("%s:%s", cache.config.Host, cache.config.Port),
		Password: cache.config.Password,
		DB:       cache.config.DB,
	})

	return cache
}

func (r *redis) InitCache() {
	ctx, cancel := context.WithTimeout(context.Background(), r.config.Timeout)
	// call the cancel variable of the above context
	defer cancel()

	_, errConnect := r.redis.Ping(ctx).Result()
	if errConnect != nil {
		helper.CustomPanic(r.locale.Get("cache_conn_failure"), errConnect)
	}
}

func (r *redis) Store(key string, data interface{}, duration time.Duration) (err error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err = r.redis.Set(context.Background(), key, bytes, duration).Err(); err != nil {
		return err
	}

	return nil
}

func (r *redis) Exists(key string) bool {
	return r.redis.Exists(context.Background(), key).Val() == 1
}

func (r *redis) Get(key string) (b []byte, err error) {
	b, err = r.redis.Get(context.Background(), key).Bytes()
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (r *redis) Delete(key string) error {
	return r.redis.Del(context.Background(), key).Err()
}
