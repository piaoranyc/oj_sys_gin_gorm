package test

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func TestGetRedis(T *testing.T) {
	val2, err := rdb.Get(ctx, "name").Result()
	if errors.Is(err, redis.Nil) {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

}

func TestSetRedis(T *testing.T) {
	err := rdb.Set(ctx, "name", "mm", time.Second*10).Err()
	if err != nil {
		panic(err)
	}
}
func TestRedisGet(T *testing.T) {
	v, err := rdb.Get(ctx, "name").Result()
	if errors.Is(err, redis.Nil) {
		fmt.Println("key2 does not exist")
	}
	fmt.Println("key2", v)
}
