//go:build pybadge
// +build pybadge

package initdisplay

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/st7735"
)

// InitDisplay initializes the display of each board.
func InitDisplay() drivers.Displayer {
	machine.SPI1.Configure(machine.SPIConfig{
		SCK:       machine.SPI1_SCK_PIN,
		SDO:       machine.SPI1_SDO_PIN,
		SDI:       machine.SPI1_SDI_PIN,
		Frequency: 8000000,
	})

	d := st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)
	d.Configure(st7735.Config{
		Rotation: st7735.ROTATION_90,
	})

	d.FillScreen(color.RGBA{255, 255, 255, 255})

	return &d
}
