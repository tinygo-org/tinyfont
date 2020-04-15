package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
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
	fs       = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	pkgname  = fs.String("package", "main", "package name")
	fontname = fs.String("fontname", "TinyFont", "font name")
	str      = fs.String("string", "", "strings for font")
	output   = fs.String("output", "", "output path")
	ascii    = fs.Bool("ascii", true, "include all glyphs in the font")
	all      = fs.Bool("all", false, "include all ascii glyphs (codepoint <= 255) in the font")
	verbose  = fs.Bool("verbose", false, "run verbosely")
	yadvance = fs.Int("yadvance", 0, "new line distance")
)

type strslice struct {
	s []string
}

func (s *strslice) String() string {
	return fmt.Sprintf("%v", s.s)
}

func (s *strslice) Set(v string) error {
	s.s = append(s.s, v)
	return nil
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: tinyfontgen [flags] [path ...]\n")
	fs.PrintDefaults()
}

// Run ...
func (c *cli) Run(args []string) error {
	var strfiles strslice
	var fonts []string

	fs.Var(&strfiles, "string-file", "strings file for font (can be set multiple times)")

	fs.Usage = usage
	fs.Parse(os.Args[1:])
	for 0 < fs.NArg() {
		fonts = append(fonts, fs.Arg(0))
		fs.Parse(fs.Args()[1:])
	}

	f := fontgen{
		pkgname:  *pkgname,
		fontname: *fontname,
		fonts:    fonts,
	}

	var w io.Writer
	w = os.Stdout
	if len(*output) > 0 {
		ow, err := os.Create(*output)
		if err != nil {
			return nil
		}
		defer ow.Close()

		w = ow
	}

	opts := []option{
		withAscii(*ascii),
		withAll(*all),
		withVerbose(*verbose),
		withYAdvance(*yadvance),
	}

	runes := []rune(*str)
	for _, f := range strfiles.s {
		b, err := ioutil.ReadFile(f)
		if err != nil {
			return err
		}
		runes = append(runes, []rune(string(b))...)
	}

	err := f.generate(w, runes, opts...)
	if err != nil {
		return err
	}

	return nil
}
