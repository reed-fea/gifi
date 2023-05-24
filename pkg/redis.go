package pkg

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis-15707.c92.us-east-1-3.ec2.cloud.redislabs.com:15707",
		Password: "X07peTaMpJqvDObMXfWLtYSSsDYapquQ",
		DB:       0,
	})

	defer fmt.Println("End")

	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(pong)
}

func Get(key string) {

	// err := rdb.Set(ctx, "name", "123000", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	ctx := context.Background()

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name", val)

	// val2, err := rdb.Get(ctx, "key2").Result()
	// if err == redis.Nil {
	// 	fmt.Println("key2 does not exist")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("key2", val2)
	// }
}
