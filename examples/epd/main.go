package main

import (
	"machine"

	"image/color"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
	"tinygo.org/x/drivers/waveshare-epd/epd2in13x"
)

var display epd2in13x.Device

func main() {
	machine.SPI0.Configure(machine.SPIConfig{
		SCK:       machine.SPI0_SCK_PIN,
		MOSI:      machine.SPI0_MOSI_PIN,
		MISO:      machine.SPI0_MISO_PIN,
		Frequency: 8000000,
		Mode:      0,
	})

	display = epd2in13x.New(machine.SPI0, machine.P6, machine.P7, machine.P8, machine.P9)
	display.Configure(epd2in13x.Config{})

	white := color.RGBA{0, 0, 0, 255}
	yellow := color.RGBA{255, 0, 0, 255}
	black := color.RGBA{1, 1, 1, 255}

	display.ClearBuffer()
	display.ClearDisplay()

	tinyfont.WriteLineRotated(&display, &freemono.Bold9pt7b, 15, 20, []byte("Hello"), black, tinyfont.NO_ROTATION)
	showRect(0, 22, 52, 20, black)
	showRect(52, 22, 52, 20, yellow)
	tinyfont.WriteLineRotated(&display, &freemono.Bold9pt7b, 85, 26, []byte("World!"), white, tinyfont.ROTATION_180)

	tinyfont.WriteLineRotated(&display, &freemono.Bold9pt7b, 55, 60, []byte("@tinyGolang"), yellow, tinyfont.ROTATION_90)
	tinyfont.WriteLineColorsRotated(&display, &freemono.Bold9pt7b, 45, 180, []byte("tinyfont"), []color.RGBA{yellow, black}, tinyfont.ROTATION_270)

	display.Display()
	display.WaitUntilIdle()
	println("You could remove power now")
}

func showRect(x int16, y int16, w int16, h int16, c color.RGBA) {
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			display.SetPixel(i, j, c)
		}
	}
}
