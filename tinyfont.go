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

func DrawChar(d hub75.Device, font Font, x int16, y int16, char byte, color color.RGBA) { //}, bgColor color.RGBA, size uint8) {
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
