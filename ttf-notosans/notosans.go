package notosans

import (
	"fmt"
	"image/color"

	"tinygo.org/x/drivers"
	"tinygo.org/x/tinyfont"
)

type NotosansGlyph struct {
	Index uint32
	Rune  rune
}

func (g NotosansGlyph) Draw(display drivers.Displayer, x int16, y int16, c1 color.RGBA) {
	info := g.Info()
	bmp := cNotoSans12pt[g.Index+10:]
	b := cNotoSans12pt[g.Index:]

	fmt.Printf("ok\n")
	fmt.Printf("%#v\n", info)
	fmt.Printf("%04X %08X\n", g.Index, uint32(b[0]))
	fmt.Printf("%02X %02X %02X %02X ", b[0], b[1], b[2], b[3])
	fmt.Printf("%02X %02X %02X %02X\n", b[4], b[5], b[6], b[7])
	fmt.Printf("%02X %02X %02X\n", info.Rune, info.Width, info.Height)
	for i := 0; i < int(b[0])+1; i++ {
		fmt.Printf("%02X ", b[i])
	}
	fmt.Printf("\n")

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

func (g NotosansGlyph) Info() *tinyfont.GlyphInfo {
	buf := cNotoSans12pt[g.Index:]
	return &tinyfont.GlyphInfo{
		Rune:     g.Rune,
		Width:    buf[5],
		Height:   buf[6],
		XAdvance: buf[7],
		XOffset:  int8(buf[8]),
		YOffset:  int8(buf[9]),
	}
}
