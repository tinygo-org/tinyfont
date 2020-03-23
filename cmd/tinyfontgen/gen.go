package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/hajimehoshi/go-jisx0208"
	"github.com/sago35/go-bdf"
	"tinygo.org/x/tinyfont"
)

type fontgen struct {
	pkgname  string
	fontname string
	fonts    []string
}

func (f *fontgen) generate(w io.Writer, runes []rune) error {
	runes = append(runes, allAscii()...)
	runes = sortAndUniq(runes)

	fonts := []*bdf.Font{}
	for _, font := range f.fonts {
		fn, err := readFont(font)
		if err != nil {
			return err
		}
		fonts = append(fonts, fn)
	}

	ufont := tinyfont.Font{}
	for _, font := range fonts {
		ufont.YAdvance = uint8(font.Size)

		code := func(r rune) (int, error) { return int(r), nil }
		switch strings.ToLower(font.CharsetRegistry) {
		case "iso08859":
		case "jisx0208.1990":
			code = jisx0208.Code
		}

		for _, r := range runes {

			c, err := code(r)
			if err != nil {
				// skip
				continue
			}

			for _, f := range font.Characters {
				if c == int(f.Encoding) {
					img := f.Alpha

					bmp := []byte{}
					bitpos := 0
					b := byte(0)
					for y := img.Rect.Min.Y; y < img.Rect.Max.Y; y++ {
						for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
							r, _, _, _ := img.At(x, y).RGBA()
							if r != 0 {
								b += (byte(1) << (8 - byte(bitpos) - 1))
							}

							bitpos++
							if bitpos == 8 {
								bitpos = 0
								bmp = append(bmp, b)
								b = 0
							}
						}
					}
					if bitpos != 0 {
						bitpos = 0
						bmp = append(bmp, b)
					}

					g := tinyfont.Glyph{
						Width:    uint8(img.Rect.Max.X),
						Height:   uint8(img.Rect.Max.Y),
						XAdvance: uint8(f.Advance[0]),
						XOffset:  int8(f.LowerPoint[0]),
						YOffset:  int8(f.LowerPoint[1]),
						Bitmaps:  bmp,
					}
					ufont.Glyphs = append(ufont.Glyphs, g)
					ufont.RuneToIndex = append(ufont.RuneToIndex, tinyfont.RuneToIndex{
						Rune:  r,
						Index: uint16(len(ufont.Glyphs) - 1),
					})
				}
			}
		}
	}

	fmt.Fprintln(w, `//`, strings.Join(os.Args, ` `))
	fmt.Fprintln(w)
	fmt.Fprintln(w, `package `+f.pkgname)
	fmt.Fprintln(w)
	fmt.Fprintln(w, `import (`)
	fmt.Fprintln(w, `	"tinygo.org/x/tinyfont"`)
	fmt.Fprintln(w, `)`)
	fmt.Fprintln(w)

	fontname := strings.ToUpper(f.fontname[0:1]) + f.fontname[1:]
	fmt.Fprintf(w, "var %s = %T{\n", fontname, ufont)
	fmt.Fprintf(w, "	Glyphs:%T{\n", ufont.Glyphs)
	for i, g := range ufont.Glyphs {
		c := fmt.Sprintf("%c", ufont.RuneToIndex[i].Rune)
		if ufont.RuneToIndex[i].Rune == 0 {
			c = ""
		}
		fmt.Fprintf(w, "		/* %s */ %#v,\n", c, g)
	}
	fmt.Fprintf(w, "	},\n")
	fmt.Fprintln(w)

	fmt.Fprintf(w, "	RuneToIndex:%T{\n", ufont.RuneToIndex)
	for _, rti := range ufont.RuneToIndex {
		c := fmt.Sprintf("%c", rti.Rune)
		if rti.Rune == 0 {
			c = ""
		}
		fmt.Fprintf(w, "		/* %s */ %#v,\n", c, rti)
	}
	fmt.Fprintf(w, "	},\n")
	fmt.Fprintln(w)
	fmt.Fprintf(w, "	YAdvance:%#v,\n", ufont.YAdvance)
	fmt.Fprintf(w, "}\n")

	return nil
}

func readFont(p string) (*bdf.Font, error) {
	buf, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}

	font, err := bdf.Parse(buf)
	if err != nil {
		return nil, err
	}

	return font, nil

}

func allAscii() []rune {
	ret := []rune{}
	for i := 0; i <= 255; i++ {
		ret = append(ret, rune(i))
	}
	return ret
}

func sortAndUniq(runes []rune) []rune {
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })

	ret := []rune{}
	var prev rune
	for _, r := range runes {
		if prev != r {
			ret = append(ret, r)
		}
		prev = r
	}
	return ret
}
