
## redis-server-docker
### install 
```shell
# Install
$ go install github.com/lemonyxk/redis-server-docker@latest 
```

### usage
```shell
# Help
$ redis-server-docker -h
you can use the following commands for more information:
redis-server-docker -help:[master|slave|sentinel|cluster]

# Help master
$ redis-server-docker -help:master
-type          -- *master or slave or sentinel or cluster
-bind          -- *listen address
-port          -- *listen port
-requirePass   -- *requirepass
-masterAuth    -- *masterauth
-dir           -- data dir
-logFile       -- log file
--daemonize    -- daemonize or not

# Help slave
$ redis-server-docker -help:slave
-type          -- *master or slave or sentinel or cluster
-bind          -- *listen address
-port          -- *listen port
-requirePass   -- *requirepass
-masterAuth    -- *masterauth
-replicaOf     -- *replicaof ip:port
-dir           -- data dir
-logFile       -- log file
--daemonize    -- daemonize or not

# Help sentinel
$ redis-server-docker -help:sentinel
-type          -- *master or slave or sentinel or cluster
-bind          -- *listen address
-port          -- *listen port
-masterAuth    -- *masterauth
-clusterName   -- *cluster name
-masterAddr    -- *master address
-dir           -- data dir
-logFile       -- log file
--daemonize    -- daemonize or not

# Help cluster
$ redis-server-docker -help:cluster
-type          -- *master or slave or sentinel or cluster
-bind          -- *listen address
-port          -- *listen port
-requirePass   -- *requirepass
-masterAuth    -- *masterauth
-dir           -- data dir
-logFile       -- log file
--daemonize    -- daemonize or not

# Create cluster
redis-cli --cluster create [ip1:port1,ip2:port2] --cluster-replicas 1 -a [masterauth]
```