package main

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/ili9341"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/notosans"
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

	str := "https://tour.golang.org/welcome/2\n" +
		"\n" +
		"Brazilian Portuguese\n    — Português do Brasil\n" +
		"Catalan — Català\n" +
		"Simplified Chinese — 中文（简体）\n" +
		"Traditional Chinese — 中文（繁體）\n" +
		"Czech — Česky\n" +
		"French — Français\n" +
		"German — Deutsch\n" +
		"Hebrew — עִבְרִית\n" +
		"Indonesian — Bahasa Indonesia\n" +
		"Italian — Italiano\n" +
		"Japanese — 日本語\n" +
		"Korean — 한국어\n" +
		"Romanian — Română\n" +
		"Russian - Русский\n" +
		"Spanish — Español\n" +
		"Thai - ภาษาไทย\n" +
		"Turkish - Türkçe\n" +
		"Ukrainian — Українська\n" +
		"Uzbek — Ўзбекча\n"

	tinyfont.WriteLine(display, &notosans.Notosans12pt, 3, 0x16, str, color.RGBA{0, 0, 0, 255})
}
