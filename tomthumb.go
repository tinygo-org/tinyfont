/**
** The original 3x5 font is licensed under the 3-clause BSD license:
**
** Copyright 1999 Brian J. Swetland
** Copyright 1999 Vassilii Khachaturov
** Portions (of vt100.c/vt100.h) copyright Dan Marks
**
** All rights reserved.
**
** Redistribution and use in source and binary forms, with or without
** modification, are permitted provided that the following conditions
** are met:
** 1. Redistributions of source code must retain the above copyright
**    notice, this list of conditions, and the following disclaimer.
** 2. Redistributions in binary form must reproduce the above copyright
**    notice, this list of conditions, and the following disclaimer in the
**    documentation and/or other materials provided with the distribution.
** 3. The name of the authors may not be used to endorse or promote products
**    derived from this software without specific prior written permission.
**
** THIS SOFTWARE IS PROVIDED BY THE AUTHOR ``AS IS'' AND ANY EXPRESS OR
** IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
** OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
** IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT,
** INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
** NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
** DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
** THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
** (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
** THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
**
** Modifications to Tom Thumb for improved readability are from Robey Pointer,
** see:
** http://robey.lag.net/2010/01/23/tiny-monospace-font.html
**
** The original author does not have any objection to relicensing of Robey
** Pointer's modifications (in this file) in a more permissive license.  See
** the discussion at the above blog, and also here:
** http://opengameart.org/forumtopic/how-to-submit-art-using-the-3-clause-bsd-license
**
** Feb 21, 2016: Conversion from Linux BDF --> Adafruit GFX font,
** with the help of this Python script:
** https://gist.github.com/skelliam/322d421f028545f16f6d
** William Skellenger (williamj@skellenger.net)
** Twitter: @skelliam
**
 */

package tinyfont

