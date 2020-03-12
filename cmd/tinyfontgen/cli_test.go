package main

import (
	"bytes"
	"testing"
)

func TestRunHelpFlag(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	c := &cli{outStream: outStream, errStream: errStream}
	args := []string{appName, `--help`}

	app.Terminate(func(i int) {
		panic("terminated")
	})

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("not terminated")
		}
	}()

	c.Run(args)
	t.Error("not terminated")
}
