package tinyfont // import "tinygo.org/x/tinyfont"

import (
	"image/color"

	"tinygo.org/x/drivers"
)

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

// Draw sets a single glyph in the buffer of the display.
func (glyph Glyph) Draw(display drivers.Displayer, x int16, y int16, color color.RGBA) {
	bitmapOffset := 0
	bitmap := byte(0)
	if len(glyph.Bitmaps) > 0 {
		bitmap = glyph.Bitmaps[bitmapOffset]
	}
	bit := uint8(0)
	for j := int16(0); j < int16(glyph.Height); j++ {
		for i := int16(0); i < int16(glyph.Width); i++ {

			if (bitmap & 0x80) != 0x00 {
				display.SetPixel(x+int16(glyph.XOffset)+i, y+int16(glyph.YOffset)+j, color)
			}
			bitmap <<= 1

			bit++
			if bit > 7 {
				bitmapOffset++
				if bitmapOffset < len(glyph.Bitmaps) {
					bitmap = glyph.Bitmaps[bitmapOffset]
				}
				bit = 0
			}
		}
	}
}

// Info returns glyph information.
func (glyph Glyph) Info() *GlyphInfo {
	return &GlyphInfo{
		Rune:     glyph.Rune,
		Width:    glyph.Width,
		Height:   glyph.Height,
		XAdvance: glyph.XAdvance,
		XOffset:  glyph.XOffset,
		YOffset:  glyph.YOffset,
	}
}

// GetYAdvance returns YAdvance of f.
func (f *Font) GetYAdvance() uint8 {
	return f.YAdvance
}

// GetGlyph returns the glyph corresponding to the specified rune in the font.
func (font *Font) GetGlyph(r rune) Glypher {
	s := 0
	e := len(font.Glyphs) - 1

	for s <= e {
		m := (s + e) / 2

		if font.Glyphs[m].Info().Rune < r {
			s = m + 1
		} else {
			e = m - 1
		}
	}

	if s == len(font.Glyphs) || font.Glyphs[s].Info().Rune != r {
		g := Glyph{
			Rune:     r,
			Width:    0,
			Height:   0,
			XAdvance: font.Glyphs[0].Info().XAdvance,
			XOffset:  font.Glyphs[0].Info().XOffset,
			YOffset:  font.Glyphs[0].Info().YOffset,
			Bitmaps:  []byte{0},
		}
		return g
	}

	return font.Glyphs[s]
}