var TomThumb = Font{
	Glyphs: []Glyph{
		/*   */ Glyph{Rune: 32, Width: 0x8, Height: 0x1, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x0}},
		/* ! */ Glyph{Rune: 33, Width: 0x8, Height: 0x5, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80, 0x80, 0x0, 0x80}},
		/* " */ Glyph{Rune: 34, Width: 0x8, Height: 0x2, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0}},
		/* # */ Glyph{Rune: 35, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xe0, 0xa0, 0xe0, 0xa0}},
		/* $ */ Glyph{Rune: 36, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xc0, 0x60, 0xc0, 0x40}},
		/* % */ Glyph{Rune: 37, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x20, 0x40, 0x80, 0x20}},
		/* & */ Glyph{Rune: 38, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xc0, 0xe0, 0xa0, 0x60}},
		/* ' */ Glyph{Rune: 39, Width: 0x8, Height: 0x2, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80}},
		/* ( */ Glyph{Rune: 40, Width: 0x8, Height: 0x5, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0x80, 0x80, 0x40}},
		/* ) */ Glyph{Rune: 41, Width: 0x8, Height: 0x5, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x40, 0x40, 0x40, 0x80}},
		/* * */ Glyph{Rune: 42, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x40, 0xa0}},
		/* + */ Glyph{Rune: 43, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x40, 0xe0, 0x40}},
		/* , */ Glyph{Rune: 44, Width: 0x8, Height: 0x2, XAdvance: 0x3, XOffset: 0, YOffset: -2, Bitmaps: []uint8{0x40, 0x80}},
		/* - */ Glyph{Rune: 45, Width: 0x8, Height: 0x1, XAdvance: 0x4, XOffset: 0, YOffset: -3, Bitmaps: []uint8{0xe0}},
		/* . */ Glyph{Rune: 46, Width: 0x8, Height: 0x1, XAdvance: 0x2, XOffset: 0, YOffset: -1, Bitmaps: []uint8{0x80}},
		/* / */ Glyph{Rune: 47, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x20, 0x40, 0x80, 0x80}},
		/* 0 */ Glyph{Rune: 48, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xa0, 0xa0, 0xa0, 0xc0}},
		/* 1 */ Glyph{Rune: 49, Width: 0x8, Height: 0x5, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xc0, 0x40, 0x40, 0x40}},
		/* 2 */ Glyph{Rune: 50, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x20, 0x40, 0x80, 0xe0}},
		/* 3 */ Glyph{Rune: 51, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x20, 0x40, 0x20, 0xc0}},
		/* 4 */ Glyph{Rune: 52, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0xe0, 0x20, 0x20}},
		/* 5 */ Glyph{Rune: 53, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x80, 0xc0, 0x20, 0xc0}},
		/* 6 */ Glyph{Rune: 54, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x80, 0xe0, 0xa0, 0xe0}},
		/* 7 */ Glyph{Rune: 55, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x20, 0x40, 0x80, 0x80}},
		/* 8 */ Glyph{Rune: 56, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0xa0, 0xe0, 0xa0, 0xe0}},
		/* 9 */ Glyph{Rune: 57, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0xa0, 0xe0, 0x20, 0xc0}},
		/* : */ Glyph{Rune: 58, Width: 0x8, Height: 0x3, XAdvance: 0x2, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x80, 0x0, 0x80}},
		/* ; */ Glyph{Rune: 59, Width: 0x8, Height: 0x4, XAdvance: 0x3, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x40, 0x0, 0x40, 0x80}},
		/* < */ Glyph{Rune: 60, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x40, 0x80, 0x40, 0x20}},
		/* = */ Glyph{Rune: 61, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xe0, 0x0, 0xe0}},
		/* > */ Glyph{Rune: 62, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x40, 0x20, 0x40, 0x80}},
		/* ? */ Glyph{Rune: 63, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x20, 0x40, 0x0, 0x40}},
		/* @ */ Glyph{Rune: 64, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xa0, 0xe0, 0x80, 0x60}},
		/* A */ Glyph{Rune: 65, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xa0, 0xe0, 0xa0, 0xa0}},
		/* B */ Glyph{Rune: 66, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xa0, 0xc0, 0xa0, 0xc0}},
		/* C */ Glyph{Rune: 67, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x80, 0x80, 0x80, 0x60}},
		/* D */ Glyph{Rune: 68, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xa0, 0xa0, 0xa0, 0xc0}},
		/* E */ Glyph{Rune: 69, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x80, 0xe0, 0x80, 0xe0}},
		/* F */ Glyph{Rune: 70, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x80, 0xe0, 0x80, 0x80}},
		/* G */ Glyph{Rune: 71, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x80, 0xe0, 0xa0, 0x60}},
		/* H */ Glyph{Rune: 72, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0xe0, 0xa0, 0xa0}},
		/* I */ Glyph{Rune: 73, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x40, 0x40, 0x40, 0xe0}},
		/* J */ Glyph{Rune: 74, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x20, 0x20, 0xa0, 0x40}},
		/* K */ Glyph{Rune: 75, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0xc0, 0xa0, 0xa0}},
		/* L */ Glyph{Rune: 76, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80, 0x80, 0x80, 0xe0}},
		/* M */ Glyph{Rune: 77, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xe0, 0xe0, 0xa0, 0xa0}},
		/* N */ Glyph{Rune: 78, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xe0, 0xe0, 0xe0, 0xa0}},
		/* O */ Glyph{Rune: 79, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xa0, 0xa0, 0xa0, 0x40}},
		/* P */ Glyph{Rune: 80, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xa0, 0xc0, 0x80, 0x80}},
		/* Q */ Glyph{Rune: 81, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xa0, 0xa0, 0xe0, 0x60}},
		/* R */ Glyph{Rune: 82, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xa0, 0xe0, 0xc0, 0xa0}},
		/* S */ Glyph{Rune: 83, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x80, 0x40, 0x20, 0xc0}},
		/* T */ Glyph{Rune: 84, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x40, 0x40, 0x40, 0x40}},
		/* U */ Glyph{Rune: 85, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0xa0, 0xa0, 0x60}},
		/* V */ Glyph{Rune: 86, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0xa0, 0x40, 0x40}},
		/* W */ Glyph{Rune: 87, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0xe0, 0xe0, 0xa0}},
		/* X */ Glyph{Rune: 88, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0x40, 0xa0, 0xa0}},
		/* Y */ Glyph{Rune: 89, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0x40, 0x40, 0x40}},
		/* Z */ Glyph{Rune: 90, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x20, 0x40, 0x80, 0xe0}},
		/* [ */ Glyph{Rune: 91, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x80, 0x80, 0x80, 0xe0}},
		/* \ */ Glyph{Rune: 92, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x80, 0x40, 0x20}},
		/* ] */ Glyph{Rune: 93, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x20, 0x20, 0x20, 0xe0}},
		/* ^ */ Glyph{Rune: 94, Width: 0x8, Height: 0x2, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xa0}},
		/* _ */ Glyph{Rune: 95, Width: 0x8, Height: 0x1, XAdvance: 0x4, XOffset: 0, YOffset: -1, Bitmaps: []uint8{0xe0}},
		/* ` */ Glyph{Rune: 96, Width: 0x8, Height: 0x2, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x40}},
		/* a */ Glyph{Rune: 97, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xc0, 0x60, 0xa0, 0xe0}},
		/* b */ Glyph{Rune: 98, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0xc0, 0xa0, 0xa0, 0xc0}},
		/* c */ Glyph{Rune: 99, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0x80, 0x80, 0x60}},
		/* d */ Glyph{Rune: 100, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x60, 0xa0, 0xa0, 0x60}},
		/* e */ Glyph{Rune: 101, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0xa0, 0xc0, 0x60}},
		/* f */ Glyph{Rune: 102, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x40, 0xe0, 0x40, 0x40}},
		/* g */ Glyph{Rune: 103, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0xa0, 0xe0, 0x20, 0x40}},
		/* h */ Glyph{Rune: 104, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0xc0, 0xa0, 0xa0, 0xa0}},
		/* i */ Glyph{Rune: 105, Width: 0x8, Height: 0x5, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x0, 0x80, 0x80, 0x80}},
		/* j */ Glyph{Rune: 106, Width: 0x8, Height: 0x6, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x0, 0x20, 0x20, 0xa0, 0x40}},
		/* k */ Glyph{Rune: 107, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0xa0, 0xc0, 0xc0, 0xa0}},
		/* l */ Glyph{Rune: 108, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x40, 0x40, 0x40, 0xe0}},
		/* m */ Glyph{Rune: 109, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xe0, 0xe0, 0xe0, 0xa0}},
		/* n */ Glyph{Rune: 110, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xc0, 0xa0, 0xa0, 0xa0}},
		/* o */ Glyph{Rune: 111, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x40, 0xa0, 0xa0, 0x40}},
		/* p */ Glyph{Rune: 112, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xc0, 0xa0, 0xa0, 0xc0, 0x80}},
		/* q */ Glyph{Rune: 113, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0xa0, 0xa0, 0x60, 0x20}},
		/* r */ Glyph{Rune: 114, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0x80, 0x80, 0x80}},
		/* s */ Glyph{Rune: 115, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0xc0, 0x60, 0xc0}},
		/* t */ Glyph{Rune: 116, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xe0, 0x40, 0x40, 0x60}},
		/* u */ Glyph{Rune: 117, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xa0, 0xa0, 0xa0, 0x60}},
		/* v */ Glyph{Rune: 118, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xa0, 0xa0, 0xe0, 0x40}},
		/* w */ Glyph{Rune: 119, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xa0, 0xe0, 0xe0, 0xe0}},
		/* x */ Glyph{Rune: 120, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xa0, 0x40, 0x40, 0xa0}},
		/* y */ Glyph{Rune: 121, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xa0, 0xa0, 0x60, 0x20, 0x40}},
		/* z */ Glyph{Rune: 122, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xe0, 0x60, 0xc0, 0xe0}},
		/* { */ Glyph{Rune: 123, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x40, 0x80, 0x40, 0x60}},
		/* | */ Glyph{Rune: 124, Width: 0x8, Height: 0x5, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80, 0x0, 0x80, 0x80}},
		/* } */ Glyph{Rune: 125, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x40, 0x20, 0x40, 0xc0}},
		/* ~ */ Glyph{Rune: 126, Width: 0x8, Height: 0x2, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xc0}},
		/*  */ Glyph{Rune: 127, Width: 0x8, Height: 0x5, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x0, 0x80, 0x80, 0x80}},
		/*  */ Glyph{Rune: 128, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xe0, 0x80, 0xe0, 0x40}},
		/*  */ Glyph{Rune: 129, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x40, 0xe0, 0x40, 0xe0}},
		/*  */ Glyph{Rune: 130, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x40, 0xe0, 0x40, 0xa0}},
		/*  */ Glyph{Rune: 131, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0x40, 0xe0, 0x40}},
		/*  */ Glyph{Rune: 132, Width: 0x8, Height: 0x5, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80, 0x0, 0x80, 0x80}},
		/*  */ Glyph{Rune: 133, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x40, 0xa0, 0x40, 0xc0}},
		/*  */ Glyph{Rune: 134, Width: 0x8, Height: 0x1, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0}},
		/*  */ Glyph{Rune: 135, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x80, 0x60}},
		/*  */ Glyph{Rune: 136, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xa0, 0xe0, 0x0, 0xe0}},
		/*  */ Glyph{Rune: 137, Width: 0x8, Height: 0x3, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0x40}},
		/*  */ Glyph{Rune: 138, Width: 0x8, Height: 0x2, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xe0, 0x20}},
		/*  */ Glyph{Rune: 139, Width: 0x8, Height: 0x1, XAdvance: 0x3, XOffset: 0, YOffset: -3, Bitmaps: []uint8{0xc0}},
		/*  */ Glyph{Rune: 140, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xc0, 0xa0}},
		/*  */ Glyph{Rune: 141, Width: 0x8, Height: 0x1, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0}},
		/*  */ Glyph{Rune: 142, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xa0, 0x40}},
		/*  */ Glyph{Rune: 143, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xe0, 0x40, 0x0, 0xe0}},
		/*  */ Glyph{Rune: 144, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x40, 0x60}},
		/*  */ Glyph{Rune: 145, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x60, 0xe0}},
		/*  */ Glyph{Rune: 146, Width: 0x8, Height: 0x2, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80}},
		/*  */ Glyph{Rune: 147, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0xa0, 0xc0, 0x80}},
		/*  */ Glyph{Rune: 148, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xa0, 0x60, 0x60, 0x60}},
		/*  */ Glyph{Rune: 149, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xe0, 0xe0, 0xe0}},
		/*  */ Glyph{Rune: 150, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -3, Bitmaps: []uint8{0x40, 0x20, 0xc0}},
		/*  */ Glyph{Rune: 151, Width: 0x8, Height: 0x3, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80, 0x80}},
		/*  */ Glyph{Rune: 152, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xa0, 0x40, 0x0, 0xe0}},
		/*  */ Glyph{Rune: 153, Width: 0x8, Height: 0x3, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x40, 0x80}},
		/*  */ Glyph{Rune: 154, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80, 0x0, 0x60, 0x20}},
		/*  */ Glyph{Rune: 155, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80, 0x0, 0xc0, 0x60}},
		/*  */ Glyph{Rune: 156, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xc0, 0x0, 0x60, 0x20}},
		/*  */ Glyph{Rune: 157, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x0, 0x40, 0x80, 0xe0}},
		/*  */ Glyph{Rune: 158, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x20, 0x40, 0xe0, 0xa0}},
		/*  */ Glyph{Rune: 159, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0x40, 0xe0, 0xa0}},
		/*   */ Glyph{Rune: 160, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0x40, 0xe0, 0xa0}},
		/* ¡ */ Glyph{Rune: 161, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xc0, 0x40, 0xe0, 0xa0}},
		/* ¢ */ Glyph{Rune: 162, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x40, 0xa0, 0xe0, 0xa0}},
		/* £ */ Glyph{Rune: 163, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xc0, 0xa0, 0xe0, 0xa0}},
		/* ¤ */ Glyph{Rune: 164, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xc0, 0xe0, 0xc0, 0xe0}},
		/* ¥ */ Glyph{Rune: 165, Width: 0x8, Height: 0x6, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x80, 0x80, 0x60, 0x20, 0x40}},
		/* ¦ */ Glyph{Rune: 166, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x20, 0xe0, 0xc0, 0xe0}},
		/* § */ Glyph{Rune: 167, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0xe0, 0xc0, 0xe0}},
		/* ¨ */ Glyph{Rune: 168, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0xe0, 0xc0, 0xe0}},
		/* © */ Glyph{Rune: 169, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0xe0, 0xc0, 0xe0}},
		/* ª */ Glyph{Rune: 170, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x20, 0xe0, 0x40, 0xe0}},
		/* « */ Glyph{Rune: 171, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0xe0, 0x40, 0xe0}},
		/* ¬ */ Glyph{Rune: 172, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0xe0, 0x40, 0xe0}},
		/* ­ */ Glyph{Rune: 173, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0xe0, 0x40, 0xe0}},
		/* ® */ Glyph{Rune: 174, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xa0, 0xe0, 0xa0, 0xc0}},
		/* ¯ */ Glyph{Rune: 175, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x60, 0xa0, 0xe0, 0xa0}},
		/* ° */ Glyph{Rune: 176, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x20, 0xe0, 0xa0, 0xe0}},
		/* ± */ Glyph{Rune: 177, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0xe0, 0xa0, 0xe0}},
		/* ² */ Glyph{Rune: 178, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0xe0, 0xa0, 0xe0}},
		/* ³ */ Glyph{Rune: 179, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x60, 0xe0, 0xa0, 0xe0}},
		/* ´ */ Glyph{Rune: 180, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0xe0, 0xa0, 0xe0}},
		/* µ */ Glyph{Rune: 181, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xa0, 0x40, 0xa0}},
		/* ¶ */ Glyph{Rune: 182, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xa0, 0xe0, 0xa0, 0xc0}},
		/* · */ Glyph{Rune: 183, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x40, 0xa0, 0xa0, 0xe0}},
		/* ¸ */ Glyph{Rune: 184, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x40, 0xa0, 0xa0, 0xe0}},
		/* ¹ */ Glyph{Rune: 185, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0xa0, 0xa0, 0xe0}},
		/* º */ Glyph{Rune: 186, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0xa0, 0xa0, 0xe0}},
		/* » */ Glyph{Rune: 187, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x40, 0xa0, 0xe0, 0x40}},
		/* ¼ */ Glyph{Rune: 188, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0xe0, 0xa0, 0xe0, 0x80}},
		/* ½ */ Glyph{Rune: 189, Width: 0x8, Height: 0x6, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xa0, 0xc0, 0xa0, 0xc0, 0x80}},
		/* ¾ */ Glyph{Rune: 190, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x20, 0x60, 0xa0, 0xe0}},
		/* ¿ */ Glyph{Rune: 191, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0x60, 0xa0, 0xe0}},
		/* À */ Glyph{Rune: 192, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0x60, 0xa0, 0xe0}},
		/* Á */ Glyph{Rune: 193, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xc0, 0x60, 0xa0, 0xe0}},
		/* Â */ Glyph{Rune: 194, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0x60, 0xa0, 0xe0}},
		/* Ã */ Glyph{Rune: 195, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x60, 0x60, 0xa0, 0xe0}},
		/* Ä */ Glyph{Rune: 196, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0xe0, 0xe0, 0xc0}},
		/* Å */ Glyph{Rune: 197, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0x80, 0x60, 0x20, 0x40}},
		/* Æ */ Glyph{Rune: 198, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x20, 0x60, 0xe0, 0x60}},
		/* Ç */ Glyph{Rune: 199, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0x60, 0xe0, 0x60}},
		/* È */ Glyph{Rune: 200, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0x60, 0xe0, 0x60}},
		/* É */ Glyph{Rune: 201, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0x60, 0xe0, 0x60}},
		/* Ê */ Glyph{Rune: 202, Width: 0x8, Height: 0x5, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x40, 0x80, 0x80, 0x80}},
		/* Ë */ Glyph{Rune: 203, Width: 0x8, Height: 0x5, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0x40, 0x40, 0x40}},
		/* Ì */ Glyph{Rune: 204, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0x40, 0x40, 0x40}},
		/* Í */ Glyph{Rune: 205, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0x40, 0x40, 0x40}},
		/* Î */ Glyph{Rune: 206, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xc0, 0x60, 0xa0, 0x60}},
		/* Ï */ Glyph{Rune: 207, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x60, 0xc0, 0xa0, 0xa0}},
		/* Ð */ Glyph{Rune: 208, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x20, 0x40, 0xa0, 0x40}},
		/* Ñ */ Glyph{Rune: 209, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0x40, 0xa0, 0x40}},
		/* Ò */ Glyph{Rune: 210, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0x40, 0xa0, 0x40}},
		/* Ó */ Glyph{Rune: 211, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x60, 0x40, 0xa0, 0x40}},
		/* Ô */ Glyph{Rune: 212, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0x40, 0xa0, 0x40}},
		/* Õ */ Glyph{Rune: 213, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x0, 0xe0, 0x0, 0x40}},
		/* Ö */ Glyph{Rune: 214, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0xe0, 0xa0, 0xc0}},
		/* × */ Glyph{Rune: 215, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x40, 0xa0, 0xa0, 0x60}},
		/* Ø */ Glyph{Rune: 216, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x40, 0xa0, 0xa0, 0x60}},
		/* Ù */ Glyph{Rune: 217, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0xa0, 0xa0, 0x60}},
		/* Ú */ Glyph{Rune: 218, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0xa0, 0xa0, 0x60}},
		/* Û */ Glyph{Rune: 219, Width: 0x8, Height: 0x6, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x40, 0xa0, 0x60, 0x20, 0x40}},
		/* Ü */ Glyph{Rune: 220, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x80, 0xc0, 0xa0, 0xc0, 0x80}},
		/* Ý */ Glyph{Rune: 221, Width: 0x8, Height: 0x6, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0xa0, 0x60, 0x20, 0x40}},
		/* Þ */ Glyph{Rune: 222, Width: 0x8, Height: 0x1, XAdvance: 0x2, XOffset: 0, YOffset: -1, Bitmaps: []uint8{0x0}},
		/* ß */ Glyph{Rune: 223, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xc0, 0xe0, 0xc0, 0x60}},
		/* à */ Glyph{Rune: 224, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0xe0, 0xc0, 0xe0}},
		/* á */ Glyph{Rune: 225, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x60, 0xc0, 0x60, 0xc0}},
		/* â */ Glyph{Rune: 226, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x60, 0xc0, 0x60, 0xc0}},
		/* ã */ Glyph{Rune: 227, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0xa0, 0x40, 0x40}},
		/* ä */ Glyph{Rune: 228, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xe0, 0x60, 0xc0, 0xe0}},
		/* å */ Glyph{Rune: 229, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xe0, 0x60, 0xc0, 0xe0}},
		/* æ */ Glyph{Rune: 230, Width: 0x8, Height: 0x1, XAdvance: 0x2, XOffset: 0, YOffset: -1, Bitmaps: []uint8{0x0}},
		/* ç */ Glyph{Rune: 231, Width: 0x8, Height: 0x1, XAdvance: 0x2, XOffset: 0, YOffset: -1, Bitmaps: []uint8{0x0}},
		/* è */ Glyph{Rune: 232, Width: 0x8, Height: 0x1, XAdvance: 0x2, XOffset: 0, YOffset: -3, Bitmaps: []uint8{0x80}},
		/* é */ Glyph{Rune: 233, Width: 0x8, Height: 0x1, XAdvance: 0x4, XOffset: 0, YOffset: -1, Bitmaps: []uint8{0xa0}},
		/* ê */ Glyph{Rune: 234, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xe0, 0xe0, 0xc0, 0x60}},
		/* ë */ Glyph{Rune: 235, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0xa0, 0xa0, 0xa0, 0xe0}},
	},

	YAdvance: 0x6,
}
