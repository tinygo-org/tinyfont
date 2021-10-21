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

type GlyphInfo struct {
	Rune     rune
	Width    uint8
	Height   uint8
	XAdvance uint8
	XOffset  int8
	YOffset  int8
}

type Glypher interface {
	Draw(display drivers.Displayer, x int16, y int16, color color.RGBA)
	Info() *GlyphInfo
}

type Font struct {
	BBox     [4]int8 // width, height, minX, minY
	Glyphs   []Glypher
	YAdvance uint8
}

// DrawChar sets a single rune in the buffer of the display.
func DrawChar(display drivers.Displayer, font Fonter, x int16, y int16, char rune, color color.RGBA) {
	DrawCharRotated(display, font, x, y, char, color, NO_ROTATION)
}

// DrawCharRotated sets a single rune in the buffer of the display.
func DrawCharRotated(display drivers.Displayer, font Fonter, x int16, y int16, char rune, color color.RGBA, rotation Rotation) {
	glyph := font.GetGlyph(char)
	display = NewRotatedLabel(display, rotation, x, y)
	glyph.Draw(display, 0, 0, color)
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
	display = NewRotatedLabel(display, rotation, x, y)

	c := uint16(0)
	l := len(text)
	nx := int16(0)
	ny := int16(0)
	for i := 0; i < l; i++ {
		if text[i] == LineFeed || text[i] == CarriageReturn {
			/* CR or LF */
			nx = 0
			ny += int16(font.GetYAdvance())
			continue
		}
		glyph := font.GetGlyph(text[i])
		glyph.Draw(display, nx, ny, colors[c])
		c++
		if c >= numColors {
			c = 0
		}
		nx += int16(glyph.Info().XAdvance)
	}
}

type Fonter interface {
	GetGlyph(r rune) Glypher
	GetYAdvance() uint8
	LineWidth(str string) (innerWidth uint32, outboxWidth uint32)
}

// LineWidth returns the inner and outbox widths corresponding to font and str.
func (f *Font) LineWidth(str string) (innerWidth uint32, outboxWidth uint32) {
	text := []rune(str)
	if len(text) == 0 {
		return
	}

	for i := range text {
		glyph := f.GetGlyph(text[i])
		outboxWidth += uint32(glyph.Info().XAdvance)
	}
	innerWidth = outboxWidth
	// first glyph
	glyph := f.GetGlyph(text[0])
	innerWidth -= uint32(glyph.Info().XOffset)

	// last glyph
	glyph = f.GetGlyph(text[len(text)-1])
	innerWidth += -uint32(glyph.Info().XAdvance) + uint32(glyph.Info().XOffset) + uint32(glyph.Info().Width)
	return
}

func (f *Font) GetYAdvance() uint8 {
	return f.YAdvance
}

// GetGlyph returns the glyph corresponding to the specified rune in the font.
func (f *Font) GetGlyph(r rune) Glypher {
	s := 0
	e := len(f.Glyphs) - 1

	for s <= e {
		m := (s + e) / 2

		if f.Glyphs[m].Info().Rune < r {
			s = m + 1
		} else {
			e = m - 1
		}
	}

	if s == len(f.Glyphs) || f.Glyphs[s].Info().Rune != r {
		g := &Glyph{
			Rune:     r,
			Width:    0,
			Height:   0,
			XAdvance: f.Glyphs[0].Info().XAdvance,
			XOffset:  f.Glyphs[0].Info().XOffset,
			YOffset:  f.Glyphs[0].Info().YOffset,
			Bitmaps:  []byte{0},
		}
		return g
	}

	return f.Glyphs[s]
}

// LineWidth returns the inner and outbox widths corresponding to font and str.
func LineWidth(f Fonter, str string) (innerWidth uint32, outboxWidth uint32) {
	return f.LineWidth(str)
}
