package main

import (
	"github.com/coreos/go-etcd/etcd"
	"strings"
)

func readFromEtcd(addr string, env *[]string, appname string) error {
	var etcdClient = etcd.NewClient([]string{addr})

	res, err := etcdClient.Get(appname, true, false)

	if err != nil {
		return err
	}

	for _, n := range res.Node.Nodes {
		key := strings.Split(n.Key, "/")
		k, v := strings.ToUpper(key[len(key)-1]), n.Value
		addToEnv(env, k, v)
	}

	addToEnv(env, "etcd", "true")

	return nil
}
