package main

import (
	"fmt"

	redis "github.com/Tobiklee/gopherless/middlewares/connectors"

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
	if err = client.Get("name", &d); err != nil {
		panic(err)
	}

	var dod *datum = d.(*datum)
	fmt.Println("Deconstruction:", dod.Host, "", dod.Port)

}
