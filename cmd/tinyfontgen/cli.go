package main

import (
	"fmt"
	"io"
	"io/ioutil"
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
	str := app.Flag(`string`, `strings for font`).String()
	strfiles := app.Flag(`string-file`, `strings for font`).ExistingFiles()
	output := app.Flag("output", "output path").String()
	all := app.Flag("all", "include all glyphs in the font").Bool()
	verbose := app.Flag("verbose", "run verbosely").Bool()
	yadvance := app.Flag("yadvance", "new line distance").Int()

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

		opts := []option{
			withAll(*all),
			withVerbose(*verbose),
			withYAdvance(*yadvance),
		}

		runes := []rune(*str)
		for _, f := range *strfiles {
			b, err := ioutil.ReadFile(f)
			if err != nil {
				return err
			}
			runes = append(runes, []rune(string(b))...)
		}

		err = f.generate(w, runes, opts...)
		if err != nil {
			return err
		}
	}

	return nil
}
