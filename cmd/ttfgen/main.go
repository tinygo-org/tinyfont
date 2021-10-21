package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
)

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
	size     = fs.Int("size", 12, "font size")
	dpi      = fs.Int("dpi", 72, "font dpi")
	fonts    []string
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

func main() {
	var strfiles strslice

	fs.Var(&strfiles, "string-file", "strings file for font (can be set multiple times)")

	fs.Usage = usage
	fs.Parse(os.Args[1:])
	for 0 < fs.NArg() {
		fonts = append(fonts, fs.Arg(0))
		fs.Parse(fs.Args()[1:])
	}

	if len(fonts) != 1 {
		log.Fatal("len(fonts) != 1")
	}

	if *yadvance == 0 {
		*yadvance = *size
	}

	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	bb, err := ioutil.ReadFile(fonts[0])
	//bb, err := ioutil.ReadFile("./_font/NotoSansMono-Regular.ttf")
	if err != nil {
		return err
	}

	ft, err := opentype.Parse(bb)
	if err != nil {
		return err
	}
	//fmt.Printf("%#v\n", ft)

	face, err := opentype.NewFace(ft, &opentype.FaceOptions{Size: float64(*size), DPI: float64(*dpi)})
	if err != nil {
		return err
	}

	sfntBuf := &sfnt.Buffer{}
	type xx struct {
		Rune  rune
		Index uint16
	}
	indexes := []xx{}
	max := rune(0x32000 * 2)
	//max = 0x10000
	if !*all {
		max = 0xFF
	}
	for r := rune(0); r < max; r++ {
		idx, err := ft.GlyphIndex(sfntBuf, r)
		if err != nil {
			return err
		}
		if idx != 0 {
			indexes = append(indexes, xx{
				Rune:  r,
				Index: uint16(idx),
			})
			//cnt++
			//fmt.Printf("%08X:%c ", r, r)
		}
	}

	font := TinyfontX{
		Name:   "NotoSans12pt",
		Glyphs: make([]GlyphBuffer, len(indexes)),
	}

	offset := 0
	fontBuffer := [256]uint8{}
	for i, xxx := range indexes {
		font.Glyphs[i].Offset = uint32(offset)
		font.Glyphs[i].Rune = xxx.Rune

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

		for len(fontBuf)%4 != 0 {
			fontBuf = append(fontBuf, 0)
		}

		font.Glyphs[i].Width = uint8(img.Bounds().Max.X - img.Bounds().Min.X)
		font.Glyphs[i].Height = uint8(img.Bounds().Max.Y - img.Bounds().Min.Y)
		font.Glyphs[i].XAdvance = uint8(adv.Ceil())
		font.Glyphs[i].XOffset = int8(dr.Min.X)
		font.Glyphs[i].YOffset = int8(dr.Min.Y)
		font.Glyphs[i].Buf = make([]byte, len(fontBuf))
		copy(font.Glyphs[i].Buf, fontBuf)

		offset += len(fontBuf)/4 + 5
	}

	font.SaveTo(os.Stdout)

	{
		w, err := os.Create(`./ttf-notosans/notosans.bin`)
		if err != nil {
			return err
		}
		defer w.Close()
		font.SaveBinaryTo(w)
	}

	return nil
}

func int2byte(x int) byte {
	if x > 0 {
		return byte(x)
	}

	ret := 0xFFFFFFFF + x
	return byte(ret)
}

func int2byte2(x int8) byte {
	if x > 0 {
		return byte(x)
	}

	ret := 0xFFFFFFFF + int(x)
	return byte(ret)
}

type TinyfontX struct {
	Name   string
	Glyphs []GlyphBuffer
}

func (f TinyfontX) SaveTo(w io.Writer) {
	fmt.Fprintf(w, "package notosans\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "import (\n")
	fmt.Fprintf(w, "	\"tinygo.org/x/tinyfont\"\n")
	fmt.Fprintf(w, ")\n")
	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "var %s = Font{\n", f.Name)
	fmt.Fprintf(w, "	Glyphs: []tinyfont.Glypher{\n")

	fmt.Fprintf(w, "	},\n")
	fmt.Fprintf(w, "	YAdvance: %d,\n", *yadvance)
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "\n")

	//fmt.Fprintf(w, "// \"[Width][Height]\" + \"[XAdvance]\" + \"[XOffset][YOffset]\" + \"[BITMAP_DATA]\"\n")
	//fmt.Fprintf(w, "const c%s = \"\" +\n", f.Name)
	//for _, g := range f.Glyphs {
	//	g.Write1(w)
	//}
	//fmt.Fprintf(w, "	\"\"\n")
	//fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "const m%s = \"\" +\n", f.Name)
	for _, g := range f.Glyphs {
		g.Write3(w)
	}
	fmt.Fprintf(w, "	\"\"\n")
}

