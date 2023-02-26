/**
* @program: redis-server-docker
*
* @description:
*
* @author: lemo
*
* @create: 2023-02-26 18:00
**/

package main

import (
	"fmt"
	"os"
)

func helpCluster() {
	var str = `
-type          -- *master or slave or sentinel or cluster
-bind          -- *listen address
-port          -- *listen port
-requirePass   -- *requirepass
-masterAuth    -- *masterauth
-dir           -- data dir
-logFile       -- log file
--daemonize    -- daemonize or not
`
	fmt.Println(str[1 : len(str)-1])
	os.Exit(0)
}

func helpSentinel() {
	var str = `
-type          -- *master or slave or sentinel or cluster
-bind          -- *listen address
-port          -- *listen port
-masterAuth    -- *masterauth
-clusterName   -- *cluster name
-masterAddr    -- *master address
-dir           -- data dir
-logFile       -- log file
--daemonize    -- daemonize or not
`
	fmt.Println(str[1 : len(str)-1])
	os.Exit(0)
}

func helpSlave() {
	var str = `
-type          -- *master or slave or sentinel or cluster
-bind          -- *listen address
-port          -- *listen port
-requirePass   -- *requirepass
-masterAuth    -- *masterauth
-replicaOf     -- *replicaof ip:port
-dir           -- data dir
-logFile       -- log file
--daemonize    -- daemonize or not
`
	fmt.Println(str[1 : len(str)-1])
	os.Exit(0)
}

func helpMaster() {
	var str = `
-type          -- *master or slave or sentinel or cluster
-bind          -- *listen address
-port          -- *listen port
-requirePass   -- *requirepass
-masterAuth    -- *masterauth
-dir           -- data dir
-logFile       -- log file
--daemonize    -- daemonize or not
`
	fmt.Println(str[1 : len(str)-1])
	os.Exit(0)
}

func help() {
	var str = `
you can use the following commands for more information:
redis-server-docker -help:[master|slave|sentinel|cluster]
`
	fmt.Println(str[1 : len(str)-1])
	os.Exit(0)
}
