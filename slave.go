/**
* @program: redis-sentinel-docker
*
* @description:
*
* @author: lemo
*
* @create: 2023-02-26 15:46
**/

package main

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
)

func slave() {
	var bind = GetArgs("-bind")
	if bind == "" {
		exitWithMessage("bind is empty")
	}

	var port = GetArgs("-port")
	if port == "" {
		exitWithMessage("port is empty")
	}

	var requirePass = GetArgs("-requirePass")
	if requirePass == "" {
		exitWithMessage("requirePass is empty")
	}

	var masterAuth = GetArgs("-masterAuth")
	if masterAuth == "" {
		exitWithMessage("masterAuth is empty")
	}

	var replicaOf = GetArgs("-replicaOf")
	if replicaOf == "" {
		exitWithMessage("replicaOf is empty")
	}

	var replica = strings.Split(replicaOf, ":")
	if len(replica) != 2 {
		exitWithMessage("replicaOf format error")
	}

	var replicaOfIP = replica[0]
	var replicaOfPort = replica[1]

	var dir = GetArgs("-dir")
	if dir == "" {
		dir = port
	}

	var logFile = GetArgs("-logFile")

	var daemonize = "no"
	if HasArgs("--daemonize") {
		daemonize = "yes"
	}

	var err error
	f, err := conf.Open("conf/slave.conf")
	if err != nil {
		exitWithMessage(err)
	}

	defer func() { _ = f.Close() }()

	data, err := io.ReadAll(f)
	if err != nil {
		exitWithMessage(err)
	}

	var slaveConf = string(data)
	_ = os.Mkdir(port, 0755)
	slaveConf = strings.ReplaceAll(slaveConf, "{{daemonize}}", daemonize)
	slaveConf = strings.ReplaceAll(slaveConf, "{{dir}}", dir)
	slaveConf = strings.ReplaceAll(slaveConf, "{{bind}}", bind)
	slaveConf = strings.ReplaceAll(slaveConf, "{{port}}", port)
	slaveConf = strings.ReplaceAll(slaveConf, "{{masterAuth}}", masterAuth)
	slaveConf = strings.ReplaceAll(slaveConf, "{{requirePass}}", requirePass)
	slaveConf = strings.ReplaceAll(slaveConf, "{{replicaOfIP}}", replicaOfIP)
	slaveConf = strings.ReplaceAll(slaveConf, "{{replicaOfPort}}", replicaOfPort)
	if logFile != "" {
		slaveConf = strings.ReplaceAll(slaveConf, `# logfile "./redis.log"`, fmt.Sprintf(`logfile "%s"`, logFile))
	}

	var fileName = filepath.Join(port, "redis.conf")
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		exitWithMessage(err)
	}

	_, err = file.WriteString(slaveConf)
	if err != nil {
		exitWithMessage(err)
	}

	defer func() { _ = file.Close() }()

	go func() {
		var cmd = Command("redis-server", fileName)
		pwd, _ := os.Getwd()
		cmd.Dir = pwd
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout

		err = cmd.Run()
		if err != nil {
			exitWithMessage(err)
		}
	}()

	// listen signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-c
}
