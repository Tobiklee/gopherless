package main

import (
	"fmt"

	redis "github.com/Tobiklee/gopherless/middlewares/connectors"
	"github.com/Tobiklee/gopherless/middlewares/validation"

	"github.com/spf13/viper"
)

type datum struct {
	Host string
	Port int16
}

func main() {
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config/") // optionally look for config in the working directory
	err := viper.ReadInConfig()      // Find and read the config file
	if err != nil {                  // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	abc := validation.ErrorResponse{}
	fmt.Println(abc)
	fmt.Println(viper.Get(redis.REDIS_HOST_CONFIG))
	fmt.Println(viper.Get(redis.REDIS_PORT_CONFIG))
	viper.Get(redis.REDIS_HOST_CONFIG)

	client := redis.ConnectToRedis(redis.RedisConfig{
		Host: viper.GetString(redis.REDIS_HOST_CONFIG),
		Port: viper.GetString(redis.REDIS_PORT_CONFIG),
	})

	fmt.Println("after connecting to redis", client.Client)
	client.Ping()
	if err := client.Set("name", datum{Host: "localhost", Port: 22}); err != nil {
		panic(err)
	}

	var d interface{} = &datum{}
	if _, err = client.Get("name", &d); err != nil {
		panic(err)
	}

	fmt.Println("after getting from redis:", d)

	var dod *datum = d.(*datum)
	fmt.Println("deconstruction:", dod.Host, "", dod.Port)

}
