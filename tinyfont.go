package tinyfont // import "tinygo.org/x/tinyfont"

import (
	"image/color"

	"tinygo.org/x/drivers"
)

const (
	NO_ROTATION  Rotation = 0
	ROTATION_90  Rotation = 1 // 90 degrees clock-wise rotation
	ROTATION_180 Rotation = 2
	ROTATION_270 Rotation = 3

	LineFeed       = 0x0A
	CarriageReturn = 0x0D
)

type Rotation uint8

type Glyph struct {
	Rune     rune
	Width    uint8
	Height   uint8
	XAdvance uint8
	XOffset  int8
	YOffset  int8
	Bitmaps  []byte
}

type Font struct {
	Glyphs   []Glyph
	YAdvance uint8
}

type Fonter interface {
	GetGlyph(r rune) Glyph
	GetYAdvance() uint8
}

// DrawChar sets a single rune in the buffer of the display.
func DrawChar(display drivers.Displayer, font Fonter, x int16, y int16, char rune, color color.RGBA) {
	DrawCharRotated(display, font, x, y, char, color, NO_ROTATION)
}

// DrawCharRotated sets a single rune in the buffer of the display.
func DrawCharRotated(display drivers.Displayer, font Fonter, x int16, y int16, char rune, color color.RGBA, rotation Rotation) {
	glyph := font.GetGlyph(char)
	drawGlyphRotated(display, x, y, glyph, color, rotation)
}

// drawGlyphRotated sets a single glyph in the buffer of the display.
func drawGlyphRotated(display drivers.Displayer, x int16, y int16, glyph Glyph, color color.RGBA, rotation Rotation) {
	bitmapOffset := 0
	bitmap := byte(0)
	if len(glyph.Bitmaps) > 0 {
		bitmap = glyph.Bitmaps[bitmapOffset]
	}
	bit := uint8(0)
	for j := int16(0); j < int16(glyph.Height); j++ {
		for i := int16(0); i < int16(glyph.Width); i++ {

			if (bitmap & 0x80) != 0x00 {
				if rotation == NO_ROTATION {
					display.SetPixel(x+int16(glyph.XOffset)+i, y+int16(glyph.YOffset)+j, color)
				} else if rotation == ROTATION_90 {
					display.SetPixel(x-int16(glyph.YOffset)-j, y+int16(glyph.XOffset)+i, color)
				} else if rotation == ROTATION_180 {
					display.SetPixel(x-int16(glyph.XOffset)-i, y-int16(glyph.YOffset)-j, color)
				} else {
					display.SetPixel(x+int16(glyph.YOffset)+j, y-int16(glyph.XOffset)-i, color)
				}
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

// WriteLine writes a string in the selected font in the buffer.
func WriteLine(display drivers.Displayer, font Fonter, x int16, y int16, str string, c color.RGBA) {
	WriteLineColorsRotated(display, font, x, y, str, []color.RGBA{c}, NO_ROTATION)
}

// WriteLineRotated writes a string in the selected font in the buffer.
func WriteLineRotated(display drivers.Displayer, font Fonter, x int16, y int16, str string, c color.RGBA, rotation Rotation) {
	WriteLineColorsRotated(display, font, x, y, str, []color.RGBA{c}, rotation)
}

// WriteLineColors writes a string in the selected font in the buffer. Each char is in a different color
// if not enough colors are defined, colors are cycled.
func WriteLineColors(display drivers.Displayer, font Fonter, x int16, y int16, str string, colors []color.RGBA) {
	WriteLineColorsRotated(display, font, x, y, str, colors, 0)
}

// WriteLineColorsRotated writes a string in the selected font in the buffer. Each char is in a different color
// if not enough colors are defined, colors are cycled.
func WriteLineColorsRotated(display drivers.Displayer, font Fonter, x int16, y int16, str string, colors []color.RGBA, rotation Rotation) {
	text := []rune(str)
	numColors := uint16(len(colors))
	if numColors == 0 {
		return
	}
	rotation = rotation % 4

	c := uint16(0)
	w, h := display.Size()
	l := len(text)
	ox := x
	oy := y
	for i := 0; i < l; i++ {
		if text[i] == LineFeed || text[i] == CarriageReturn {
			/* CR or LF */
			if rotation == NO_ROTATION {
				x = ox
				y += int16(font.GetYAdvance())
			} else if rotation == ROTATION_90 {
				x -= int16(font.GetYAdvance())
				y = oy
			} else if rotation == ROTATION_180 {
				x = ox
				y -= int16(font.GetYAdvance())
			} else {
				x += int16(font.GetYAdvance())
				y = oy
			}
			continue
		}
		glyph := font.GetGlyph(text[i])
		drawGlyphRotated(display, x, y, glyph, colors[c], rotation)
		c++
		if c >= numColors {
			c = 0
		}
		if rotation == NO_ROTATION {
			x += int16(glyph.XAdvance)
		} else if rotation == ROTATION_90 {
			y += int16(glyph.XAdvance)
		} else if rotation == ROTATION_180 {
			x -= int16(glyph.XAdvance)
		} else {
			y -= int16(glyph.XAdvance)
		}

		// speed up?
		if (rotation == NO_ROTATION && x > w) ||
			(rotation == ROTATION_90 && x > h) ||
			(rotation == ROTATION_180 && x < 0) ||
			(rotation == ROTATION_270 && y < 0) {
			break
		}
	}
}

// LineWidth returns the inner and outbox widths corresponding to font and str.
func LineWidth(font Fonter, str string) (innerWidth uint32, outboxWidth uint32) {
	text := []rune(str)
	for i := range text {
		glyph := font.GetGlyph(text[i])
		outboxWidth += uint32(glyph.XAdvance)
	}
	innerWidth = outboxWidth
	// first glyph
	glyph := font.GetGlyph(text[0])
	innerWidth -= uint32(glyph.XOffset)

	// last glyph
	glyph = font.GetGlyph(text[len(text)-1])
	innerWidth += -uint32(glyph.XAdvance) + uint32(glyph.XOffset) + uint32(glyph.Width)
	return
}

// GetGlyph returns the glyph corresponding to the specified rune in the font.
func (f *Font) GetGlyph(r rune) Glyph {
	s := 0
	e := len(f.Glyphs) - 1

	for s <= e {
		m := (s + e) / 2

		if f.Glyphs[m].Rune < r {
			s = m + 1
		} else {
			e = m - 1
		}
	}

	if s == len(f.Glyphs) || f.Glyphs[s].Rune != r {
		g := Glyph{
			Rune:     r,
			Width:    0,
			Height:   0,
			XAdvance: f.Glyphs[0].XAdvance,
			XOffset:  f.Glyphs[0].XOffset,
			YOffset:  f.Glyphs[0].YOffset,
			Bitmaps:  []byte{0},
		}
		return g
	}

	return f.Glyphs[s]
}

func (f *Font) GetYAdvance() uint8 {
	return f.YAdvance
}
