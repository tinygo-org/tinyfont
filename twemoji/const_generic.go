//go:build !baremetal
// +build !baremetal

package twemoji

import "io/ioutil"

var cTinyFont []byte

func init() {
	var err error
	cTinyFont, err = ioutil.ReadFile("./twemoji/twemoji.bin")
	if err != nil {
		panic(err)
	}
}

func (g Twemoji12ptGlyph) Bytes() []byte {
	buf := []byte(cTinyFont[g.Offset : g.Offset+12*12*2])
	return buf
}
