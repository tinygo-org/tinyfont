//go:build !baremetal
// +build !baremetal

package notosans

import (
	_ "embed"
)

//go:embed notosans.bin
var cNotoSans12pt []byte

func (g NotoSans12ptGlyph) Bytes() []byte {
	buf := cNotoSans12pt[g.Offset : g.Offset+2]
	buf = cNotoSans12pt[g.Offset : g.Offset+uint32(buf[0])*uint32(buf[1])/4+5+1]
	return buf
}
