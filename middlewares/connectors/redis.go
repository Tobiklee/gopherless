package connectors

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

const (
	REDIS_HOST_CONFIG = "services.redis.host"
	REDIS_PORT_CONFIG = "services.redis.port"
)

type (
	RedisConfig struct {
		Host     string
		Port     string
		Password string
		DB       int
	}

	RedisClient struct {
		Client *redis.Client
		Config RedisConfig
	}
)

func ConnectToRedis(config RedisConfig) RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       config.DB,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return RedisClient{Client: client, Config: config}
}

func (client *RedisClient) Ping() {
	_, err := client.Client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}

func (client *RedisClient) Set(key string, value interface{}) error {
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}
	fmt.Println("Redis Set ", p)
	err = client.Client.Set(context.Background(), key, p, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (client *RedisClient) Get(key string, dest *interface{}) (interface{}, error) {
	g := client.Client.Get(context.Background(), key)
	if g.Err() != nil {
		return nil, g.Err()
	}
	fmt.Println(" Get ", g.Val())
	json.Unmarshal([]byte(g.Val()), dest)

	return g.Val(), nil
	//return json.Unmarshal(p, dest)
}
