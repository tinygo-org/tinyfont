package main

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/ili9341"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/freeserif"
	"tinygo.org/x/tinyfont/gophers"
)

func main() {
	display := ili9341.NewParallel(
		machine.LCD_DATA0,
		machine.TFT_WR,
		machine.TFT_DC,
		machine.TFT_CS,
		machine.TFT_RESET,
		machine.TFT_RD,
	)
	display.Configure(ili9341.Config{})
	display.FillScreen(color.RGBA{255, 255, 255, 255})

	machine.TFT_BACKLIGHT.Configure(machine.PinConfig{machine.PinOutput})
	machine.TFT_BACKLIGHT.High()

	mycolors := make([]color.RGBA, 20)
	for k := 0; k < 20; k++ {
		mycolors[k] = getRainbowRGB(uint8(k * 14))
	}
	tinyfont.WriteLineColors(display, &freesans.Bold18pt7b, 10, 35, "HELLO", mycolors)
	tinyfont.WriteLineColorsRotated(display, &freemono.Bold18pt7b, 150, 100, "Gophers", mycolors[6:], tinyfont.ROTATION_180)
	tinyfont.WriteLineColorsRotated(display, &freeserif.Bold9pt7b, 150, 90, "TinyGo", mycolors[12:], tinyfont.ROTATION_270)
	tinyfont.WriteLineColorsRotated(display, &tinyfont.Org01, 10, 40, "TinyGo", mycolors[18:], tinyfont.ROTATION_90)

	tinyfont.WriteLineColors(display, &gophers.Regular58pt, 18, 90, "ABC", mycolors[9:])
}

func getRainbowRGB(i uint8) color.RGBA {
	if i < 85 {
		return color.RGBA{i * 3, 255 - i*3, 0, 255}
	} else if i < 170 {
		i -= 85
		return color.RGBA{255 - i*3, 0, i * 3, 255}
	}
	i -= 170
	return color.RGBA{0, i * 3, 255 - i*3, 255}
}
