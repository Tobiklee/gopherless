package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type (
	Author struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
)

func main() {
	var err error
	fmt.Println("Go Redis tutorial")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	json, err := json.Marshal(Author{
		Name: "Elliot",
		Age:  25})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set("id1234", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set("name", "James", 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	//getFromRedis(client, "name")
	//getFromRedis(client, "id1234")
	getFromRedis(client, "sam")
	getFromRedis(client, "able")

}

func getFromRedis(client *redis.Client, id string) {
	val, err := client.Get(id).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}
