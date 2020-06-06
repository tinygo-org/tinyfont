// +build pybadge

package main

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/st7735"
)

func initDisplay() drivers.Displayer {
	machine.SPI1.Configure(machine.SPIConfig{
		SCK:       machine.SPI1_SCK_PIN,
		MOSI:      machine.SPI1_MOSI_PIN,
		MISO:      machine.SPI1_MISO_PIN,
		Frequency: 8000000,
	})

	d := st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)
	d.Configure(st7735.Config{
		Rotation: st7735.ROTATION_90,
	})

	d.FillScreen(color.RGBA{255, 255, 255, 255})

	return &d
}
