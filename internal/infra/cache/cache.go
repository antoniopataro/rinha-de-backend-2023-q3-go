package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/antoniopataro/rinha-de-backend-2023-q3-go/internal/config"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
}

func (cache *Cache) Get(ctx context.Context, key string) (string, error) {
	query := cache.client.Get(ctx, key)

	if query.Err() != nil {
		return "", query.Err()
	}

	return query.Val(), nil
}

func (cache *Cache) Set(ctx context.Context, key string, value interface{}) error {
	value, err := json.Marshal(value)

	if err != nil {
		return err
	}

	query := cache.client.Set(ctx, key, value, 0)

	if query.Err() != nil {
		return query.Err()
	}

	return nil
}

func MakeCache(envs *config.Envs) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", envs.CACHE_ADDRESS, envs.CACHE_PORT),
		DB:       0,
		Password: "",
	})

	if ping := client.Ping(context.Background()); ping.Err() != nil {
		log.Fatal(ping.Err())
	}

	return &Cache{
		client: client,
	}
}
