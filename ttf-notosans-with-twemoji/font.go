package nososans

import (
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
	YAdvance: 14,
}

func (f Font) GetGlyph(r rune) tinyfont.Glypher {
	g := twemoji.Twemoji12pt.GetGlyph(r)
	if g != nil {
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
