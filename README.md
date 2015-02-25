# Habitat

## What

  Habitat is a golang clone of coreutils *env* that talks to many service discovery systems to create an enviroment and run applications.
  Habitat retrieves key/values from you service discovery/distributed kv and add it to the env passed to applications. It enables dynamic 12 factor app with custom runners.
  It is perfect for containers or applications which need distributed configuration values or quick time to spin with zero touch.


## Usage
Run your application with habitat app_name for a local env, habitat -e for etcd based env and habitat -c for consul based queries.
To eliminate the local env add the -n option. To access a daemon on a remote server use -p host:port
Your keys need to be added in a way where the root name is the appname.

## Build
	$ go get github.com/gleicon/habitat and it will be available at $GOPATH/bin

	If you want to clone the repo, just run
	$ go build

## Options
	- -e etcd <etcd addr:port>
	- -c consul <consul addr:port>
	- -r redis <redis addr:port>
	- -n do not add local env
	- -p host:port for remote servers
	- -k key prefix. If not set the app name will be used as root or prefix depending on the database


## Examples (using env to show the new env vars)
	- etcd
		$ curl http://127.0.0.1:4001/v1/keys/env/db -d value="newdb"
		$ curl http://127.0.0.1:4001/v1/keys/env/cache -d value="newcache"
		$ curl http://127.0.0.1:4001/v1/keys/env/queue -d value="datqueue"
		$ habitat -e 127.0.0.1:4001 env

	- redis
		$ redis-cli hset env db newdb
		$ redis-cli hset env cache newcache
		$ redis-cli hset env queue newqueue
		$ habitat -r 127.0.0.1:6379 env

	- consul
		$ curl -X PUT -d 'newdb' http://localhost:8500/v1/kv/env/db
		$ curl -X PUT -d 'newcache' http://localhost:8500/v1/kv/env/cache
		$ curl -X PUT -d 'newqueue' http://localhost:8500/v1/kv/env/queue
		$ habitat -c 127.0.0.1:8500 env

	- consul with a different app name
		$ curl -X PUT -d 'newdb' http://localhost:8500/v1/kv/newapp/db
		$ curl -X PUT -d 'newcache' http://localhost:8500/v1/kv/newapp/cache
		$ curl -X PUT -d 'newqueue' http://localhost:8500/v1/kv/newapp/queue
		$ habitat -k newapp -c 127.0.0.1:8500 env

	You can mix data coming from all sources too.


## Authors
	Gleicon <gleicon@gmail.com>

## License MIT
