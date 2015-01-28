package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"strings"
)

func readFromConsul(addr string, env *[]string, appname string) error {
	df := api.DefaultConfig()
	df.Address = addr
	client, _ := api.NewClient(df)
	kv := client.KV()
	keys, _, err := kv.List(appname, nil)
	if err != nil {
		return err
	}

	for _, kp := range keys {
		kk := strings.Split(kp.Key, "/")
		key, value := strings.ToUpper(kk[len(kk)-1]), kp.Value
		ev := fmt.Sprintf("%s=%s", key, value)
		*env = append(*env, ev)
	}
	*env = append(*env, "consul=true")
	return nil
}
