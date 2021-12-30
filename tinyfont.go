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

// GlyphInfo is a structure that holds information about glyphs.
type GlyphInfo struct {
	Rune     rune
	Width    uint8
	Height   uint8
	XAdvance uint8
	XOffset  int8
	YOffset  int8
}

// Fonter is an interface that represents a set of glyphs.
type Fonter interface {
	GetGlyph(r rune) Glypher
	GetYAdvance() uint8
}

// Glypher is glyph itself, and it knows how to draw itself.
type Glypher interface {
	Draw(display drivers.Displayer, x int16, y int16, color color.RGBA)
	Info() *GlyphInfo
}

// DrawChar sets a single rune in the buffer of the display.
func DrawChar(display drivers.Displayer, font Fonter, x int16, y int16, char rune, color color.RGBA) {
	DrawCharRotated(display, font, x, y, char, color, NO_ROTATION)
}

// DrawCharRotated sets a single rune in the buffer of the display.
func DrawCharRotated(display drivers.Displayer, font Fonter, x int16, y int16, char rune, color color.RGBA, rotation Rotation) {
	glyph := font.GetGlyph(char)
	rd := NewRotatedDisplay(display, rotation, x, y)
	glyph.Draw(rd, 0, 0, color)
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
	rd := NewRotatedDisplay(display, rotation, x, y)

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
		glyph.Draw(rd, nx, ny, colors[c])
		c++
		if c >= numColors {
			c = 0
		}
		nx += int16(glyph.Info().XAdvance)
	}
}

// LineWidth returns the inner and outbox widths corresponding to font and str.
func LineWidth(f Fonter, str string) (innerWidth uint32, outboxWidth uint32) {
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

// GetGlyph returns the glyph corresponding to the specified rune in the font.
func GetGlyph(f Fonter, r rune) Glypher {
	return f.GetGlyph(r)
}
