package quanpixel8

import (
	"image/color"

	"tinygo.org/x/drivers"
	"tinygo.org/x/tinyfont"
)

// Font is an 8x8 fixed-cell bitmap font derived from QuanPixel.
//
// The wrapper code follows the tinyfont project license.
// The embedded font data is derived from QuanPixel and is distributed under
// the SIL Open Font License 1.1. See LICENSE_OFL.txt and NOTICE.md.
type Font struct{}

// QP8 is an 8x8 bitmap font derived from QuanPixel.
//
// The name avoids using the reserved font name "QuanPixel" as the font
// variable name for the generated/converted font data.
var QP8 Font

type glyph8 struct {
	r     rune
	index uint32
}

const noGlyph = ^uint32(0)

var emptyGlyph = glyph8{r: ' ', index: noGlyph}

func (Font) GetYAdvance() uint8 {
	return 8
}

func (Font) GetGlyph(r rune) tinyfont.Glypher {
	lo, hi := 0, len(glyphs)
	for lo < hi {
		mid := (lo + hi) >> 1
		if glyphs[mid].r < r {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	if lo < len(glyphs) && glyphs[lo].r == r {
		return &glyphs[lo]
	}

	return &emptyGlyph
}

func (g *glyph8) Info() tinyfont.GlyphInfo {
	return tinyfont.GlyphInfo{
		Rune:     g.r,
		Width:    8,
		Height:   8,
		XAdvance: 8,
		XOffset:  0,
		YOffset:  0,
	}
}

func (g *glyph8) Draw(display drivers.Displayer, x, y int16, c color.RGBA) {
	if g.index == noGlyph {
		return
	}

	off := int(g.index) * 8

	for yy := int16(0); yy < 8; yy++ {
		row := bitmaps[off+int(yy)]

		for xx := int16(0); xx < 8; xx++ {
			if row&(0x80>>uint8(xx)) != 0 {
				display.SetPixel(x+xx, y+yy, c)
			}
		}
	}
}
