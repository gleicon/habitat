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
	fmt.Println("In all the sources the root will be the app name.")
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
	noEnv := flag.Bool("n", false, "No local env")
	flag.Usage = help
	flag.Parse()

	args := flag.Args()
	var env []string
	if len(args) < 1 {
		fmt.Println("No application given")
		help()
		os.Exit(1)
	}

	if *noEnv == false {
		env = os.Environ()
	}

	cmd := args[0]
	bin, lookErr := exec.LookPath(cmd)

	if lookErr != nil {
		panic(lookErr)
	}

	if *fromConsul != "" {
		err := readFromConsul(*fromConsul, &env, cmd)
		if err != nil {
			panic(err)
		}
	}

	if *fromEtcd != "" {
		err := readFromEtcd(*fromEtcd, &env, cmd)
		if err != nil {
			panic(err)
		}
	}

	if *fromRedis != "" {
		err := readFromRedis(*fromRedis, &env, cmd)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(env)
	execErr := syscall.Exec(bin, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
