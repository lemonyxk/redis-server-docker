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

func sentinel() {
	var bind = GetArgs("-bind")
	if bind == "" {
		exitWithMessage("bind is empty")
	}

	var port = GetArgs("-port")
	if port == "" {
		exitWithMessage("port is empty")
	}

	var masterAuth = GetArgs("-masterAuth")
	if masterAuth == "" {
		exitWithMessage("masterAuth is empty")
	}

	var clusterName = GetArgs("-clusterName")
	if clusterName == "" {
		exitWithMessage("clusterName is empty")
	}

	var masterAddr = GetArgs("-masterAddr")
	if masterAddr == "" {
		exitWithMessage("masterAddr is empty")
	}

	var master = strings.Split(masterAddr, ":")
	if len(master) != 2 {
		exitWithMessage("masterAddr format error")
	}

	var dir = GetArgs("-dir")
	if dir == "" {
		dir = port
	}

	var logFile = GetArgs("-logFile")

	var daemonize = "no"
	if HasArgs("--daemonize") {
		daemonize = "yes"
	}

	var masterIP = master[0]
	var masterPort = master[1]

	var err error
	f, err := conf.Open("conf/sentinel.conf")
	if err != nil {
		exitWithMessage(err)
	}

	defer func() { _ = f.Close() }()

	data, err := io.ReadAll(f)
	if err != nil {
		exitWithMessage(err)
	}

	var sentinelConf = string(data)
	_ = os.Mkdir(port, 0755)
	sentinelConf = strings.ReplaceAll(sentinelConf, "{{daemonize}}", daemonize)
	sentinelConf = strings.ReplaceAll(sentinelConf, "{{dir}}", dir)
	sentinelConf = strings.ReplaceAll(sentinelConf, "{{bind}}", bind)
	sentinelConf = strings.ReplaceAll(sentinelConf, "{{port}}", port)
	sentinelConf = strings.ReplaceAll(sentinelConf, "{{masterAuth}}", masterAuth)
	sentinelConf = strings.ReplaceAll(sentinelConf, "{{clusterName}}", clusterName)
	sentinelConf = strings.ReplaceAll(sentinelConf, "{{masterIP}}", masterIP)
	sentinelConf = strings.ReplaceAll(sentinelConf, "{{masterPort}}", masterPort)
	if logFile != "" {
		sentinelConf = strings.ReplaceAll(sentinelConf, `# logfile "./sentinel.log"`, fmt.Sprintf(`logfile "%s"`, logFile))
	}

	var fileName = filepath.Join(port, "sentinel.conf")
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		exitWithMessage(err)
	}

	_, err = file.WriteString(sentinelConf)
	if err != nil {
		exitWithMessage(err)
	}

	defer func() { _ = file.Close() }()

	defer func() { _ = os.Remove(fileName) }()

	go func() {
		var cmd = Command("redis-sentinel", fileName)
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
