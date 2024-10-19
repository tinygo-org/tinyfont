// Package const1bit is for fonts that use 2 bits per pixel.
package const1bit

import (
	"image/color"

	"tinygo.org/x/drivers"
	"tinygo.org/x/tinyfont"
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

	OffsetMap string
	Data      string
	Name      string
	glyph     Glyph
}

// Draw sets a single glyph in the buffer of the display.
func (glyph Glyph) Draw(display drivers.Displayer, x int16, y int16, c color.RGBA) {
	bitmapOffset := 0
	bitmap := byte(0)
	if len(glyph.Bitmaps) > 0 {
		bitmap = glyph.Bitmaps[bitmapOffset]
	}
	bit := uint8(0)
	for j := int16(0); j < int16(glyph.Height); j++ {
		for i := int16(0); i < int16(glyph.Width); i++ {

			if (bitmap & 0x80) != 0x00 {
				display.SetPixel(x+int16(glyph.XOffset)+i, y+int16(glyph.YOffset)+j, c)
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
func (glyph Glyph) Info() tinyfont.GlyphInfo {
	return tinyfont.GlyphInfo{
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
// Since there is only one glyph used for the return value in this package,
// concurrent access is not allowed. Normally, there is no issue when using it from tinyfont.
func (font *Font) GetGlyph(r rune) tinyfont.Glypher {
	s := 0
	e := len(font.OffsetMap)/6 - 1

	for s <= e {
		m := (s + e) / 2

		r2 := rune(font.OffsetMap[m*6])<<16 + rune(font.OffsetMap[m*6+1])<<8 + rune(font.OffsetMap[m*6+2])
		if r2 < r {
			s = m + 1
		} else {
			e = m - 1
		}
	}

	if s > len(font.OffsetMap)/6-1 {
		s = 0
	}

	offset := uint32(font.OffsetMap[s*6+3])<<16 + uint32(font.OffsetMap[s*6+4])<<8 + uint32(font.OffsetMap[s*6+5])
	sz := uint32(len(font.Data[offset+5:]))
	if s*6+6 < len(font.OffsetMap) {
		sz = uint32(font.OffsetMap[s*6+9])<<16 + uint32(font.OffsetMap[s*6+10])<<8 + uint32(font.OffsetMap[s*6+11]) - offset
	}

	font.glyph.Rune = r
	font.glyph.Width = font.Data[offset+0]
	font.glyph.Height = font.Data[offset+1]
	font.glyph.XAdvance = font.Data[offset+2]
	font.glyph.XOffset = int8(font.Data[offset+3])
	font.glyph.YOffset = int8(font.Data[offset+4])
	font.glyph.Bitmaps = []byte(font.Data[offset+5 : offset+5+sz])
	return &(font.glyph)
}
