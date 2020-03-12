package tinyfont

import (
	"fmt"
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
	Width    uint8
	Height   uint8
	XAdvance uint8
	XOffset  int8
	YOffset  int8
	Bitmaps  []byte
}

type RuneToIndex struct {
	Rune  rune
	Index uint16
}

type Font struct {
	Glyphs      []Glyph
	RuneToIndex []RuneToIndex
	YAdvance    uint8
}

// DrawChar sets a single char in the buffer of the display
func DrawChar(display drivers.Displayer, font *Font, x int16, y int16, char rune, color color.RGBA) {
	DrawCharRotated(display, font, x, y, char, color, 0)
}

// DrawCharRotated sets a single char in the buffer of the display
func DrawCharRotated(display drivers.Displayer, font *Font, x int16, y int16, char rune, color color.RGBA, rotation Rotation) {
	glyph, err := GetGlyph(font, char)
	if err != nil {
		// TODO
	}
	bitmapOffset := 0
	bitmap := glyph.Bitmaps[bitmapOffset]
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
				if bitmapOffset < len(glyph.Bitmaps) {
					bitmap = glyph.Bitmaps[bitmapOffset]
				}
				bit = 0
			}
		}
	}
}

// WriteLine writes a string in the selected font in the buffer
func WriteLine(display drivers.Displayer, font *Font, x int16, y int16, text []rune, color color.RGBA) {
	WriteLineRotated(display, font, x, y, text, color, 0)
}

// WriteLineRotated writes a string in the selected font in the buffer
func WriteLineRotated(display drivers.Displayer, font *Font, x int16, y int16, text []rune, color color.RGBA, rotation Rotation) {
	rotation = rotation % 4
	w, h := display.Size()
	l := len(text)
	ox := x
	oy := y
	for i := 0; i < l; i++ {
		if text[i] == 0x0A || text[i] == 0x0D {
			/* CR or LF */
			if rotation == 0 {
				x = ox
				y += int16(font.YAdvance) + 1
			} else if rotation == 1 {
				x -= int16(font.YAdvance) + 1
				y = oy
			} else if rotation == 2 {
				x = ox
				y -= int16(font.YAdvance) + 1
			} else {
				x += int16(font.YAdvance) + 1
				y = oy
			}
			continue
		}

		glyph, err := GetGlyph(font, text[i])
		if err != nil {
			// TODO
		}
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
func WriteLineColors(display drivers.Displayer, font *Font, x int16, y int16, text []rune, colors []color.RGBA) {
	WriteLineColorsRotated(display, font, x, y, text, colors, 0)
}

// WriteLineColorsRotated writes a string in the selected font in the buffer. Each char is in a different color
// if not enough colors are defined, colors are cycled.
func WriteLineColorsRotated(display drivers.Displayer, font *Font, x int16, y int16, text []rune, colors []color.RGBA, rotation Rotation) {
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
		if text[i] == 0x0A || text[i] == 0x0D {
			/* CR or LF */
			if rotation == 0 {
				x = ox
				y += int16(font.YAdvance) + 1
			} else if rotation == 1 {
				x -= int16(font.YAdvance) + 1
				y = oy
			} else if rotation == 2 {
				x = ox
				y -= int16(font.YAdvance) + 1
			} else {
				x += int16(font.YAdvance) + 1
				y = oy
			}
			continue
		}
		glyph, err := GetGlyph(font, text[i])
		if err != nil {
			// TODO
		}
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

func LineWidth(font *Font, text []rune) (innerWidth uint32, outboxWidth uint32) {
	for i := range text {
		glyph, err := GetGlyph(font, text[i])
		if err != nil {
			// TODO
		}
		outboxWidth += uint32(glyph.XAdvance)
	}
	innerWidth = outboxWidth
	// first glyph
	glyph, err := GetGlyph(font, text[0])
	if err != nil {
		// TODO
	}
	innerWidth -= uint32(glyph.XOffset)

	// last glyph
	glyph, err = GetGlyph(font, text[len(text)-1])
	if err != nil {
		// TODO
	}
	innerWidth += -uint32(glyph.XAdvance) + uint32(glyph.XOffset) + uint32(glyph.Width)
	return
}

func GetGlyph(font *Font, r rune) (Glyph, error) {
	s := 0
	e := len(font.RuneToIndex) - 1

	for s <= e {
		m := (s + e) / 2

		if font.RuneToIndex[m].Rune < r {
			s = m + 1
		} else {
			e = m - 1
		}
	}

	if s == len(font.RuneToIndex) || font.RuneToIndex[s].Rune != r {
		return Glyph{}, fmt.Errorf("glyph not found")
	}

	return font.Glyphs[s], nil
}
