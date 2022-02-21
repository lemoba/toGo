package main

import (
	"context"
	"fmt"

	"github.com/lemoba/toGo/17-redis/cache"
)

var ctx = context.Background()

func main() {
	err := cache.RedisClient.Set(ctx, "name", "ranen", 0)

	if err != nil {
		fmt.Printf("set error: %v\n", err)
	}
}
