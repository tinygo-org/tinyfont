package main

import (
	"fmt"
	"os"
)

// VERSION : Tool version
// ex: go build -ldflags="-X main.VERSION=1.2.3"
var VERSION string

// BUILDDATE : Build Date
// ex: go build -ldflags="-X "main.buildDate=2006/01/02 15:04:05 -0700 MST"
var BUILDDATE string

func main() {
	c := &cli{outStream: os.Stdout, errStream: os.Stderr}
	err := c.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
