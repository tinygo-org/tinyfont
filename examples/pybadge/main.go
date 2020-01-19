package main

import (
	"github.com/conejoninja/tinyfont/gophers"
	"image/color"
	"machine"

	"github.com/conejoninja/tinyfont/freeserif"

	"github.com/conejoninja/tinyfont/freemono"
	"github.com/conejoninja/tinyfont/freesans"

	"github.com/conejoninja/tinyfont"
	"tinygo.org/x/drivers/st7735"
)

func main() {
	machine.SPI1.Configure(machine.SPIConfig{
		SCK:       machine.SPI1_SCK_PIN,
		MOSI:      machine.SPI1_MOSI_PIN,
		MISO:      machine.SPI1_MISO_PIN,
		Frequency: 8000000,
	})

	display := st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)
	display.Configure(st7735.Config{
		Rotation: st7735.ROTATION_90,
	})

	display.FillScreen(color.RGBA{255, 255, 255, 255})
	mycolors := make([]color.RGBA, 20)
	for k := 0; k < 20; k++ {
		mycolors[k] = getRainbowRGB(uint8(k * 14))
	}
	tinyfont.WriteLineColors(&display, &freesans.Bold18pt7b, 10, 35, []byte("HELLO"), mycolors)
	tinyfont.WriteLineColorsRotated(&display, &freemono.Bold18pt7b, 150, 100, []byte("Gophers"), mycolors[6:], tinyfont.ROTATION_180)
	tinyfont.WriteLineColorsRotated(&display, &freeserif.Bold9pt7b, 150, 90, []byte("TinyGo"), mycolors[12:], tinyfont.ROTATION_270)
	tinyfont.WriteLineColorsRotated(&display, &tinyfont.Org01, 10, 40, []byte("TinyGo"), mycolors[18:], tinyfont.ROTATION_90)

	tinyfont.WriteLineColors(&display, &gophers.Regular58pt, 18, 90, []byte("ABC"), mycolors[9:])


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