func (f TinyfontX) SaveBinaryTo(w io.Writer) {
	for _, g := range f.Glyphs {
		g.Write1bin(w)
	}
}

type GlyphBuffer struct {
	Offset   uint32
	Rune     rune
	Width    uint8
	Height   uint8
	XAdvance uint8
	XOffset  int8
	YOffset  int8
	Buf      []byte
}

func (g *GlyphBuffer) Write1(w io.Writer) {
	switch g.Rune {
	case 0x00, 0x0D, 0xFEFF:
		fmt.Fprintf(w, "	/* '\\x%02X' */ ", g.Rune)
	default:
		fmt.Fprintf(w, "	/* '%c' */ ", g.Rune)
	}

	//fmt.Fprintf(w, "\"\\x%02X\" + ", len(g.Buf)+9)
	//fmt.Fprintf(w, "\"\\x%02X\" + ", 0)
	//fmt.Fprintf(w, "\"\\x%02X\\x%02X\\x%02X\\x%02X\" + ", byte(g.Rune>>24), byte(g.Rune>>16), byte(g.Rune>>8), byte(g.Rune))
	fmt.Fprintf(w, "\"\\x%02X\\x%02X\" + ", g.Width, g.Height)
	fmt.Fprintf(w, "\"\\x%02X\" + ", g.XAdvance)
	fmt.Fprintf(w, "\"\\x%02X\\x%02X\" + ", int2byte2(g.XOffset), int2byte2(g.YOffset))
	fmt.Fprintf(w, "\"")
	for i := 0; i < len(g.Buf); i += 4 {
		fmt.Fprintf(w, "\\x%02X", (g.Buf[i]<<6)+(g.Buf[i+1]<<4)+(g.Buf[i+2]<<2)+g.Buf[i+3])
	}
	fmt.Fprintf(w, "\" +\n")
}

func (g *GlyphBuffer) Write1bin(w io.Writer) {
	binary.Write(w, binary.BigEndian, g.Width)
	binary.Write(w, binary.BigEndian, g.Height)
	binary.Write(w, binary.BigEndian, g.XAdvance)
	binary.Write(w, binary.BigEndian, int2byte2(g.XOffset))
	binary.Write(w, binary.BigEndian, int2byte2(g.YOffset))
	for i := 0; i < len(g.Buf); i += 4 {
		w.Write([]byte{(g.Buf[i] << 6) + (g.Buf[i+1] << 4) + (g.Buf[i+2] << 2) + g.Buf[i+3]})
	}
	//w.Write(g.Buf)
	//binary.Write(w, binary.BigEndian, g.Buf)
}

//func (g *GlyphBuffer) Write2(w io.Writer) {
//	if g.Rune <= 0x20 || g.Rune == 0xFEFF {
//		fmt.Fprintf(w, "		NotoSans12ptGlyph{Offset: 0x%08X, Rune: 0x%08X}, // \\x%02X\n", g.Offset, g.Rune, g.Rune)
//	} else {
//		fmt.Fprintf(w, "		NotoSans12ptGlyph{Offset: 0x%08X, Rune: 0x%08X}, // %c\n", g.Offset, g.Rune, g.Rune)
//	}
//}

func (g *GlyphBuffer) Write3(w io.Writer) {
	switch g.Rune {
	case 0x00, 0x0D, 0xFEFF:
		fmt.Fprintf(w, "	/* '\\x%06X' */ ", g.Rune)
	default:
		fmt.Fprintf(w, "	/* '%c' */ ", g.Rune)
	}

	if g.Offset > 0x00FFFFFF || g.Rune > 0x00FFFFFF {
		panic("overflow")
	}

	fmt.Fprintf(w, "\"\\x%02X\\x%02X\\x%02X\" + \"\\x%02X\\x%02X\\x%02X\" +\n",
		byte(g.Rune>>16),
		byte(g.Rune>>8),
		byte(g.Rune>>0),
		byte(g.Offset>>16),
		byte(g.Offset>>8),
		byte(g.Offset>>0),
	)
}
