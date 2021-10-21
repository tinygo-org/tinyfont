//go:build baremetal
// +build baremetal

package notosans

import (
	"machine"

	"tinygo.org/x/drivers/flash"
)

var (
	dev     *flash.Device
	fontBuf [1024]byte
)

func init() {
	dev = flash.NewQSPI(
		machine.QSPI_CS,
		machine.QSPI_SCK,
		machine.QSPI_DATA0,
		machine.QSPI_DATA1,
		machine.QSPI_DATA2,
		machine.QSPI_DATA3,
	)
	dev.Configure(&flash.DeviceConfig{
		Identifier: flash.DefaultDeviceIdentifier,
	})
}

func (g NotoSans12ptGlyph) Bytes() []byte {
	buf := fontBuf[:2]
	for {
		_, err := dev.ReadAt(buf, int64(g.Offset))
		if err != nil {
			continue
		}
		break
	}

	buf = fontBuf[:uint32(buf[0])*uint32(buf[1])/4+5+1]
	for {
		_, err := dev.ReadAt(buf, int64(g.Offset))
		if err != nil {
			continue
		}
		break
	}
	return buf
}
