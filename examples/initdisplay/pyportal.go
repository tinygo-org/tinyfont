//go:build pyportal
// +build pyportal

package initdisplay

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/ili9341"
)

// InitDisplay initializes the display of each board.
func InitDisplay() drivers.Displayer {
	d := ili9341.NewParallel(
		machine.LCD_DATA0,
		machine.TFT_WR,
		machine.TFT_DC,
		machine.TFT_CS,
		machine.TFT_RESET,
		machine.TFT_RD,
	)
	d.Configure(ili9341.Config{})
	d.FillScreen(color.RGBA{255, 255, 255, 255})

	machine.TFT_BACKLIGHT.Configure(machine.PinConfig{machine.PinOutput})
	machine.TFT_BACKLIGHT.High()

	return d
}
