package connectors

import (
	"fmt"

	"github.com/go-redis/redis"
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

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return RedisClient{Config: config}
}

func (client *RedisClient) Ping() {
	pong, err := client.Client.Ping().Result()
	fmt.Println(pong, err)
}

func (client *RedisClient) Set() {
	err := client.Client.Set("name", "James", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func (client *RedisClient) Get(id string) {
	val, err := client.Client.Get(id).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}
