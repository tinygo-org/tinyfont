package notosans

import (
	"image/color"

	"tinygo.org/x/drivers"
	"tinygo.org/x/tinyfont"
)

type NotoSans12ptGlyph struct {
	Offset uint32
	Rune   rune
}

func (g NotoSans12ptGlyph) Draw(display drivers.Displayer, x int16, y int16, c1 color.RGBA) {
	info := g.Info()
	bmp := cNotoSans12pt[g.Offset+5:]

	for j := int16(0); j < int16(info.Height); j++ {
		for i := int16(0); i < int16(info.Width); i++ {
			//display.SetPixel(x+int16(info.XOffset)+i, y+int16(info.YOffset)+j, color.RGBA{0x00, 0x00, 0x00, 0xFF})

			shiftNum := 0
			idx := i + j*int16(info.Width)
			switch idx % 4 {
			case 0:
				shiftNum = 6
			case 1:
				shiftNum = 4
			case 2:
				shiftNum = 2
			case 3:
				shiftNum = 0
			}

			c := (bmp[idx/4] >> shiftNum) & 0x03

			//fmt.Printf("+%d+%d+%x ", idx, shiftNum, c)
			//fmt.Printf("%x ", c)
			//if i == int16(info.Width)-1 {
			//	fmt.Printf("\n")
			//}
			switch c {
			case 0:
				c = 0xFF
			case 1:
				c = 0xA0
			case 2:
				c = 0x60
			case 3:
				c = 0x00
			default:
				panic("err")
			}
			//c = 0xFF - c

			display.SetPixel(x+int16(info.XOffset)+i, y+int16(info.YOffset)+j, color.RGBA{c, c, c, 0xFF})
		}
	}
}

func (g NotoSans12ptGlyph) Info() *tinyfont.GlyphInfo {
	buf := cNotoSans12pt[g.Offset:]
	return &tinyfont.GlyphInfo{
		Rune:     g.Rune,
		Width:    buf[0],
		Height:   buf[1],
		XAdvance: buf[2],
		XOffset:  int8(buf[3]),
		YOffset:  int8(buf[4]),
	}
}
