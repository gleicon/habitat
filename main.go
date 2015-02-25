package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func help() {
	fmt.Println("Usage: habitat [-c <ip:port>] [-e <ip:port>] [-r <ip:port>] [-n] <application>")
	fmt.Println("Habitat is just live env but can read pairs of k/v as env vars from remote sources")
	fmt.Println("By default, In all the sources the root key or prefix will be the app name.")
	fmt.Println("To use a custom prefix -k <prefix>")
	fmt.Println("To read a list of env vars from consul -c 127.0.0.1:8500")
	fmt.Println("To read a list of env vars from etcd -e 127.0.0.1:4001")
	fmt.Println("To read a list of env vars from a hash on redis -r 127.0.0.1:6379")
	fmt.Println("All sources can be mixed plus the local env. To skip the local env use -n")
	os.Exit(1)
}

func main() {
	fromConsul := flag.String("c", "", "Read from consul")
	fromEtcd := flag.String("e", "", "Read from etcd")
	fromRedis := flag.String("r", "", "Read from Redis")
	skipLocalEnv := flag.Bool("n", false, "No local env")
	keyPrefix := flag.String("k", "", "A key prefix other than appname")
	flag.Usage = help
	flag.Parse()

	args := flag.Args()
	var env []string
	if len(args) < 1 {
		fmt.Println("No application given")
		help()
		os.Exit(1)
	}

	if *skipLocalEnv == false {
		env = os.Environ()
	}

	cmd := args[0]

	bin, lookErr := exec.LookPath(cmd)

	if lookErr != nil {
		panic(lookErr)
	}

	if *keyPrefix == "" {
		*keyPrefix = cmd
	}

	if *fromConsul != "" {
		err := readFromConsul(*fromConsul, &env, *keyPrefix)
		if err != nil {
			panic(err)
		}
	}

	if *fromEtcd != "" {
		err := readFromEtcd(*fromEtcd, &env, *keyPrefix)
		if err != nil {
			panic(err)
		}
	}

	if *fromRedis != "" {
		err := readFromRedis(*fromRedis, &env, *keyPrefix)
		if err != nil {
			panic(err)
		}
	}

	execErr := syscall.Exec(bin, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
