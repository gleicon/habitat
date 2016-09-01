package main

import (
	"strings"

	"github.com/fiorix/go-redis/redis"
)

func readFromRedis(addr string, env *[]string, appname string) error {
	rc := redis.New(addr)
	defer rc.Close()
	keys, err := rc.HGetAll(appname)
	if err != nil {
		return err
	}
	for k, v := range keys {
		key, value := strings.ToUpper(k), v
		addToEnv(env, key, value)
	}
	addToEnv(env, "redis", "true")
	return nil
}
