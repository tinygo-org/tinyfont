package main

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	appName        = "tinyfontgen"
	appDescription = ""
)

type cli struct {
	outStream io.Writer
	errStream io.Writer
}

var (
	app = kingpin.New(appName, appDescription)
)

// Run ...
func (c *cli) Run(args []string) error {
	app.UsageWriter(c.errStream)

	if VERSION != "" {
		app.Version(fmt.Sprintf("%s version %s build %s", appName, VERSION, BUILDDATE))
	} else {
		app.Version(fmt.Sprintf("%s version - build -", appName))
	}
	app.HelpFlag.Short('h')

	pkgname := app.Flag("package", "package name").Default("main").String()
	fontname := app.Flag("fontname", "font name").Default("TinyFont").String()
	fonts := app.Flag("font", "font path (*.bdf)").Required().ExistingFiles()
	str := app.Arg(`string`, `strings for font`).String()
	output := app.Flag("output", "output path").String()

	k, err := app.Parse(args[1:])
	if err != nil {
		return err
	}

	switch k {
	default:
		f := fontgen{
			pkgname:  *pkgname,
			fontname: *fontname,
			fonts:    *fonts,
		}

		w, err := os.Create(*output)
		if err != nil {
			return nil
		}
		defer w.Close()
		err = f.generate(w, []rune(*str))
		if err != nil {
			return err
		}
	}

	return nil
}
