package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
	_ "github.com/go-redis/redis"

	_ "github.com/go-sql-driver/mysql"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func initClient(ctx context.Context) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "root", // no password set
		DB:       0,      // use default DB
	})

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
func main() {
	//测试rset
	//fix fatal
	//connect redis
	ctx := context.Background()
	err := initClient(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("connect redis successfly")

}
