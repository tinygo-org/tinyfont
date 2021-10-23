//go:build baremetal
// +build baremetal

package twemoji

import (
	"tinygo.org/x/drivers/flash"
)

var (
	dev     *flash.Device
	fontBuf [1024]byte
	offset  = 0x000A0000
)

func FlashDevice(d *flash.Device) {
	dev = d
}

func (g Twemoji12ptGlyph) Bytes() []byte {
	if dev == nil {
		panic("dev == nil")
	}

	buf := fontBuf[:12*12*2]
	for {
		_, err := dev.ReadAt(buf, int64(g.Offset)+int64(offset))
		if err != nil {
			continue
		}
		break
	}
	return buf
}
