package tinyfont

import (
	"image/color"

	"tinygo.org/x/drivers"
)

type Glyph struct {
	Rune     rune
	Width    uint8
	Height   uint8
	XAdvance uint8
	XOffset  int8
	YOffset  int8
	Bitmaps  []byte
}

// Draw sets a single glyph in the buffer of the display.
func (g Glyph) Draw(display drivers.Displayer, x int16, y int16, color color.RGBA) {
	bitmapOffset := 0
	bitmap := byte(0)
	if len(g.Bitmaps) > 0 {
		bitmap = g.Bitmaps[bitmapOffset]
	}
	bit := uint8(0)

	for j := int16(0); j < int16(g.Height); j++ {
		for i := int16(0); i < int16(g.Width); i++ {

			if (bitmap & 0x80) != 0x00 {
				display.SetPixel(x+int16(g.XOffset)+i, y+int16(g.YOffset)+j, color)
			}
			bitmap <<= 1

			bit++
			if bit > 7 {
				bitmapOffset++
				if bitmapOffset < len(g.Bitmaps) {
					bitmap = g.Bitmaps[bitmapOffset]
				}
				bit = 0
			}
		}
	}
}

// Draw sets a single glyph in the buffer of the display.
func (g Glyph) Info() *GlyphInfo {
	return &GlyphInfo{
		Rune:     g.Rune,
		Width:    g.Width,
		Height:   g.Height,
		XAdvance: g.XAdvance,
		XOffset:  g.XOffset,
		YOffset:  g.YOffset,
	}
}
