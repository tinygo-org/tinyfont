package main

import (
	"machine"

	"image/color"

	"github.com/conejoninja/tinyfont"
	"github.com/tinygo-org/drivers/hub75"
)

var display Displayer

func main() {
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
		Mode:      0},
	)

	display = hub75.New(machine.SPI0, 11, 12, 6, 10, 18, 20)
	display.Configure(hub75.Config{
		Width:      64,
		Height:     32,
		RowPattern: 16,
		ColorDepth: 6,
	})

	colors := []color.RGBA{
		{255, 0, 0, 255},
		{255, 255, 0, 255},
		{0, 255, 0, 255},
		{0, 255, 255, 255},
		{0, 0, 255, 255},
		{255, 0, 255, 255},
		{255, 255, 255, 255},
	}

	display.ClearDisplay()
	display.SetBrightness(100)

	tinyfont.WriteLine(display, &tinyfont.TomThumb, 18, 12, []byte("Hello"), colors[0])
	tinyfont.WriteLineColors(display, &tinyfont.Org01, 12, 28, []byte("Gophers!"), colors)
	for {
		display.Display()
	}
}
