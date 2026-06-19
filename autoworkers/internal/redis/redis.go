package redis

import (
	"context"
	"fmt"

	goredis "github.com/redis/go-redis/v9"
)

type Redis struct{
	Client *goredis.Client
}
func Constructor() *Redis{
	ctx := context.Background()
	r := goredis.NewClient(&goredis.Options{Addr: "localhost:6379",
										Password: "",
										DB: 0})
	b,err := r.Ping(ctx).Result()
	if err!=nil{
		fmt.Println(err)

	}else{

		fmt.Println(b)
	}
	st := &Redis{
		Client: r,
	}
	return st
	
}