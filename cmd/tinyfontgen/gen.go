package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"unsafe"

	"github.com/hajimehoshi/go-jisx0208"
	"github.com/sago35/go-bdf"
	"tinygo.org/x/tinyfont"
)

type fontgen struct {
	pkgname  string
	fontname string
	fonts    []string
}

type option func(*options)

type options struct {
	all      bool
	verbose  bool
	yadvance int
}

var defaultOption = options{
	all:      false,
	verbose:  false,
	yadvance: 0,
}

func withAll(b bool) option {
	return func(o *options) {
		o.all = b
	}
}

func withVerbose(b bool) option {
	return func(o *options) {
		o.verbose = b
	}
}

func withYAdvance(a int) option {
	return func(o *options) {
		o.yadvance = a
	}
}

func (f *fontgen) generate(w io.Writer, runes []rune, opt ...option) error {
	opts := defaultOption
	for _, o := range opt {
		o(&opts)
	}

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
	exists := map[rune]bool{}
	for _, font := range fonts {
		ufont.YAdvance = uint8(float64(font.Size) * float64(font.DPI[1]) / 75)
		if opts.yadvance > 0 {
			ufont.YAdvance = uint8(opts.yadvance)
		}

		code2rune := func(c int) (rune, error) { return rune(c), nil }
		switch strings.ToLower(font.CharsetRegistry) {
		case "iso08859":
		case "jisx0208.1990":
			code2rune = jisx0208.Rune
		}

		for _, f := range font.Characters {
			fr, err := code2rune(int(f.Encoding))
			if err != nil {
				// skip
				continue
			}

			if opts.all {
				if exists[fr] {
					continue
				}
				exists[fr] = true

				g, err := bdf2glyph(f, fr)
				if err != nil {
					continue
				}
				ufont.Glyphs = append(ufont.Glyphs, g)
			} else {
				for _, r := range runes {
					if r == fr {
						if exists[fr] {
							continue
						}
						exists[fr] = true

						g, err := bdf2glyph(f, fr)
						if err != nil {
							continue
						}
						ufont.Glyphs = append(ufont.Glyphs, g)
					}
				}
			}
		}
	}
	ufont.Glyphs = sortGlyphs(ufont.Glyphs)

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
		c := fmt.Sprintf("%c", ufont.Glyphs[i].Rune)
		if ufont.Glyphs[i].Rune == 0 {
			c = ""
		}
		fmt.Fprintf(w, "		/* %s */ %#v,\n", c, g)
	}
	fmt.Fprintf(w, "	},\n")
	fmt.Fprintln(w)

	fmt.Fprintf(w, "	YAdvance:%#v,\n", ufont.YAdvance)
	fmt.Fprintf(w, "}\n")

	if opts.verbose {
		fmt.Printf("Approx. %d bytes\n", calcSize(ufont))
	}

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

func calcSize(f tinyfont.Font) int {
	sz := 0
	for _, g := range f.Glyphs {
		sz += int(unsafe.Sizeof(g.Rune))
		sz += int(unsafe.Sizeof(g.Width))
		sz += int(unsafe.Sizeof(g.Height))
		sz += int(unsafe.Sizeof(g.XAdvance))
		sz += int(unsafe.Sizeof(g.XOffset))
		sz += int(unsafe.Sizeof(g.YOffset))
		sz += len(*(*[]byte)(unsafe.Pointer(&g.Bitmaps)))
	}
	sz += int(unsafe.Sizeof(f.YAdvance))
	return sz

}

func bdf2glyph(f bdf.Character, r rune) (tinyfont.Glyph, error) {
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
		Rune:     r,
		Width:    uint8(img.Rect.Max.X),
		Height:   uint8(img.Rect.Max.Y),
		XAdvance: uint8(f.Advance[0]),
		XOffset:  int8(f.LowerPoint[0]),
		YOffset:  -1 * int8(f.Alpha.Rect.Max.Y+f.LowerPoint[1]),
		Bitmaps:  bmp,
	}

	return g, nil
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

func sortGlyphs(glyphs []tinyfont.Glyph) []tinyfont.Glyph {
	sort.Slice(glyphs, func(i, j int) bool { return glyphs[i].Rune < glyphs[j].Rune })
	return glyphs
}
