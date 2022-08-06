package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
)

var (
	fs       = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	pkgname  = fs.String("package", "main", "package name")
	fontname = fs.String("fontname", "TinyFont", "font name")
	sz       = fs.Int("size", 12, "size for font")
	dpi      = fs.Int("dpi", 75, "dpi for font")
	str      = fs.String("string", "", "strings for font")
	output   = fs.String("output", "", "output path")
	ascii    = fs.Bool("ascii", true, "include all glyphs in the font")
	all      = fs.Bool("all", false, "include all ascii glyphs (codepoint <= 255) in the font")
	verbose  = fs.Bool("verbose", false, "run verbosely")
	yadvance = fs.Int("yadvance", 0, "new line distance")

	runesMap map[rune]bool
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
	fmt.Fprintf(os.Stderr, "usage: tinyfontgen-ttf [flags] [path ...]\n")
	fs.PrintDefaults()
}

func main() {
	var strfiles strslice
	var fonts []string

	fs.Var(&strfiles, "string-file", "strings file for font (can be set multiple times)")

	fs.Usage = usage
	fs.Parse(os.Args[1:])
	for 0 < fs.NArg() {
		fonts = append(fonts, fs.Arg(0))
		fs.Parse(fs.Args()[1:])
	}

	if *yadvance == 0 {
		*yadvance = *sz
	}

	runesMap = map[rune]bool{}
	if *ascii {
		for i := rune(0); i < 256; i++ {
			runesMap[i] = true
		}
	}
	for _, r := range *str {
		runesMap[r] = true
	}
	for _, f := range strfiles.s {
		b, err := ioutil.ReadFile(f)
		if err != nil {
			log.Fatal(err)
		}
		for _, r := range string(b) {
			runesMap[r] = true
		}
	}

	err := run(*sz, *dpi, fonts[0])
	if err != nil {
		log.Fatal(err)
	}
}

func run(size, dpi int, fontfile string) error {
	bb, err := ioutil.ReadFile(fontfile)
	if err != nil {
		return err
	}

	ft, err := opentype.Parse(bb)
	if err != nil {
		return err
	}

	face, err := opentype.NewFace(ft, &opentype.FaceOptions{Size: float64(size), DPI: float64(dpi)})
	if err != nil {
		return err
	}

	sfntBuf := &sfnt.Buffer{}
	type xx struct {
		Rune  rune
		Index uint16
	}
	indexes := []xx{}

	const runeMax = 0x3FFFF
	for r := rune(0); r < runeMax; r++ {
		idx, err := ft.GlyphIndex(sfntBuf, r)
		if err != nil {
			return err
		}
		if idx != 0 && (*all || runesMap[r]) {
			indexes = append(indexes, xx{
				Rune:  r,
				Index: uint16(idx),
			})
		}
	}

	fontBuffer := [256]uint8{}
	font := Font{
		Glyphs: make([]Glyph, len(indexes)),
	}
	for i, xxx := range indexes {
		dr, img, _, adv, ok := face.Glyph(fixed.Point26_6{}, xxx.Rune)
		if !ok {
			continue
		}

		// font
		fontBuf := fontBuffer[:0]
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				z := img.At(x, y)
				r, _, _, _ := z.RGBA()
				fontBuf = append(fontBuf, byte((r&0xC000)>>14))
			}
		}

		switch len(fontBuf) % 4 {
		case 0:
			// skip
		case 1:
			fontBuf = append(fontBuf, 0, 0, 0)
		case 2:
			fontBuf = append(fontBuf, 0, 0)
		case 3:
			fontBuf = append(fontBuf, 0)
		}

		font.Glyphs[i].Rune = xxx.Rune
		font.Glyphs[i].Width = uint8(img.Bounds().Max.X - img.Bounds().Min.X)
		font.Glyphs[i].Height = uint8(img.Bounds().Max.Y - img.Bounds().Min.Y)
		font.Glyphs[i].XAdvance = uint8(adv.Ceil())
		font.Glyphs[i].XOffset = int8(dr.Min.X)
		font.Glyphs[i].YOffset = int8(dr.Min.Y)
		//font.Glyphs[i].Bitmaps = make([]byte, len(fontBuf))
		//copy(font.Glyphs[i].Bitmaps, fontBuf)
		font.Glyphs[i].Bitmaps = make([]byte, len(fontBuf)/4)
		for j := range font.Glyphs[i].Bitmaps {
			font.Glyphs[i].Bitmaps[j] = fontBuf[j*4] << 6
			font.Glyphs[i].Bitmaps[j] += fontBuf[j*4+1] << 4
			font.Glyphs[i].Bitmaps[j] += fontBuf[j*4+2] << 2
			font.Glyphs[i].Bitmaps[j] += fontBuf[j*4+3]
		}
	}

	var w io.Writer
	w = os.Stdout
	if *output != "" {
		os.MkdirAll(filepath.Dir(*output), 0666)
		wc, err := os.Create(*output)
		if err != nil {
			return err
		}
		defer wc.Close()
		w = wc
	}
	font.SaveTo(w)

	return nil
}

