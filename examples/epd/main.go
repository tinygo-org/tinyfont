package main

import (
	"machine"
	"time"

	"image/color"

	"tinygo.org/x/drivers/waveshare-epd/epd2in13x"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

var display epd2in13x.Device

func main() {
	machine.SPI0.Configure(machine.SPIConfig{
		SCK:       machine.SPI0_SCK_PIN,
		SDO:       machine.SPI0_SDO_PIN,
		SDI:       machine.SPI0_SDI_PIN,
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

	tinyfont.WriteLineRotated(&display, &freemono.Bold9pt7b, 15, 20, "Hello", black, tinyfont.NO_ROTATION)
	showRect(0, 22, 52, 20, black)
	showRect(52, 22, 52, 20, yellow)
	tinyfont.WriteLineRotated(&display, &freemono.Bold9pt7b, 85, 26, "World!", white, tinyfont.ROTATION_180)

	tinyfont.WriteLineRotated(&display, &freemono.Bold9pt7b, 55, 60, "@tinyGolang", yellow, tinyfont.ROTATION_90)
	tinyfont.WriteLineColorsRotated(&display, &freemono.Bold9pt7b, 45, 180, "tinyfont", []color.RGBA{yellow, black}, tinyfont.ROTATION_270)

	display.Display()
	display.WaitUntilIdle()
	println("You could remove power now")

	for {
		time.Sleep(time.Hour)
	}
}

func showRect(x int16, y int16, w int16, h int16, c color.RGBA) {
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			display.SetPixel(i, j, c)
		}
	}
}
