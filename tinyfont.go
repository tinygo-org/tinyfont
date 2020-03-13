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
)

type Rotation uint8

type Glyph struct {
	BitmapIndex uint16
	Width       uint8
	Height      uint8
	XAdvance    uint8
	XOffset     int8
	YOffset     int8
}

type Font struct {
	Bitmaps  []byte
	Glyphs   []Glyph
	First    byte
	Last     byte
	YAdvance uint8
}

// DrawChar sets a single char in the buffer of the display
func DrawChar(display drivers.Displayer, font *Font, x int16, y int16, char byte, color color.RGBA) {
	DrawCharRotated(display, font, x, y, char, color, 0)
}

// DrawCharRotated sets a single char in the buffer of the display
func DrawCharRotated(display drivers.Displayer, font *Font, x int16, y int16, char byte, color color.RGBA, rotation Rotation) {
	if char < font.First || char > font.Last {
		return
	}
	glyph := font.Glyphs[char-font.First]
	bitmapOffset := glyph.BitmapIndex
	bitmap := font.Bitmaps[bitmapOffset]
	bit := uint8(0)
	for j := int16(0); j < int16(glyph.Height); j++ {
		for i := int16(0); i < int16(glyph.Width); i++ {

			if (bitmap & 0x80) != 0x00 {
				if rotation == 0 {
					display.SetPixel(x+int16(glyph.XOffset)+i, y+int16(glyph.YOffset)+j, color)
				} else if rotation == 1 {
					display.SetPixel(x-int16(glyph.YOffset)-j, y+int16(glyph.XOffset)+i, color)
				} else if rotation == 2 {
					display.SetPixel(x-int16(glyph.XOffset)-i, y-int16(glyph.YOffset)-j, color)
				} else {
					display.SetPixel(x+int16(glyph.YOffset)+j, y-int16(glyph.XOffset)-i, color)
				}
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
func WriteLine(display drivers.Displayer, font *Font, x int16, y int16, text []byte, color color.RGBA) {
	WriteLineRotated(display, font, x, y, text, color, 0)
}

// WriteLineRotated writes a string in the selected font in the buffer
func WriteLineRotated(display drivers.Displayer, font *Font, x int16, y int16, text []byte, color color.RGBA, rotation Rotation) {
	rotation = rotation % 4
	w, h := display.Size()
	l := len(text)
	for i := 0; i < l; i++ {
		glyph := font.Glyphs[text[i]-font.First]
		//if x+int16(glyph.XAdvance) >= 0 {
		DrawCharRotated(display, font, x, y, text[i], color, rotation)
		//}
		if rotation == 0 {
			x += int16(glyph.XAdvance)
		} else if rotation == 1 {
			y += int16(glyph.XAdvance)
		} else if rotation == 2 {
			x -= int16(glyph.XAdvance)
		} else {
			y -= int16(glyph.XAdvance)
		}

		// speed up?
		if (rotation == 0 && x > w) ||
			(rotation == 1 && x > h) ||
			(rotation == 2 && x < 0) ||
			(rotation == 3 && y < 0) {
			break
		}
	}
}

// WriteLineColors writes a string in the selected font in the buffer. Each char is in a different color
// if not enough colors are defined, colors are cycled.
func WriteLineColors(display drivers.Displayer, font *Font, x int16, y int16, text []byte, colors []color.RGBA) {
	WriteLineColorsRotated(display, font, x, y, text, colors, 0)
}

// WriteLineColorsRotated writes a string in the selected font in the buffer. Each char is in a different color
// if not enough colors are defined, colors are cycled.
func WriteLineColorsRotated(display drivers.Displayer, font *Font, x int16, y int16, text []byte, colors []color.RGBA, rotation Rotation) {
	numColors := uint16(len(colors))
	if numColors == 0 {
		return
	}
	rotation = rotation % 4

	c := uint16(0)
	w, h := display.Size()
	l := len(text)
	for i := 0; i < l; i++ {
		glyph := font.Glyphs[text[i]-font.First]
		//if x+int16(glyph.XAdvance) >= 0 {
		DrawCharRotated(display, font, x, y, text[i], colors[c], rotation)
		//}
		c++
		if c >= numColors {
			c = 0
		}
		if rotation == 0 {
			x += int16(glyph.XAdvance)
		} else if rotation == 1 {
			y += int16(glyph.XAdvance)
		} else if rotation == 2 {
			x -= int16(glyph.XAdvance)
		} else {
			y -= int16(glyph.XAdvance)
		}

		// speed up?
		if (rotation == 0 && x > w) ||
			(rotation == 1 && x > h) ||
			(rotation == 2 && x < 0) ||
			(rotation == 3 && y < 0) {
			break
		}
	}
}

func LineWidth(font *Font, text []byte) (innerWidth uint32, outboxWidth uint32) {
	for i := range text {
		glyph := font.Glyphs[text[i]-font.First]
		outboxWidth += uint32(glyph.XAdvance)
	}
	innerWidth = outboxWidth
	// first glyph
	glyph := font.Glyphs[text[0]-font.First]
	innerWidth -= uint32(glyph.XOffset)

	// last glyph
	glyph = font.Glyphs[text[0]-font.First]
	innerWidth += -uint32(glyph.XAdvance) + uint32(glyph.XOffset) + uint32(glyph.Width)
	return
}
