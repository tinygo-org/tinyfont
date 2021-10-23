package twemoji

import (
	"image/color"

	"tinygo.org/x/drivers"
	"tinygo.org/x/tinyfont"
)

var Twemoji12pt = Font{
	Glyphs:   []tinyfont.Glypher{},
	YAdvance: 12,
}

type Twemoji12ptGlyph struct {
	Offset uint32
	Rune   rune
}

func (g Twemoji12ptGlyph) Draw(display drivers.Displayer, x int16, y int16, c color.RGBA) {
	//idx := 12 * 12 * 2 * g.Offset
	info := g.Info()
	buf := g.Bytes()
	idx := 0
	for yy := int16(0); yy < 12; yy++ {
		for xx := int16(0); xx < 12; xx++ {
			if true {
				rgb565 := uint16(buf[idx])<<8 + uint16(buf[idx+1])
				display.SetPixel(xx+x+int16(info.XOffset), yy+y+int16(info.YOffset), RGB565ToRGBA(rgb565))
			}
			idx += 2
		}
	}
}

func (g Twemoji12ptGlyph) Info() *tinyfont.GlyphInfo {
	return &tinyfont.GlyphInfo{
		Rune:     g.Rune,
		Width:    12,
		Height:   12,
		XAdvance: 12,
		XOffset:  0,
		YOffset:  -12,
	}
}

func RGB565ToRGBA(c uint16) color.RGBA {
	return color.RGBA{
		R: uint8((c & 0xF800) >> 8),
		G: uint8((c & 0x07E0) >> 3),
		B: uint8((c & 0x001F) << 3),
		A: 0xFF,
	}
}
