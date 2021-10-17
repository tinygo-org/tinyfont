package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	bb, err := ioutil.ReadFile("./_font/NotoSansJP-Regular.otf")
	//bb, err := ioutil.ReadFile("./_font/NotoSansMono-Regular.ttf")
	if err != nil {
		return err
	}

	ft, err := opentype.Parse(bb)
	if err != nil {
		return err
	}
	//fmt.Printf("%#v\n", ft)

	face, err := opentype.NewFace(ft, &opentype.FaceOptions{Size: 12, DPI: 72})
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
	if len(os.Args) < 1 {
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
		font.Glyphs[i].Index = uint32(offset)
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

		offset += len(fontBuf)/4 + 9 + 1
	}

	font.SaveTo(os.Stdout)

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

	fmt.Fprintf(w, "var %s = tinyfont.Font{\n", f.Name)
	fmt.Fprintf(w, "	Glyphs: []tinyfont.Glypher{\n")

	for _, g := range f.Glyphs {
		g.Write2(w)
	}
	fmt.Fprintf(w, "	},\n")
	fmt.Fprintf(w, "	YAdvance: 12,\n")
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "const c%s = \"\" +\n", f.Name)
	for _, g := range f.Glyphs {
		g.Write1(w)
	}
	fmt.Fprintf(w, "	\"\"\n")
}

func (f TinyfontX) SaveBinaryTo(w io.Writer) {
}

type GlyphBuffer struct {
	Index    uint32
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
	case 0x00, 0x0D:
		fmt.Fprintf(w, "	/* '\\x%02X' */ ", g.Rune)
	default:
		fmt.Fprintf(w, "	/* '%c' */ ", g.Rune)
	}

	fmt.Fprintf(w, "\"\\x%02X\" + ", len(g.Buf)+9)
	fmt.Fprintf(w, "\"\\x%02X\\x%02X\\x%02X\\x%02X\" + ", byte(g.Rune>>24), byte(g.Rune>>16), byte(g.Rune>>8), byte(g.Rune))
	fmt.Fprintf(w, "\"\\x%02X\\x%02X\" + ", g.Width, g.Height)
	fmt.Fprintf(w, "\"\\x%02X\" + ", g.XAdvance)
	fmt.Fprintf(w, "\"\\x%02X\\x%02X\" + ", int2byte2(g.XOffset), int2byte2(g.YOffset))
	fmt.Fprintf(w, "\"")
	for i := 0; i < len(g.Buf); i += 4 {
		fmt.Fprintf(w, "\\x%02X", (g.Buf[i]<<6)+(g.Buf[i+1]<<4)+(g.Buf[i+2]<<2)+g.Buf[i+3])
	}
	fmt.Fprintf(w, "\" +\n")
}

func (g *GlyphBuffer) Write2(w io.Writer) {
	if g.Rune <= 0x20 {
		fmt.Fprintf(w, "		NotosansGlyph{Index: 0x%08X, Rune: 0x%08X}, // \\x%02X\n", g.Index, g.Rune, g.Rune)
	} else {
		fmt.Fprintf(w, "		NotosansGlyph{Index: 0x%08X, Rune: 0x%08X}, // %c\n", g.Index, g.Rune, g.Rune)
	}
}
