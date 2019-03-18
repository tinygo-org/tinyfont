package main

import (
	"machine"

	"image/color"

	"github.com/conejoninja/tinyfont"
	"github.com/tinygo-org/drivers/hub75"
)

var display hub75.Device
var font tinyfont.Font = tinyfont.FreeSerif9pt7b

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

	bytesAr := []byte("Hello")
	x := int16(4)
	for i := range bytesAr {
		glyph := font.Glyphs[bytesAr[i]-font.First]
		tinyfont.DrawChar(display, font, x, 12, bytesAr[i], colors[i%7])
		x += int16(glyph.XAdvance)

	}
	bytesAr = []byte("Gophers!")
	x = int16(-1)
	for i := range bytesAr {
		glyph := font.Glyphs[bytesAr[i]-font.First]
		tinyfont.DrawChar(display, font, x, 28, bytesAr[i], colors[(80-1-i)%7])
		x += int16(glyph.XAdvance)

	}
	for {
		display.Display()
	}
}
