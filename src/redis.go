package main

import (
	"fmt"
	"github.com/fiorix/go-redis/redis"
	"strings"
)

func readFromRedis(addr string, env *[]string, appname string) error {
	rc := redis.New(addr)
	defer rc.CloseAll()
	keys, err := rc.HGetAll(appname)
	if err != nil {
		return err
	}
	for k, v := range keys {
		key, value := strings.ToUpper(k), v
		ev := fmt.Sprintf("%s=%s", key, value)
		*env = append(*env, ev)
	}
	*env = append(*env, "redis=true")
	return nil
}
