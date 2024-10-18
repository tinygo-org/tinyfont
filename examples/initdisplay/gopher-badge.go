//go:build gopher_badge

package initdisplay

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/st7789"
)

// InitDisplay initializes the display of each board.
func InitDisplay() drivers.Displayer {
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
		Mode:      0,
	})

	d := st7789.New(machine.SPI0,
		machine.TFT_RST,       // TFT_RESET
		machine.TFT_WRX,       // TFT_DC
		machine.TFT_CS,        // TFT_CS
		machine.TFT_BACKLIGHT) // TFT_LITE

	d.Configure(st7789.Config{
		Rotation: st7789.NO_ROTATION,
		Height:   320,
	})
	d.FillScreen(color.RGBA{255, 255, 255, 255})

	return &d
}
