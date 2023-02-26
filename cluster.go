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

func cluster() {

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
	f, err := conf.Open("conf/cluster.conf")
	if err != nil {
		exitWithMessage(err)
	}

	defer func() { _ = f.Close() }()

	data, err := io.ReadAll(f)
	if err != nil {
		exitWithMessage(err)
	}

	var clusterConf = string(data)
	_ = os.Mkdir(port, 0755)
	clusterConf = strings.ReplaceAll(clusterConf, "{{daemonize}}", daemonize)
	clusterConf = strings.ReplaceAll(clusterConf, "{{dir}}", dir)
	clusterConf = strings.ReplaceAll(clusterConf, "{{bind}}", bind)
	clusterConf = strings.ReplaceAll(clusterConf, "{{port}}", port)
	clusterConf = strings.ReplaceAll(clusterConf, "{{masterAuth}}", masterAuth)
	clusterConf = strings.ReplaceAll(clusterConf, "{{requirePass}}", requirePass)
	clusterConf = strings.ReplaceAll(clusterConf, "{{nodesConfig}}", "nodes.conf")
	if logFile != "" {
		clusterConf = strings.ReplaceAll(clusterConf, `# logfile "./cluster.log"`, fmt.Sprintf(`logfile "%s"`, logFile))
	}

	var fileName = filepath.Join(port, "cluster.conf")
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		exitWithMessage(err)
	}

	_, err = file.WriteString(clusterConf)
	if err != nil {
		exitWithMessage(err)
	}

	defer func() { _ = file.Close() }()

	defer func() { _ = os.Remove(fileName) }()

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
