//go:build !baremetal
// +build !baremetal

package notosans

import (
	"io/ioutil"
)

var cNotoSans12pt []byte

func init() {
	var err error
	cNotoSans12pt, err = ioutil.ReadFile("./ttf-notosans/notosans.bin")
	if err != nil {
		panic(err)
	}
}

func (g NotoSans12ptGlyph) Bytes() []byte {
	buf := cNotoSans12pt[g.Offset : g.Offset+2]
	buf = cNotoSans12pt[g.Offset : g.Offset+uint32(buf[0])*uint32(buf[1])/4+5+1]
	return buf
}
