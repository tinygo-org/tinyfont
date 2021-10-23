//go:build !baremetal
// +build !baremetal

package twemoji

import (
	_ "embed"
)

//go:embed twemoji.bin
var cTinyFont []byte

func (g Twemoji12ptGlyph) Bytes() []byte {
	buf := []byte(cTinyFont[g.Offset : g.Offset+12*12*2])
	return buf
}
