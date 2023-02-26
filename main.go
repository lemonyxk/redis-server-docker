package main

import (
	"embed"
	_ "embed"
)

//go:embed conf/*
var conf embed.FS

func main() {

	if HasArgs("-h", "--help") {
		help()
	}

	if HasArgs("-help:master") {
		helpMaster()
	}

	if HasArgs("-help:slave") {
		helpSlave()
	}

	if HasArgs("-help:sentinel") {
		helpSentinel()
	}

	if HasArgs("-help:cluster") {
		helpCluster()
	}

	var nodeType = GetArgs("-type")
	if nodeType == "" {
		exitWithMessage("type is empty")
	}

	switch nodeType {
	case "master":
		master()
	case "slave":
		slave()
	case "sentinel":
		sentinel()
	case "cluster":
		cluster()
	default:
		exitWithMessage("type must be master or slave or sentinel")
	}
}