// Glyph is a struct that implements Glypher interface.
type Glyph struct {
	Rune     rune
	Width    uint8
	Height   uint8
	XAdvance uint8
	XOffset  int8
	YOffset  int8
	Bitmaps  []byte
}

// Font is a struct that implements Fonter interface.
type Font struct {
	BBox     [4]int8 // width, height, minX, minY
	Glyphs   []Glyph
	YAdvance uint8
}

func (f Font) SaveTo(w io.Writer) {
	pkg := *pkgname
	fontname := *fontname
	yadv := *yadvance

	fmt.Fprintf(w, "// ttfgen %s\n", strings.Join(os.Args[1:], " "))
	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "package %s\n", pkg)
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "import (\n")
	fmt.Fprintf(w, "	\"tinygo.org/x/tinyfont/const2bit\"\n")
	fmt.Fprintf(w, ")\n")
	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "var %s = const2bit.Font{\n", fontname)
	fmt.Fprintf(w, "	OffsetMap: m%s,\n", fontname)
	fmt.Fprintf(w, "	Data:      d%s,\n", fontname)
	fmt.Fprintf(w, "	YAdvance:  %d,\n", yadv)
	fmt.Fprintf(w, "	Name:      %q,\n", fontname)
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "\n")

	offset := 0
	fmt.Fprintf(w, "// rune (3byte) + offset (3byte)\n")
	fmt.Fprintf(w, "const m%s = \"\" +\n", fontname)
	for _, x := range f.Glyphs {
		fmt.Fprintf(w, `	"\x%02X\x%02X\x%02X" + `, byte(x.Rune>>16), byte(x.Rune>>8), byte(x.Rune))
		fmt.Fprintf(w, `"\x%02X\x%02X\x%02X" + `, byte(offset>>16), byte(offset>>8), byte(offset))
		if x.Rune > 0 {
			fmt.Fprintf(w, `// %c`, x.Rune)
		} else {
			fmt.Fprintf(w, `//`)
		}
		fmt.Fprintf(w, "\n")

		// width + height + xadvance + xoffset + yoffset + len([]bitmaps)
		offset += 1 + 1 + 1 + 1 + 1 + len(x.Bitmaps)
	}
	fmt.Fprintf(w, "	\"\"\n")

	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "// width + height + xadvance + xoffset + yoffset + []bitmaps\n")
	fmt.Fprintf(w, "const d%s = \"\" +\n", fontname)
	for _, x := range f.Glyphs {
		fmt.Fprintf(w, `	"\x%02X\x%02X\x%02X\x%02X\x%02X" + `, x.Width, x.Height, x.XAdvance, byte(x.XOffset), byte(x.YOffset))
		fmt.Fprintf(w, `"`)
		for _, y := range x.Bitmaps {
			fmt.Fprintf(w, `\x%02X`, y)
		}
		fmt.Fprintf(w, `" +`)
		fmt.Fprintf(w, "\n")
	}
	fmt.Fprintf(w, "	\"\"\n")
}
