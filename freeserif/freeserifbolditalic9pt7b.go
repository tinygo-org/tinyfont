package freeserif

import (
	"tinygo.org/x/tinyfont"
)

var BoldItalic9pt7b = tinyfont.Font{
	BBox: [4]int8{20, 17, -2, -12},
	Glyphs: []tinyfont.Glyph{
		/*   */ tinyfont.Glyph{Rune: 32, Width: 0x0, Height: 0x0, XAdvance: 0x5, XOffset: 0, YOffset: 1, Bitmaps: []uint8{}},
		/* ! */ tinyfont.Glyph{Rune: 33, Width: 0x6, Height: 0xd, XAdvance: 0x7, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xc, 0x31, 0xc6, 0x18, 0x41, 0x8, 0x20, 0xe, 0x38, 0xe0}},
		/* " */ tinyfont.Glyph{Rune: 34, Width: 0x6, Height: 0x5, XAdvance: 0xa, XOffset: 3, YOffset: -11, Bitmaps: []uint8{0xcf, 0x38, 0xa2, 0x88}},
		/* # */ tinyfont.Glyph{Rune: 35, Width: 0xb, Height: 0xd, XAdvance: 0x9, XOffset: -1, YOffset: -12, Bitmaps: []uint8{0x2, 0x40, 0xc8, 0x13, 0x6, 0x43, 0xfc, 0x32, 0x6, 0x40, 0x98, 0x7f, 0x84, 0xc0, 0x90, 0x32, 0x4, 0xc0}},
		/* $ */ tinyfont.Glyph{Rune: 36, Width: 0xb, Height: 0xf, XAdvance: 0x9, XOffset: -1, YOffset: -12, Bitmaps: []uint8{0x1, 0x1, 0xf0, 0x4b, 0x99, 0x33, 0x24, 0x78, 0x7, 0x80, 0x38, 0xb, 0x89, 0x31, 0x26, 0x64, 0xc7, 0x30, 0x3c, 0x4, 0x0}},
		/* % */ tinyfont.Glyph{Rune: 37, Width: 0xe, Height: 0xd, XAdvance: 0xf, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x38, 0x41, 0x9f, 0x6, 0x48, 0x31, 0x60, 0xcd, 0x3, 0x2c, 0x7, 0x27, 0x81, 0x39, 0x5, 0xc4, 0x26, 0x10, 0x98, 0x84, 0x66, 0x10, 0xe0}},
		/* & */ tinyfont.Glyph{Rune: 38, Width: 0xd, Height: 0xd, XAdvance: 0xe, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x3, 0x80, 0x22, 0x3, 0x10, 0x19, 0x0, 0xf0, 0xf, 0x3c, 0xf8, 0xcc, 0xc4, 0xe7, 0x47, 0x3e, 0x38, 0xe1, 0xe7, 0x97, 0xcf, 0x0}},
		/* ' */ tinyfont.Glyph{Rune: 39, Width: 0x2, Height: 0x5, XAdvance: 0x5, XOffset: 3, YOffset: -11, Bitmaps: []uint8{0xfa, 0x80}},
		/* ( */ tinyfont.Glyph{Rune: 40, Width: 0x5, Height: 0x10, XAdvance: 0x6, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x8, 0x88, 0x84, 0x62, 0x10, 0x84, 0x21, 0x8, 0x41, 0x0}},
		/* ) */ tinyfont.Glyph{Rune: 41, Width: 0x5, Height: 0x10, XAdvance: 0x6, XOffset: -1, YOffset: -11, Bitmaps: []uint8{0x20, 0x84, 0x10, 0x84, 0x21, 0x8, 0xc6, 0x23, 0x11, 0x0}},
		/* * */ tinyfont.Glyph{Rune: 42, Width: 0x8, Height: 0x8, XAdvance: 0x9, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x18, 0x18, 0xd6, 0x38, 0x18, 0xf7, 0x18, 0x18}},
		/* + */ tinyfont.Glyph{Rune: 43, Width: 0x9, Height: 0x9, XAdvance: 0xa, XOffset: 0, YOffset: -8, Bitmaps: []uint8{0x8, 0x4, 0x2, 0x1, 0xf, 0xf8, 0x40, 0x20, 0x10, 0x8, 0x0}},
		/* , */ tinyfont.Glyph{Rune: 44, Width: 0x3, Height: 0x6, XAdvance: 0x5, XOffset: -1, YOffset: -2, Bitmaps: []uint8{0x6d, 0x95, 0x0}},
		/* - */ tinyfont.Glyph{Rune: 45, Width: 0x5, Height: 0x2, XAdvance: 0x6, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xff, 0xc0}},
		/* . */ tinyfont.Glyph{Rune: 46, Width: 0x3, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -1, Bitmaps: []uint8{0xff, 0x80}},
		/* / */ tinyfont.Glyph{Rune: 47, Width: 0x7, Height: 0xc, XAdvance: 0x6, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x6, 0xc, 0x30, 0x60, 0x83, 0x4, 0x18, 0x20, 0xc1, 0x6, 0x0}},
		/* 0 */ tinyfont.Glyph{Rune: 48, Width: 0x9, Height: 0xd, XAdvance: 0x9, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xf, 0xc, 0x8c, 0x6e, 0x37, 0x1b, 0x1f, 0x8f, 0xc7, 0xc7, 0x63, 0xb1, 0x89, 0x83, 0x80}},
		/* 1 */ tinyfont.Glyph{Rune: 49, Width: 0x8, Height: 0xd, XAdvance: 0x9, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x6, 0x1e, 0xe, 0xe, 0xc, 0xc, 0x1c, 0x18, 0x18, 0x18, 0x38, 0x38, 0xfc}},
		/* 2 */ tinyfont.Glyph{Rune: 50, Width: 0x9, Height: 0xd, XAdvance: 0x9, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x1f, 0x13, 0xd0, 0xe0, 0x70, 0x38, 0x38, 0x18, 0x18, 0x18, 0x8, 0x8, 0x4f, 0xcf, 0xe0}},
		/* 3 */ tinyfont.Glyph{Rune: 51, Width: 0x9, Height: 0xd, XAdvance: 0x9, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x1f, 0x11, 0xc0, 0xe0, 0x60, 0xc1, 0xf0, 0x38, 0xc, 0x6, 0x3, 0x1, 0x19, 0x8f, 0x0}},
		/* 4 */ tinyfont.Glyph{Rune: 52, Width: 0x9, Height: 0xc, XAdvance: 0x9, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x0, 0x80, 0xc0, 0xe1, 0xe0, 0xb0, 0x98, 0x9c, 0x8c, 0xff, 0x7, 0x3, 0x1, 0x80}},
		/* 5 */ tinyfont.Glyph{Rune: 53, Width: 0x9, Height: 0xd, XAdvance: 0x9, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xf, 0x88, 0x8, 0x7, 0x83, 0xe0, 0x78, 0x1c, 0x6, 0x3, 0x1, 0x80, 0x9c, 0x87, 0x80}},
		/* 6 */ tinyfont.Glyph{Rune: 54, Width: 0x9, Height: 0xd, XAdvance: 0x9, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x3, 0x87, 0x7, 0x7, 0x7, 0x3, 0xe3, 0x99, 0xcc, 0xc6, 0x63, 0x33, 0x89, 0x87, 0x80}},
		/* 7 */ tinyfont.Glyph{Rune: 55, Width: 0x9, Height: 0xc, XAdvance: 0x9, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x3f, 0xbf, 0x90, 0x80, 0xc0, 0x40, 0x60, 0x20, 0x30, 0x30, 0x10, 0x18, 0x8, 0x0}},
		/* 8 */ tinyfont.Glyph{Rune: 56, Width: 0x8, Height: 0xd, XAdvance: 0x9, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x1e, 0x13, 0x31, 0x31, 0x3a, 0x1c, 0x1c, 0x6e, 0xc6, 0xc6, 0xc6, 0x44, 0x38}},
		/* 9 */ tinyfont.Glyph{Rune: 57, Width: 0x9, Height: 0xd, XAdvance: 0x9, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xe, 0x1c, 0x8c, 0x6c, 0x36, 0x3b, 0x1d, 0x8e, 0x7e, 0xe, 0x7, 0x7, 0xe, 0xc, 0x0}},
		/* : */ tinyfont.Glyph{Rune: 58, Width: 0x5, Height: 0x9, XAdvance: 0x5, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x39, 0xce, 0x0, 0x3, 0x9c, 0xe0}},
		/* ; */ tinyfont.Glyph{Rune: 59, Width: 0x5, Height: 0xb, XAdvance: 0x5, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x39, 0xce, 0x0, 0x1, 0x8c, 0x22, 0x20}},
		/* < */ tinyfont.Glyph{Rune: 60, Width: 0x9, Height: 0xa, XAdvance: 0xa, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0x0, 0x1, 0xc3, 0xc7, 0x8e, 0x6, 0x1, 0xe0, 0x3c, 0x7, 0x80, 0x40}},
		/* = */ tinyfont.Glyph{Rune: 61, Width: 0x9, Height: 0x5, XAdvance: 0xa, XOffset: 1, YOffset: -6, Bitmaps: []uint8{0xff, 0x80, 0x0, 0x0, 0xf, 0xf8}},
		/* > */ tinyfont.Glyph{Rune: 62, Width: 0x9, Height: 0xa, XAdvance: 0xa, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0x0, 0x60, 0x1e, 0x3, 0xc0, 0x78, 0x1c, 0x3c, 0x78, 0xf0, 0x40, 0x0}},
		/* ? */ tinyfont.Glyph{Rune: 63, Width: 0x8, Height: 0xd, XAdvance: 0x9, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x1c, 0x27, 0x37, 0x7, 0xe, 0x1c, 0x30, 0x60, 0x40, 0x0, 0xe0, 0xe0, 0xe0}},
		/* @ */ tinyfont.Glyph{Rune: 64, Width: 0xd, Height: 0xd, XAdvance: 0xf, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0xf, 0x80, 0xc3, 0x8, 0x4, 0xc3, 0x3c, 0x24, 0xe2, 0x27, 0x33, 0x39, 0x11, 0xc9, 0x93, 0x77, 0x18, 0x0, 0x70, 0x40, 0xfc, 0x0}},
		/* A */ tinyfont.Glyph{Rune: 65, Width: 0xc, Height: 0xd, XAdvance: 0xd, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x0, 0x80, 0x18, 0x1, 0x80, 0x38, 0x5, 0x80, 0x5c, 0x9, 0xc1, 0x1c, 0x1f, 0xc2, 0xc, 0x20, 0xc4, 0xe, 0xf3, 0xf0}},
		/* B */ tinyfont.Glyph{Rune: 66, Width: 0xc, Height: 0xd, XAdvance: 0xc, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x3f, 0xe0, 0xc7, 0xc, 0x71, 0xc7, 0x1c, 0xe1, 0xf0, 0x39, 0xc3, 0x8e, 0x38, 0xe3, 0xe, 0x71, 0xe7, 0x1c, 0xff, 0x0}},
		/* C */ tinyfont.Glyph{Rune: 67, Width: 0xc, Height: 0xd, XAdvance: 0xb, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x7, 0xd1, 0xc7, 0x38, 0x27, 0x2, 0x70, 0xf, 0x0, 0xe0, 0xe, 0x0, 0xe0, 0xe, 0x0, 0x60, 0x87, 0x18, 0x1e, 0x0}},
		/* D */ tinyfont.Glyph{Rune: 68, Width: 0xe, Height: 0xd, XAdvance: 0xd, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x3f, 0xe0, 0x30, 0xe0, 0xc1, 0x87, 0x7, 0x1c, 0x1c, 0x60, 0x73, 0x81, 0xce, 0x7, 0x38, 0x38, 0xc0, 0xe7, 0x7, 0x1c, 0x78, 0xff, 0x80}},
		/* E */ tinyfont.Glyph{Rune: 69, Width: 0xd, Height: 0xd, XAdvance: 0xb, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x1f, 0xf8, 0x61, 0xc3, 0x4, 0x38, 0x81, 0xcc, 0xf, 0xe0, 0xe2, 0x7, 0x10, 0x38, 0x81, 0x81, 0x1c, 0x18, 0xe3, 0x8f, 0xfc, 0x0}},
		/* F */ tinyfont.Glyph{Rune: 70, Width: 0xd, Height: 0xd, XAdvance: 0xb, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x3f, 0xf8, 0x61, 0xc3, 0x4, 0x38, 0x81, 0xcc, 0xf, 0xe0, 0xe2, 0x7, 0x10, 0x38, 0x81, 0x80, 0x1c, 0x0, 0xe0, 0xf, 0x80, 0x0}},
		/* G */ tinyfont.Glyph{Rune: 71, Width: 0xc, Height: 0xd, XAdvance: 0xd, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x7, 0x91, 0xc7, 0x38, 0x27, 0x0, 0x70, 0xf, 0x0, 0xe1, 0xfe, 0xe, 0xe0, 0xce, 0xc, 0x60, 0xc7, 0x1c, 0x1f, 0x0}},
		/* H */ tinyfont.Glyph{Rune: 72, Width: 0xf, Height: 0xd, XAdvance: 0xe, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x1f, 0x7e, 0x1c, 0x38, 0x30, 0x60, 0xe1, 0xc1, 0xc3, 0x83, 0x6, 0xf, 0xfc, 0x1c, 0x38, 0x38, 0x70, 0x60, 0xc1, 0xc3, 0x83, 0x87, 0xf, 0x9f, 0x0}},
		/* I */ tinyfont.Glyph{Rune: 73, Width: 0x8, Height: 0xd, XAdvance: 0x7, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x3f, 0xc, 0xc, 0x1c, 0x1c, 0x18, 0x38, 0x38, 0x38, 0x30, 0x70, 0x70, 0xf8}},
		/* J */ tinyfont.Glyph{Rune: 74, Width: 0xa, Height: 0xe, XAdvance: 0x9, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x7, 0xc0, 0xe0, 0x38, 0xc, 0x7, 0x1, 0xc0, 0x70, 0x18, 0xe, 0x3, 0x80, 0xc3, 0x30, 0xdc, 0x1e, 0x0}},
		/* K */ tinyfont.Glyph{Rune: 75, Width: 0xd, Height: 0xd, XAdvance: 0xc, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x1f, 0x78, 0x71, 0x83, 0x18, 0x39, 0x81, 0xd0, 0xd, 0x0, 0xfc, 0x7, 0x60, 0x3b, 0x81, 0x8c, 0x1c, 0x70, 0xe1, 0x8f, 0xbe, 0x0}},
		/* L */ tinyfont.Glyph{Rune: 76, Width: 0xc, Height: 0xd, XAdvance: 0xb, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x1f, 0x0, 0xc0, 0xc, 0x1, 0xc0, 0x1c, 0x1, 0x80, 0x38, 0x3, 0x80, 0x38, 0x3, 0x1, 0x70, 0x37, 0xe, 0xff, 0xe0}},
		/* M */ tinyfont.Glyph{Rune: 77, Width: 0x11, Height: 0xd, XAdvance: 0x10, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x1e, 0x7, 0x87, 0x7, 0x83, 0x83, 0x82, 0xc3, 0xc1, 0x62, 0xe0, 0xb1, 0x70, 0x99, 0x30, 0x4d, 0xb8, 0x27, 0x9c, 0x13, 0x8c, 0x11, 0xc6, 0xc, 0xc7, 0xf, 0x47, 0xc0}},
		/* N */ tinyfont.Glyph{Rune: 78, Width: 0xe, Height: 0xd, XAdvance: 0xd, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x3c, 0x3c, 0x38, 0x20, 0xe0, 0x85, 0xc4, 0x13, 0x10, 0x4e, 0x42, 0x3a, 0x8, 0x78, 0x21, 0xe0, 0x83, 0x84, 0xc, 0x18, 0x10, 0x0, 0x40}},
		/* O */ tinyfont.Glyph{Rune: 79, Width: 0xc, Height: 0xd, XAdvance: 0xc, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x7, 0xc1, 0xce, 0x38, 0x73, 0x87, 0x70, 0x77, 0x7, 0xf0, 0xfe, 0xe, 0xe0, 0xee, 0x1c, 0xe1, 0xc6, 0x38, 0x3e, 0x0}},
		/* P */ tinyfont.Glyph{Rune: 80, Width: 0xc, Height: 0xd, XAdvance: 0xb, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x3f, 0xc0, 0xc7, 0xc, 0x71, 0xc7, 0x1c, 0x71, 0x8e, 0x3f, 0xc3, 0x80, 0x30, 0x3, 0x0, 0x70, 0x7, 0x0, 0xf8, 0x0}},
		/* Q */ tinyfont.Glyph{Rune: 81, Width: 0xc, Height: 0x10, XAdvance: 0xc, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x7, 0xc0, 0xce, 0x38, 0x73, 0x87, 0x70, 0x77, 0x7, 0xf0, 0x7e, 0xe, 0xe0, 0xee, 0xc, 0xe1, 0xc6, 0x38, 0x36, 0x1, 0x80, 0x3c, 0x2d, 0xfc}},
		/* R */ tinyfont.Glyph{Rune: 82, Width: 0xc, Height: 0xd, XAdvance: 0xc, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x3f, 0xc0, 0xe7, 0xc, 0x71, 0xc7, 0x1c, 0x71, 0x8e, 0x3f, 0x83, 0xb8, 0x3b, 0x83, 0x3c, 0x71, 0xc7, 0x1c, 0xf9, 0xf0}},
		/* S */ tinyfont.Glyph{Rune: 83, Width: 0x9, Height: 0xd, XAdvance: 0x9, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xc, 0x89, 0x8c, 0x46, 0x23, 0x80, 0xe0, 0x78, 0xe, 0x3, 0x21, 0x90, 0xcc, 0xc9, 0xc0}},
		/* T */ tinyfont.Glyph{Rune: 84, Width: 0xb, Height: 0xd, XAdvance: 0xb, XOffset: 2, YOffset: -11, Bitmaps: []uint8{0x7f, 0xe9, 0xdf, 0x31, 0x4e, 0x21, 0xc0, 0x38, 0x6, 0x1, 0xc0, 0x38, 0x6, 0x0, 0xc0, 0x38, 0xf, 0xc0}},
		/* U */ tinyfont.Glyph{Rune: 85, Width: 0xc, Height: 0xd, XAdvance: 0xd, XOffset: 2, YOffset: -11, Bitmaps: []uint8{0x7c, 0xf3, 0x82, 0x30, 0x27, 0x4, 0x70, 0x46, 0x4, 0xe0, 0x4e, 0x8, 0xe0, 0x8e, 0x8, 0xe1, 0xf, 0x30, 0x3c, 0x0}},
		/* V */ tinyfont.Glyph{Rune: 86, Width: 0xc, Height: 0xc, XAdvance: 0xd, XOffset: 2, YOffset: -11, Bitmaps: []uint8{0xfc, 0x73, 0x82, 0x38, 0x23, 0x84, 0x38, 0x83, 0x90, 0x39, 0x1, 0xa0, 0x1c, 0x1, 0xc0, 0x18, 0x1, 0x0}},
		/* W */ tinyfont.Glyph{Rune: 87, Width: 0x10, Height: 0xc, XAdvance: 0x11, XOffset: 2, YOffset: -11, Bitmaps: []uint8{0xf9, 0xf7, 0x30, 0xe2, 0x30, 0xc2, 0x38, 0xc4, 0x3b, 0xc4, 0x3a, 0xe8, 0x3c, 0xe8, 0x3c, 0xf0, 0x18, 0xf0, 0x18, 0x60, 0x10, 0x60, 0x10, 0x40}},
		/* X */ tinyfont.Glyph{Rune: 88, Width: 0xd, Height: 0xd, XAdvance: 0xd, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x3f, 0x78, 0x61, 0x83, 0x98, 0x1d, 0x0, 0x70, 0x3, 0x80, 0x1c, 0x1, 0x60, 0xb, 0x80, 0x9c, 0x8, 0x60, 0xc3, 0x8f, 0x7e, 0x0}},
		/* Y */ tinyfont.Glyph{Rune: 89, Width: 0xb, Height: 0xd, XAdvance: 0xb, XOffset: 2, YOffset: -11, Bitmaps: []uint8{0xf9, 0xe6, 0x18, 0xc2, 0x1c, 0x81, 0xa0, 0x34, 0x7, 0x0, 0xc0, 0x18, 0x7, 0x0, 0xe0, 0x1c, 0xf, 0xc0}},
		/* Z */ tinyfont.Glyph{Rune: 90, Width: 0xb, Height: 0xd, XAdvance: 0xa, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x3f, 0xe6, 0x19, 0x87, 0x21, 0xc0, 0x30, 0xe, 0x3, 0x80, 0x60, 0x1c, 0x7, 0x5, 0xc1, 0x38, 0xef, 0xfc}},
		/* [ */ tinyfont.Glyph{Rune: 91, Width: 0x8, Height: 0xf, XAdvance: 0x6, XOffset: -1, YOffset: -11, Bitmaps: []uint8{0xe, 0x8, 0x18, 0x18, 0x18, 0x10, 0x30, 0x30, 0x30, 0x20, 0x60, 0x60, 0x60, 0x40, 0xf0}},
		/* \ */ tinyfont.Glyph{Rune: 92, Width: 0x5, Height: 0xc, XAdvance: 0x7, XOffset: 2, YOffset: -11, Bitmaps: []uint8{0xc6, 0x10, 0xc6, 0x10, 0x86, 0x30, 0x86, 0x30}},
		/* ] */ tinyfont.Glyph{Rune: 93, Width: 0x7, Height: 0xf, XAdvance: 0x6, XOffset: -1, YOffset: -11, Bitmaps: []uint8{0x1e, 0xc, 0x18, 0x20, 0xc1, 0x83, 0x4, 0x18, 0x30, 0x60, 0x83, 0x6, 0x3c, 0x0}},
		/* ^ */ tinyfont.Glyph{Rune: 94, Width: 0x8, Height: 0x7, XAdvance: 0xa, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x18, 0x1c, 0x34, 0x26, 0x66, 0x43, 0xc3}},
		/* _ */ tinyfont.Glyph{Rune: 95, Width: 0x9, Height: 0x1, XAdvance: 0x9, XOffset: 0, YOffset: 3, Bitmaps: []uint8{0xff, 0x80}},
		/* ` */ tinyfont.Glyph{Rune: 96, Width: 0x4, Height: 0x3, XAdvance: 0x6, XOffset: 2, YOffset: -11, Bitmaps: []uint8{0xc6, 0x30}},
		/* a */ tinyfont.Glyph{Rune: 97, Width: 0x9, Height: 0x9, XAdvance: 0x9, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xd, 0x9d, 0x8c, 0xcc, 0x6e, 0x26, 0x33, 0x19, 0xbe, 0x66, 0x0}},
		/* b */ tinyfont.Glyph{Rune: 98, Width: 0x8, Height: 0xe, XAdvance: 0x9, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0x0, 0x78, 0x18, 0x30, 0x30, 0x3e, 0x73, 0x63, 0x63, 0x63, 0xc6, 0xc6, 0xcc, 0x70}},
		/* c */ tinyfont.Glyph{Rune: 99, Width: 0x8, Height: 0x9, XAdvance: 0x8, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xf, 0x3b, 0x70, 0x70, 0xe0, 0xe0, 0xe2, 0xe4, 0x78}},
		/* d */ tinyfont.Glyph{Rune: 100, Width: 0xa, Height: 0xe, XAdvance: 0x9, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0x0, 0x0, 0xf0, 0x1c, 0x6, 0x1, 0x83, 0xe3, 0x30, 0xcc, 0x63, 0x19, 0xcc, 0x63, 0x38, 0xcf, 0x1d, 0x80}},
		/* e */ tinyfont.Glyph{Rune: 101, Width: 0x7, Height: 0x9, XAdvance: 0x7, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xe, 0x75, 0xcb, 0xbe, 0xde, 0x38, 0x72, 0x78}},
		/* f */ tinyfont.Glyph{Rune: 102, Width: 0xb, Height: 0x11, XAdvance: 0x9, XOffset: -2, YOffset: -12, Bitmaps: []uint8{0x0, 0xe0, 0x34, 0xc, 0x1, 0x80, 0x30, 0x1f, 0x1, 0x80, 0x30, 0x6, 0x1, 0xc0, 0x30, 0x6, 0x0, 0xc0, 0x30, 0x6, 0x4, 0x80, 0xe0, 0x0}},
		/* g */ tinyfont.Glyph{Rune: 103, Width: 0x9, Height: 0xc, XAdvance: 0x9, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x1c, 0x19, 0xd8, 0xcc, 0x66, 0x60, 0xe1, 0x80, 0xf0, 0x7e, 0x43, 0x21, 0x8f, 0x0}},
		/* h */ tinyfont.Glyph{Rune: 104, Width: 0x9, Height: 0xe, XAdvance: 0xa, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0x0, 0x1e, 0x7, 0x3, 0x1, 0x80, 0xd8, 0xfc, 0x76, 0x33, 0x19, 0x99, 0xcc, 0xd6, 0x77, 0x30}},
		/* i */ tinyfont.Glyph{Rune: 105, Width: 0x5, Height: 0xd, XAdvance: 0x5, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x39, 0xc0, 0xf, 0x31, 0x8c, 0xc6, 0x31, 0xae, 0x0}},
		/* j */ tinyfont.Glyph{Rune: 106, Width: 0x9, Height: 0x10, XAdvance: 0x6, XOffset: -1, YOffset: -11, Bitmaps: []uint8{0x3, 0x81, 0xc0, 0x0, 0x0, 0xe0, 0x30, 0x18, 0x18, 0xc, 0x6, 0x3, 0x3, 0x1, 0x80, 0xc2, 0xc1, 0xc0}},
		/* k */ tinyfont.Glyph{Rune: 107, Width: 0xa, Height: 0xe, XAdvance: 0x9, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0x0, 0xf, 0x0, 0xc0, 0x60, 0x18, 0x6, 0xf3, 0x90, 0xc8, 0x34, 0xf, 0x6, 0xc1, 0x98, 0x66, 0xb9, 0xc0}},
		/* l */ tinyfont.Glyph{Rune: 108, Width: 0x5, Height: 0xe, XAdvance: 0x5, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x3, 0xcc, 0x63, 0x39, 0x8c, 0x66, 0x31, 0x8e, 0x70}},
		/* m */ tinyfont.Glyph{Rune: 109, Width: 0xd, Height: 0x9, XAdvance: 0xe, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x7b, 0x99, 0xaf, 0xce, 0x66, 0x63, 0x67, 0x33, 0x31, 0x99, 0x8c, 0xcc, 0xe7, 0xc6, 0x30}},
		/* n */ tinyfont.Glyph{Rune: 110, Width: 0x8, Height: 0x9, XAdvance: 0x9, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x73, 0x7f, 0x73, 0x73, 0x63, 0x67, 0xe6, 0xc7, 0xc6}},
		/* o */ tinyfont.Glyph{Rune: 111, Width: 0x8, Height: 0x9, XAdvance: 0x9, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x1e, 0x33, 0x63, 0x63, 0xc3, 0xc6, 0xc6, 0xcc, 0x78}},
		/* p */ tinyfont.Glyph{Rune: 112, Width: 0xa, Height: 0xc, XAdvance: 0x9, XOffset: -2, YOffset: -7, Bitmaps: []uint8{0x1d, 0xc3, 0xb1, 0xcc, 0x63, 0x19, 0xce, 0x63, 0x18, 0xcc, 0x3e, 0x1c, 0x6, 0x3, 0xe0}},
		/* q */ tinyfont.Glyph{Rune: 113, Width: 0x9, Height: 0xc, XAdvance: 0x9, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xd, 0x99, 0x8c, 0xcc, 0x6e, 0x76, 0x33, 0x19, 0x9c, 0x7c, 0x6, 0x7, 0x7, 0xc0}},
		/* r */ tinyfont.Glyph{Rune: 114, Width: 0x8, Height: 0x8, XAdvance: 0x7, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x76, 0x3a, 0x30, 0x70, 0x60, 0x60, 0x60, 0xe0}},
		/* s */ tinyfont.Glyph{Rune: 115, Width: 0x6, Height: 0x9, XAdvance: 0x6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x3d, 0x14, 0x58, 0x38, 0x60, 0xa2, 0xf0}},
		/* t */ tinyfont.Glyph{Rune: 116, Width: 0x5, Height: 0xc, XAdvance: 0x5, XOffset: 1, YOffset: -10, Bitmaps: []uint8{0x8, 0xcc, 0xf6, 0x31, 0x98, 0xc6, 0x35, 0xc0}},
		/* u */ tinyfont.Glyph{Rune: 117, Width: 0x8, Height: 0x9, XAdvance: 0xa, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0xe3, 0x63, 0x66, 0x66, 0x66, 0xcc, 0xcc, 0xfe, 0xec}},
		/* v */ tinyfont.Glyph{Rune: 118, Width: 0x7, Height: 0x8, XAdvance: 0x8, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0xe6, 0xcd, 0x8b, 0x26, 0x8e, 0x18, 0x20}},
		/* w */ tinyfont.Glyph{Rune: 119, Width: 0xa, Height: 0x8, XAdvance: 0xc, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0xe4, 0xd9, 0x36, 0xe5, 0xda, 0x77, 0x19, 0xc6, 0x61, 0x10}},
		/* x */ tinyfont.Glyph{Rune: 120, Width: 0xa, Height: 0x9, XAdvance: 0x9, XOffset: -1, YOffset: -7, Bitmaps: []uint8{0x39, 0xc7, 0xb0, 0xc0, 0x30, 0xc, 0x3, 0x0, 0xe1, 0x5a, 0x67, 0x0}},
		/* y */ tinyfont.Glyph{Rune: 121, Width: 0x9, Height: 0xc, XAdvance: 0x8, XOffset: -1, YOffset: -7, Bitmaps: []uint8{0x39, 0x8c, 0xc3, 0x21, 0xa0, 0xd0, 0x68, 0x38, 0xc, 0x4, 0x4, 0x14, 0xc, 0x0}},
		/* z */ tinyfont.Glyph{Rune: 122, Width: 0x8, Height: 0x9, XAdvance: 0x7, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x3e, 0x46, 0xc, 0x8, 0x10, 0x20, 0x70, 0x1a, 0xe}},
		/* { */ tinyfont.Glyph{Rune: 123, Width: 0x8, Height: 0x10, XAdvance: 0x6, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0x3, 0xe, 0xc, 0xc, 0x8, 0x18, 0x18, 0x10, 0x60, 0x30, 0x30, 0x30, 0x60, 0x60, 0x60, 0x30}},
		/* | */ tinyfont.Glyph{Rune: 124, Width: 0x1, Height: 0xc, XAdvance: 0x5, XOffset: 2, YOffset: -11, Bitmaps: []uint8{0xff, 0xf0}},
		/* } */ tinyfont.Glyph{Rune: 125, Width: 0x8, Height: 0x10, XAdvance: 0x6, XOffset: -2, YOffset: -12, Bitmaps: []uint8{0xc, 0x6, 0x6, 0x6, 0x4, 0xc, 0xc, 0xc, 0x6, 0x18, 0x18, 0x18, 0x30, 0x30, 0x30, 0xe0}},
		/* ~ */ tinyfont.Glyph{Rune: 126, Width: 0x8, Height: 0x2, XAdvance: 0xa, XOffset: 1, YOffset: -4, Bitmaps: []uint8{0x71, 0x8f}},
	},

	YAdvance: 0x16,
}
