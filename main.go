package main

import (
	"github.com/benjlevesque/task/cmd"
)

// set by ldflags at build
var (
	version = "<not set>"
	commit  = "<not set>"
	date    = "<not set>"
)

func main() {
	cmd.SetVersion(version, commit, date)
	cmd.Execute()
}
