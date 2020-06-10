// +build wioterminal

package main

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/ili9341"
)

func initDisplay() drivers.Displayer {
	machine.SPI3.Configure(machine.SPIConfig{
		SCK:       machine.LCD_SCK_PIN,
		MOSI:      machine.LCD_MOSI_PIN,
		MISO:      machine.LCD_MISO_PIN,
		Frequency: 40000000,
	})

	d := ili9341.NewSpi(
		machine.SPI3,
		machine.LCD_DC,
		machine.LCD_SS_PIN,
		machine.LCD_RESET,
	)
	d.Configure(ili9341.Config{})
	d.FillScreen(color.RGBA{255, 255, 255, 255})

	machine.LCD_BACKLIGHT.Configure(machine.PinConfig{machine.PinOutput})
	machine.LCD_BACKLIGHT.High()

	return d
}
