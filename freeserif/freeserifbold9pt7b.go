package freeserif

import (
	"tinygo.org/x/tinyfont"
)

var Bold9pt7b = tinyfont.Font{
	BBox: [4]int8{18, 17, 0, -12},
	Glyphs: []tinyfont.Glyph{
		/*   */ tinyfont.Glyph{Rune: 32, Width: 0x0, Height: 0x0, XAdvance: 0x5, XOffset: 0, YOffset: 1, Bitmaps: []uint8{}},
		/* ! */ tinyfont.Glyph{Rune: 33, Width: 0x3, Height: 0xc, XAdvance: 0x6, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xff, 0xf4, 0x92, 0x1f, 0xf0}},
		/* " */ tinyfont.Glyph{Rune: 34, Width: 0x6, Height: 0x5, XAdvance: 0xa, XOffset: 2, YOffset: -11, Bitmaps: []uint8{0xcf, 0x3c, 0xe3, 0x88}},
		/* # */ tinyfont.Glyph{Rune: 35, Width: 0x9, Height: 0xd, XAdvance: 0x9, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0x13, 0x9, 0x84, 0xc2, 0x47, 0xf9, 0x90, 0xc8, 0x4c, 0xff, 0x13, 0x9, 0xc, 0x86, 0x40}},
		/* $ */ tinyfont.Glyph{Rune: 36, Width: 0x8, Height: 0xe, XAdvance: 0x9, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x10, 0x38, 0xd6, 0x92, 0xd2, 0xf0, 0x7c, 0x3e, 0x17, 0x93, 0x93, 0xd6, 0x7c, 0x10}},
		/* % */ tinyfont.Glyph{Rune: 37, Width: 0xe, Height: 0xc, XAdvance: 0x12, XOffset: 2, YOffset: -11, Bitmaps: []uint8{0x3c, 0x21, 0xcf, 0xe, 0x24, 0x30, 0xa0, 0xc5, 0x3, 0x34, 0xe7, 0x26, 0x40, 0xb9, 0x4, 0xc4, 0x23, 0x30, 0x8c, 0x84, 0x1c}},
		/* & */ tinyfont.Glyph{Rune: 38, Width: 0xd, Height: 0xc, XAdvance: 0xf, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xf, 0x0, 0xcc, 0x6, 0x60, 0x3e, 0x0, 0xe7, 0x8f, 0x18, 0x9c, 0x8c, 0xe4, 0xe3, 0xc7, 0x9e, 0x3c, 0x72, 0xfd, 0xe0}},
		/* ' */ tinyfont.Glyph{Rune: 39, Width: 0x2, Height: 0x5, XAdvance: 0x5, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xff, 0x80}},
		/* ( */ tinyfont.Glyph{Rune: 40, Width: 0x4, Height: 0xf, XAdvance: 0x6, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x32, 0x44, 0xcc, 0xcc, 0xcc, 0xc4, 0x62, 0x10}},
		/* ) */ tinyfont.Glyph{Rune: 41, Width: 0x4, Height: 0xf, XAdvance: 0x6, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x84, 0x22, 0x33, 0x33, 0x33, 0x32, 0x64, 0x80}},
		/* * */ tinyfont.Glyph{Rune: 42, Width: 0x7, Height: 0x7, XAdvance: 0x9, XOffset: 2, YOffset: -11, Bitmaps: []uint8{0x31, 0x6b, 0xb1, 0x8e, 0xd6, 0x8c, 0x0}},
		/* + */ tinyfont.Glyph{Rune: 43, Width: 0x9, Height: 0x9, XAdvance: 0xc, XOffset: 1, YOffset: -8, Bitmaps: []uint8{0x8, 0x4, 0x2, 0x1, 0xf, 0xf8, 0x40, 0x20, 0x10, 0x8, 0x0}},
		/* , */ tinyfont.Glyph{Rune: 44, Width: 0x3, Height: 0x6, XAdvance: 0x4, XOffset: 1, YOffset: -2, Bitmaps: []uint8{0xdf, 0x95, 0x0}},
		/* - */ tinyfont.Glyph{Rune: 45, Width: 0x4, Height: 0x2, XAdvance: 0x6, XOffset: 1, YOffset: -4, Bitmaps: []uint8{0xff}},
		/* . */ tinyfont.Glyph{Rune: 46, Width: 0x3, Height: 0x3, XAdvance: 0x4, XOffset: 1, YOffset: -2, Bitmaps: []uint8{0xff, 0x80}},
		/* / */ tinyfont.Glyph{Rune: 47, Width: 0x6, Height: 0xd, XAdvance: 0x5, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xc, 0x21, 0x86, 0x10, 0xc3, 0x8, 0x61, 0x84, 0x30, 0xc0}},
		/* 0 */ tinyfont.Glyph{Rune: 48, Width: 0x9, Height: 0xc, XAdvance: 0x9, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x1c, 0x33, 0x98, 0xdc, 0x7e, 0x3f, 0x1f, 0x8f, 0xc7, 0xe3, 0xb1, 0x98, 0xc3, 0x80}},
		/* 1 */ tinyfont.Glyph{Rune: 49, Width: 0x6, Height: 0xc, XAdvance: 0x9, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x8, 0xe3, 0x8e, 0x38, 0xe3, 0x8e, 0x38, 0xe3, 0xbf}},
		/* 2 */ tinyfont.Glyph{Rune: 50, Width: 0x9, Height: 0xc, XAdvance: 0x9, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x3c, 0x3f, 0x23, 0xc0, 0xe0, 0x70, 0x30, 0x38, 0x18, 0x18, 0x18, 0x5f, 0xdf, 0xe0}},
		/* 3 */ tinyfont.Glyph{Rune: 51, Width: 0x8, Height: 0xc, XAdvance: 0x9, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x7c, 0x8e, 0xe, 0xe, 0xc, 0x1e, 0x7, 0x3, 0x3, 0x2, 0xe6, 0xf8}},
		/* 4 */ tinyfont.Glyph{Rune: 52, Width: 0x8, Height: 0xc, XAdvance: 0x9, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x6, 0xe, 0xe, 0x3e, 0x2e, 0x4e, 0x8e, 0x8e, 0xff, 0xff, 0xe, 0xe}},
		/* 5 */ tinyfont.Glyph{Rune: 53, Width: 0x8, Height: 0xc, XAdvance: 0x9, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x3f, 0x7e, 0x40, 0x40, 0xf8, 0xfc, 0x1e, 0x6, 0x2, 0x2, 0xe4, 0xf8}},
		/* 6 */ tinyfont.Glyph{Rune: 54, Width: 0x8, Height: 0xc, XAdvance: 0x9, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x7, 0x1c, 0x30, 0x70, 0xfc, 0xe6, 0xe7, 0xe7, 0xe7, 0x67, 0x66, 0x3c}},
		/* 7 */ tinyfont.Glyph{Rune: 55, Width: 0x9, Height: 0xc, XAdvance: 0x9, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x7f, 0x3f, 0xa0, 0xd0, 0x40, 0x60, 0x30, 0x10, 0x18, 0xc, 0x4, 0x6, 0x3, 0x0}},
		/* 8 */ tinyfont.Glyph{Rune: 56, Width: 0x8, Height: 0xc, XAdvance: 0x9, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x3c, 0xc6, 0xc6, 0xc6, 0xfc, 0x7c, 0x3e, 0xcf, 0xc7, 0xc7, 0xc6, 0x7c}},
		/* 9 */ tinyfont.Glyph{Rune: 57, Width: 0x9, Height: 0xc, XAdvance: 0x9, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x3e, 0x33, 0xb8, 0xdc, 0x7e, 0x3f, 0x1d, 0xce, 0x7f, 0x7, 0x7, 0xf, 0x1c, 0x0}},
		/* : */ tinyfont.Glyph{Rune: 58, Width: 0x3, Height: 0x9, XAdvance: 0x6, XOffset: 1, YOffset: -8, Bitmaps: []uint8{0xff, 0x80, 0x3f, 0xe0}},
		/* ; */ tinyfont.Glyph{Rune: 59, Width: 0x3, Height: 0xc, XAdvance: 0x6, XOffset: 2, YOffset: -8, Bitmaps: []uint8{0xff, 0x80, 0x37, 0xe5, 0x40}},
		/* < */ tinyfont.Glyph{Rune: 60, Width: 0xa, Height: 0xa, XAdvance: 0xc, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0x0, 0x0, 0x70, 0x78, 0x78, 0x78, 0x38, 0x3, 0x80, 0x3c, 0x3, 0xc0, 0x30}},
		/* = */ tinyfont.Glyph{Rune: 61, Width: 0xa, Height: 0x5, XAdvance: 0xc, XOffset: 1, YOffset: -6, Bitmaps: []uint8{0xff, 0xc0, 0x0, 0x0, 0x0, 0xff, 0xc0}},
		/* > */ tinyfont.Glyph{Rune: 62, Width: 0xa, Height: 0xa, XAdvance: 0xc, XOffset: 1, YOffset: -8, Bitmaps: []uint8{0xc0, 0x3c, 0x3, 0xc0, 0x1c, 0x1, 0xc1, 0xe1, 0xe1, 0xe0, 0xe0, 0x0, 0x0}},
		/* ? */ tinyfont.Glyph{Rune: 63, Width: 0x7, Height: 0xc, XAdvance: 0x9, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x3d, 0x9f, 0x3e, 0x70, 0xe1, 0x4, 0x8, 0x0, 0x70, 0xe1, 0xc0}},
		/* @ */ tinyfont.Glyph{Rune: 64, Width: 0xd, Height: 0xc, XAdvance: 0x11, XOffset: 2, YOffset: -11, Bitmaps: []uint8{0xf, 0x81, 0x83, 0x18, 0xc4, 0x89, 0x9c, 0x4c, 0xe4, 0x67, 0x22, 0x39, 0x22, 0x4f, 0xe3, 0x0, 0xc, 0x10, 0x1f, 0x0}},
		/* A */ tinyfont.Glyph{Rune: 65, Width: 0xd, Height: 0xc, XAdvance: 0xd, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x2, 0x0, 0x30, 0x1, 0xc0, 0xe, 0x0, 0xb8, 0x5, 0xc0, 0x4f, 0x2, 0x38, 0x3f, 0xe1, 0x7, 0x18, 0x3d, 0xe3, 0xf0}},
		/* B */ tinyfont.Glyph{Rune: 66, Width: 0xb, Height: 0xc, XAdvance: 0xc, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xff, 0x87, 0x1c, 0xe3, 0x9c, 0x73, 0x9c, 0x7f, 0xe, 0x71, 0xc7, 0x38, 0xe7, 0x1c, 0xe7, 0x7f, 0xc0}},
		/* C */ tinyfont.Glyph{Rune: 67, Width: 0xb, Height: 0xc, XAdvance: 0xd, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x1f, 0x26, 0x1d, 0xc1, 0xb0, 0x1e, 0x1, 0xc0, 0x38, 0x7, 0x0, 0xe0, 0xe, 0x4, 0xe1, 0xf, 0xc0}},
		/* D */ tinyfont.Glyph{Rune: 68, Width: 0xb, Height: 0xc, XAdvance: 0xd, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xff, 0xe, 0x71, 0xc7, 0x38, 0x77, 0xe, 0xe1, 0xdc, 0x3b, 0x87, 0x70, 0xce, 0x39, 0xc6, 0x7f, 0x80}},
		/* E */ tinyfont.Glyph{Rune: 69, Width: 0xb, Height: 0xc, XAdvance: 0xc, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xff, 0xce, 0x19, 0xc1, 0x38, 0x87, 0x30, 0xfe, 0x1c, 0xc3, 0x88, 0x70, 0x2e, 0xd, 0xc3, 0x7f, 0xe0}},
		/* F */ tinyfont.Glyph{Rune: 70, Width: 0xa, Height: 0xc, XAdvance: 0xb, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xff, 0xdc, 0x37, 0x5, 0xc4, 0x73, 0x1f, 0xc7, 0x31, 0xc4, 0x70, 0x1c, 0x7, 0x3, 0xe0}},
		/* G */ tinyfont.Glyph{Rune: 71, Width: 0xc, Height: 0xc, XAdvance: 0xe, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x1f, 0x23, 0xe, 0x70, 0x6e, 0x2, 0xe0, 0xe, 0x0, 0xe1, 0xfe, 0xe, 0x60, 0xe7, 0xe, 0x38, 0xe0, 0xf8}},
		/* H */ tinyfont.Glyph{Rune: 72, Width: 0xc, Height: 0xc, XAdvance: 0xe, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xf9, 0xf7, 0xe, 0x70, 0xe7, 0xe, 0x70, 0xe7, 0xfe, 0x70, 0xe7, 0xe, 0x70, 0xe7, 0xe, 0x70, 0xef, 0x9f}},
		/* I */ tinyfont.Glyph{Rune: 73, Width: 0x5, Height: 0xc, XAdvance: 0x7, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xfb, 0x9c, 0xe7, 0x39, 0xce, 0x73, 0x9d, 0xf0}},
		/* J */ tinyfont.Glyph{Rune: 74, Width: 0x8, Height: 0xe, XAdvance: 0x9, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x1f, 0xe, 0xe, 0xe, 0xe, 0xe, 0xe, 0xe, 0xe, 0xe, 0xe, 0xce, 0xcc, 0x78}},
		/* K */ tinyfont.Glyph{Rune: 75, Width: 0xd, Height: 0xc, XAdvance: 0xe, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xf9, 0xf3, 0x82, 0x1c, 0x20, 0xe2, 0x7, 0x20, 0x3f, 0x1, 0xdc, 0xe, 0x70, 0x73, 0xc3, 0x8f, 0x1c, 0x3d, 0xf3, 0xf0}},
		/* L */ tinyfont.Glyph{Rune: 76, Width: 0xb, Height: 0xc, XAdvance: 0xc, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xf8, 0xe, 0x1, 0xc0, 0x38, 0x7, 0x0, 0xe0, 0x1c, 0x3, 0x80, 0x70, 0x2e, 0x9, 0xc3, 0x7f, 0xe0}},
		/* M */ tinyfont.Glyph{Rune: 77, Width: 0x10, Height: 0xc, XAdvance: 0x11, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xf8, 0xf, 0x3c, 0x1e, 0x3c, 0x1e, 0x2e, 0x2e, 0x2e, 0x2e, 0x26, 0x4e, 0x27, 0x4e, 0x27, 0x4e, 0x23, 0x8e, 0x23, 0x8e, 0x21, 0xe, 0x71, 0x1f}},
		/* N */ tinyfont.Glyph{Rune: 78, Width: 0xb, Height: 0xc, XAdvance: 0xd, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xf0, 0xee, 0x9, 0xe1, 0x3e, 0x25, 0xe4, 0x9e, 0x91, 0xd2, 0x1e, 0x43, 0xc8, 0x39, 0x3, 0x70, 0x20}},
		/* O */ tinyfont.Glyph{Rune: 79, Width: 0xc, Height: 0xc, XAdvance: 0xe, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x1f, 0x83, 0xc, 0x70, 0xee, 0x7, 0xe0, 0x7e, 0x7, 0xe0, 0x7e, 0x7, 0xe0, 0x77, 0xe, 0x30, 0xc1, 0xf8}},
		/* P */ tinyfont.Glyph{Rune: 80, Width: 0xa, Height: 0xc, XAdvance: 0xb, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xff, 0x1c, 0xe7, 0x1d, 0xc7, 0x71, 0xdc, 0xe7, 0xf1, 0xc0, 0x70, 0x1c, 0x7, 0x3, 0xe0}},
		/* Q */ tinyfont.Glyph{Rune: 81, Width: 0xc, Height: 0xf, XAdvance: 0xe, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xf, 0x83, 0x9c, 0x70, 0xe6, 0x6, 0xe0, 0x7e, 0x7, 0xe0, 0x7e, 0x7, 0xe0, 0x76, 0x6, 0x30, 0xc1, 0x98, 0xf, 0x0, 0x78, 0x3, 0xe0}},
		/* R */ tinyfont.Glyph{Rune: 82, Width: 0xc, Height: 0xc, XAdvance: 0xd, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xff, 0x7, 0x38, 0x71, 0xc7, 0x1c, 0x71, 0xc7, 0x38, 0x7e, 0x7, 0x70, 0x77, 0x87, 0x3c, 0x71, 0xef, 0x8f}},
		/* S */ tinyfont.Glyph{Rune: 83, Width: 0x8, Height: 0xc, XAdvance: 0xa, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x39, 0x47, 0xc1, 0xc0, 0xf0, 0x7c, 0x3e, 0xf, 0x83, 0xc3, 0xc6, 0xbc}},
		/* T */ tinyfont.Glyph{Rune: 84, Width: 0xc, Height: 0xc, XAdvance: 0xc, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xff, 0xfc, 0xe3, 0x8e, 0x10, 0xe0, 0xe, 0x0, 0xe0, 0xe, 0x0, 0xe0, 0xe, 0x0, 0xe0, 0xe, 0x1, 0xf0}},
		/* U */ tinyfont.Glyph{Rune: 85, Width: 0xb, Height: 0xc, XAdvance: 0xd, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xf8, 0xee, 0x9, 0xc1, 0x38, 0x27, 0x4, 0xe0, 0x9c, 0x13, 0x82, 0x70, 0x4e, 0x8, 0xe2, 0xf, 0x80}},
		/* V */ tinyfont.Glyph{Rune: 86, Width: 0xd, Height: 0xd, XAdvance: 0xd, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xfc, 0x7b, 0xc1, 0xe, 0x8, 0x70, 0x81, 0xc4, 0xe, 0x20, 0x7a, 0x1, 0xd0, 0xe, 0x80, 0x38, 0x1, 0xc0, 0x4, 0x0, 0x20, 0x0}},
		/* W */ tinyfont.Glyph{Rune: 87, Width: 0x12, Height: 0xc, XAdvance: 0x12, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xfd, 0xfb, 0xdc, 0x38, 0x43, 0x87, 0x10, 0xe1, 0xc4, 0x38, 0xf2, 0x7, 0x2e, 0x81, 0xd3, 0xa0, 0x34, 0x70, 0xe, 0x1c, 0x3, 0x87, 0x0, 0x60, 0x80, 0x10, 0x20}},
		/* X */ tinyfont.Glyph{Rune: 88, Width: 0xd, Height: 0xc, XAdvance: 0xd, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xfe, 0xf3, 0xc3, 0xf, 0x10, 0x39, 0x0, 0xf0, 0x3, 0x80, 0x1e, 0x1, 0x70, 0x9, 0xc0, 0x8f, 0x8, 0x3d, 0xf3, 0xf0}},
		/* Y */ tinyfont.Glyph{Rune: 89, Width: 0xd, Height: 0xc, XAdvance: 0xd, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xfc, 0x7b, 0xc1, 0x8e, 0x8, 0x38, 0x81, 0xe8, 0x7, 0x40, 0x1c, 0x0, 0xe0, 0x7, 0x0, 0x38, 0x1, 0xc0, 0x1f, 0x0}},
		/* Z */ tinyfont.Glyph{Rune: 90, Width: 0xb, Height: 0xc, XAdvance: 0xc, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xff, 0xd8, 0x72, 0x1e, 0x43, 0x80, 0xe0, 0x1c, 0x7, 0x1, 0xc0, 0x38, 0x2e, 0xf, 0x83, 0x7f, 0xe0}},
		/* [ */ tinyfont.Glyph{Rune: 91, Width: 0x4, Height: 0xf, XAdvance: 0x6, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xfc, 0xcc, 0xcc, 0xcc, 0xcc, 0xcc, 0xcc, 0xf0}},
		/* \ */ tinyfont.Glyph{Rune: 92, Width: 0x6, Height: 0xd, XAdvance: 0x5, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xc1, 0x6, 0x18, 0x20, 0xc3, 0x4, 0x18, 0x60, 0x83, 0xc}},
		/* ] */ tinyfont.Glyph{Rune: 93, Width: 0x4, Height: 0xf, XAdvance: 0x6, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xf3, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0xf0}},
		/* ^ */ tinyfont.Glyph{Rune: 94, Width: 0x8, Height: 0x7, XAdvance: 0xa, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x18, 0x1c, 0x34, 0x26, 0x62, 0x43, 0xc1}},
		/* _ */ tinyfont.Glyph{Rune: 95, Width: 0x9, Height: 0x1, XAdvance: 0x9, XOffset: 0, YOffset: 3, Bitmaps: []uint8{0xff, 0x80}},
		/* ` */ tinyfont.Glyph{Rune: 96, Width: 0x4, Height: 0x3, XAdvance: 0x6, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0xc6, 0x30}},
		/* a */ tinyfont.Glyph{Rune: 97, Width: 0x9, Height: 0x9, XAdvance: 0x9, XOffset: 0, YOffset: -8, Bitmaps: []uint8{0x7c, 0x63, 0xb1, 0xc0, 0xe1, 0xf3, 0x3b, 0x9d, 0xce, 0xff, 0x80}},
		/* b */ tinyfont.Glyph{Rune: 98, Width: 0xa, Height: 0xc, XAdvance: 0xa, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xf0, 0x1c, 0x7, 0x1, 0xdc, 0x7b, 0x9c, 0x77, 0x1d, 0xc7, 0x71, 0xdc, 0x77, 0x39, 0x3c}},
		/* c */ tinyfont.Glyph{Rune: 99, Width: 0x7, Height: 0x9, XAdvance: 0x8, XOffset: 0, YOffset: -8, Bitmaps: []uint8{0x3c, 0xed, 0x9f, 0xe, 0x1c, 0x38, 0x39, 0x3c}},
		/* d */ tinyfont.Glyph{Rune: 100, Width: 0xa, Height: 0xc, XAdvance: 0xa, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x7, 0x80, 0xe0, 0x38, 0xee, 0x77, 0xb8, 0xee, 0x3b, 0x8e, 0xe3, 0xb8, 0xe7, 0x78, 0xef}},
		/* e */ tinyfont.Glyph{Rune: 101, Width: 0x8, Height: 0x9, XAdvance: 0x8, XOffset: 0, YOffset: -8, Bitmaps: []uint8{0x3c, 0x66, 0xe6, 0xfe, 0xe0, 0xe0, 0xe0, 0x72, 0x3c}},
		/* f */ tinyfont.Glyph{Rune: 102, Width: 0x7, Height: 0xc, XAdvance: 0x7, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x3e, 0xed, 0xc7, 0xc7, 0xe, 0x1c, 0x38, 0x70, 0xe1, 0xc7, 0xc0}},
		/* g */ tinyfont.Glyph{Rune: 103, Width: 0x7, Height: 0xd, XAdvance: 0x9, XOffset: 1, YOffset: -8, Bitmaps: []uint8{0x31, 0xdf, 0xbf, 0x7e, 0xe7, 0x90, 0x60, 0xfc, 0xfe, 0xc, 0x17, 0xc0}},
		/* h */ tinyfont.Glyph{Rune: 104, Width: 0xa, Height: 0xc, XAdvance: 0xa, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xf0, 0x1c, 0x7, 0x1, 0xdc, 0x7b, 0x9c, 0xe7, 0x39, 0xce, 0x73, 0x9c, 0xe7, 0x3b, 0xff}},
		/* i */ tinyfont.Glyph{Rune: 105, Width: 0x5, Height: 0xc, XAdvance: 0x5, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x73, 0x9d, 0xe7, 0x39, 0xce, 0x73, 0x9d, 0xf0}},
		/* j */ tinyfont.Glyph{Rune: 106, Width: 0x6, Height: 0x10, XAdvance: 0x7, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0x1c, 0x71, 0xcf, 0x1c, 0x71, 0xc7, 0x1c, 0x71, 0xc7, 0x1c, 0x7d, 0xbe}},
		/* k */ tinyfont.Glyph{Rune: 107, Width: 0xa, Height: 0xc, XAdvance: 0xa, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xf0, 0x1c, 0x7, 0x1, 0xce, 0x71, 0x1c, 0x87, 0x41, 0xf8, 0x77, 0x1c, 0xe7, 0x1b, 0xef}},
		/* l */ tinyfont.Glyph{Rune: 108, Width: 0x5, Height: 0xc, XAdvance: 0x5, XOffset: 0, YOffset: -11, Bitmaps: []uint8{0xf3, 0x9c, 0xe7, 0x39, 0xce, 0x73, 0x9d, 0xf0}},
		/* m */ tinyfont.Glyph{Rune: 109, Width: 0xf, Height: 0x9, XAdvance: 0xf, XOffset: 0, YOffset: -8, Bitmaps: []uint8{0xf7, 0x38, 0xf7, 0xb9, 0xce, 0x73, 0x9c, 0xe7, 0x39, 0xce, 0x73, 0x9c, 0xe7, 0x39, 0xce, 0xff, 0xfe}},
		/* n */ tinyfont.Glyph{Rune: 110, Width: 0xa, Height: 0x9, XAdvance: 0xa, XOffset: 0, YOffset: -8, Bitmaps: []uint8{0xf7, 0x1e, 0xe7, 0x39, 0xce, 0x73, 0x9c, 0xe7, 0x39, 0xce, 0xff, 0xc0}},
		/* o */ tinyfont.Glyph{Rune: 111, Width: 0x9, Height: 0x9, XAdvance: 0x9, XOffset: 0, YOffset: -8, Bitmaps: []uint8{0x3e, 0x31, 0xb8, 0xfc, 0x7e, 0x3f, 0x1f, 0x8e, 0xc6, 0x3e, 0x0}},
		/* p */ tinyfont.Glyph{Rune: 112, Width: 0xa, Height: 0xd, XAdvance: 0xa, XOffset: 0, YOffset: -8, Bitmaps: []uint8{0xf7, 0x1e, 0xe7, 0x1d, 0xc7, 0x71, 0xdc, 0x77, 0x1d, 0xce, 0x7f, 0x1c, 0x7, 0x1, 0xc0, 0xf8, 0x0}},
		/* q */ tinyfont.Glyph{Rune: 113, Width: 0xa, Height: 0xd, XAdvance: 0xa, XOffset: 0, YOffset: -8, Bitmaps: []uint8{0x3c, 0x9c, 0xee, 0x3b, 0x8e, 0xe3, 0xb8, 0xee, 0x39, 0xce, 0x3f, 0x80, 0xe0, 0x38, 0xe, 0x7, 0xc0}},
		/* r */ tinyfont.Glyph{Rune: 114, Width: 0x8, Height: 0x9, XAdvance: 0x8, XOffset: 0, YOffset: -8, Bitmaps: []uint8{0xf7, 0x7b, 0x70, 0x70, 0x70, 0x70, 0x70, 0x70, 0xf8}},
		/* s */ tinyfont.Glyph{Rune: 115, Width: 0x5, Height: 0x9, XAdvance: 0x7, XOffset: 1, YOffset: -8, Bitmaps: []uint8{0x7e, 0x73, 0xc7, 0x8e, 0x39, 0xb0}},
		/* t */ tinyfont.Glyph{Rune: 116, Width: 0x6, Height: 0xb, XAdvance: 0x6, XOffset: 0, YOffset: -10, Bitmaps: []uint8{0x10, 0xcf, 0x9c, 0x71, 0xc7, 0x1c, 0x71, 0xd3, 0x80}},
		/* u */ tinyfont.Glyph{Rune: 117, Width: 0xa, Height: 0x9, XAdvance: 0xa, XOffset: 0, YOffset: -8, Bitmaps: []uint8{0xf7, 0x9c, 0xe7, 0x39, 0xce, 0x73, 0x9c, 0xe7, 0x39, 0xce, 0x3f, 0xc0}},
		/* v */ tinyfont.Glyph{Rune: 118, Width: 0x9, Height: 0x9, XAdvance: 0x9, XOffset: 0, YOffset: -8, Bitmaps: []uint8{0xfb, 0xb8, 0x8c, 0x87, 0x43, 0xc0, 0xe0, 0x70, 0x10, 0x8, 0x0}},
		/* w */ tinyfont.Glyph{Rune: 119, Width: 0xc, Height: 0x9, XAdvance: 0xd, XOffset: 0, YOffset: -8, Bitmaps: []uint8{0xf7, 0xb6, 0x31, 0x73, 0xa3, 0x3a, 0x3d, 0xa3, 0xdc, 0x18, 0xc1, 0x88, 0x10, 0x80}},
		/* x */ tinyfont.Glyph{Rune: 120, Width: 0x9, Height: 0x9, XAdvance: 0x9, XOffset: 0, YOffset: -8, Bitmaps: []uint8{0xfb, 0xb8, 0x8e, 0x83, 0x81, 0xc0, 0xf0, 0x98, 0xce, 0xef, 0x80}},
		/* y */ tinyfont.Glyph{Rune: 121, Width: 0x8, Height: 0xd, XAdvance: 0x9, XOffset: 0, YOffset: -8, Bitmaps: []uint8{0xf7, 0x62, 0x72, 0x34, 0x34, 0x3c, 0x18, 0x18, 0x10, 0x10, 0x10, 0xe0, 0xe0}},
		/* z */ tinyfont.Glyph{Rune: 122, Width: 0x7, Height: 0x9, XAdvance: 0x8, XOffset: 1, YOffset: -8, Bitmaps: []uint8{0xff, 0x1c, 0x70, 0xe3, 0x87, 0x1c, 0x71, 0xfe}},
		/* { */ tinyfont.Glyph{Rune: 123, Width: 0x5, Height: 0x10, XAdvance: 0x7, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0x19, 0x8c, 0x63, 0x18, 0xcc, 0x61, 0x8c, 0x63, 0x18, 0xc3}},
		/* | */ tinyfont.Glyph{Rune: 124, Width: 0x1, Height: 0xd, XAdvance: 0x4, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0xff, 0xf8}},
		/* } */ tinyfont.Glyph{Rune: 125, Width: 0x5, Height: 0x10, XAdvance: 0x7, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0xc3, 0x18, 0xc6, 0x31, 0x86, 0x33, 0x18, 0xc6, 0x31, 0x98}},
		/* ~ */ tinyfont.Glyph{Rune: 126, Width: 0x8, Height: 0x2, XAdvance: 0x9, XOffset: 1, YOffset: -4, Bitmaps: []uint8{0xf0, 0x8e}},
	},

	YAdvance: 0x16,
}
