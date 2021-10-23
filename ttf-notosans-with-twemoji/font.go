package nososans

import (
	"fmt"

	"tinygo.org/x/tinyfont"
	notosans "tinygo.org/x/tinyfont/ttf-notosans"
	"tinygo.org/x/tinyfont/twemoji"
)

type Font struct {
	BBox     [4]int8 // width, height, minX, minY
	Glyphs   []tinyfont.Glypher
	YAdvance uint8
}

var NotoSans12pt = Font{
	Glyphs:   []tinyfont.Glypher{},
	YAdvance: 12,
}

func (f Font) GetGlyph(r rune) tinyfont.Glypher {
	g := twemoji.Twemoji12pt.GetGlyph(r)
	if g != nil {
		fmt.Printf("%c", r)
		return g
	}

	g = notosans.NotoSans12pt.GetGlyph(r)
	if g != nil {
		return g
	}

	panic("not impl")
	return nil
}

func (f Font) GetYAdvance() uint8 {
	return f.YAdvance
}

func (f Font) LineWidth(str string) (innerWidth uint32, outboxWidth uint32) {
	return 0, 0
}
