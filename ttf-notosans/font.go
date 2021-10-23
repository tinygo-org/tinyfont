package notosans

import (
	"tinygo.org/x/tinyfont"
)

type Font struct {
	BBox     [4]int8 // width, height, minX, minY
	Glyphs   []tinyfont.Glypher
	YAdvance uint8
}

func (f Font) GetGlyph(r rune) tinyfont.Glypher {
	s := 0
	e := len(mNotoSans12pt)/6 - 1

	for s <= e {
		m := (s + e) / 2

		rr := rune(mNotoSans12pt[m*6])<<16 + rune(mNotoSans12pt[m*6+1])<<8 + rune(mNotoSans12pt[m*6+2])
		if rr < r {
			s = m + 1
		} else {
			e = m - 1
		}
	}

	if s == len(mNotoSans12pt)/6 {
		return nil
	}

	offset := uint32(mNotoSans12pt[s*6+3])<<16 + uint32(mNotoSans12pt[s*6+4])<<8 + uint32(mNotoSans12pt[s*6+5])
	g := &NotoSans12ptGlyph{
		Rune:   r,
		Offset: offset,
	}
	return g
}

func (f Font) GetYAdvance() uint8 {
	return f.YAdvance
}

func (f Font) LineWidth(str string) (innerWidth uint32, outboxWidth uint32) {
	return 0, 0
}
