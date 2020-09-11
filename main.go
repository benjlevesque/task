package main

import (
	"flag"

	"github.com/benjlevesque/task/cmd"
)

func main() {
	var docs bool
	var docPath string
	flag.BoolVar(&docs, "docs", false, "Generates doc")
	flag.StringVar(&docPath, "docs-path", "./docs", "Docs path")
	flag.Parse()

	if docs {
		cmd.Docs(docPath)
	} else {
		cmd.Execute()
	}
}
