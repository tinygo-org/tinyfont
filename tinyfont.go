package tinyfont

import (
	"image/color"

	"github.com/tinygo-org/drivers/hub75"
)

type Glyph struct {
	bitmapIndex uint16
	width       uint8
	height      uint8
	XAdvance    uint8
	xOffset     int8
	yOffset     int8
}

type Font struct {
	Bitmaps  []byte
	Glyphs   []Glyph
	First    byte
	Last     byte
	YAdvance uint8
}

// DrawChar sets a single char in the buffer of the display
func DrawChar(d hub75.Device, font Font, x int16, y int16, char byte, color color.RGBA) {
	if char < font.First || char > font.Last {
		return
	}
	glyph := font.Glyphs[char-font.First]
	bitmapOffset := glyph.bitmapIndex
	bitmap := font.Bitmaps[bitmapOffset]
	bit := uint8(0)
	for j := int16(0); j < int16(glyph.height); j++ {
		for i := int16(0); i < int16(glyph.width); i++ {

			if (bitmap & 0x80) != 0x00 {
				d.SetPixel(x+int16(glyph.xOffset)+i, y+int16(glyph.yOffset)+j, color)
			}
			bitmap <<= 1

			bit++
			if bit > 7 {
				bitmapOffset++
				bitmap = font.Bitmaps[bitmapOffset]
				bit = 0
			}
		}
	}
}

// WriteLine writes a string in the selected font in the buffer
func WriteLine(display hub75.Device, font Font, x int16, y int16, text []byte, color color.RGBA) {
	for i := range text {
		glyph := font.Glyphs[text[i]-font.First]
		DrawChar(display, font, x, y, text[i], color)
		x += int16(glyph.XAdvance)
	}
}

// WriteLineColors writes a string in the selected font in the buffer. Each char is in a different color
// if not enough colors are defined, colors are cycled.
func WriteLineColors(display hub75.Device, font Font, x int16, y int16, text []byte, colors []color.RGBA) {
	numColors := uint16(len(colors))
	if numColors == 0 {
		return
	}

	c := uint16(0)
	for i := range text {
		glyph := font.Glyphs[text[i]-font.First]
		DrawChar(display, font, x, y, text[i], colors[c])
		x += int16(glyph.XAdvance)
		c = (c + 1) % numColors
	}
}
