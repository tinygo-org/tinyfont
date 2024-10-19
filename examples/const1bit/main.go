package main

import (
	"image/color"
	"time"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/examples/initdisplay"
	"tinygo.org/x/tinyfont/shnmk12"
)

func main() {
	display := initdisplay.InitDisplay()

	str := "As an example of displaying Japanese\nfonts, I will try displaying text from\n`The Tale of the Heike` and\n`Essays in Idleness (Tsurezuregusa)`.\n\n" +
		"[平家物語] (The Tale of the Heike)\n" +
		"祇園精舎の鐘の声、諸行無常の響きあり。\n沙羅双樹の花の色、盛者必衰の理をあらは\nす。おごれる人も久しからず、ただ春の夜\nの夢のごとし。\n\n" +
		"[徒然草] (Essays in Idleness)\n" +
		"つれづれなるままに、日ぐらし硯にむかひ\nて、心にうつりゆくよしなしごとを、そこ\nはかとなく書きつくれば、あやしうこそも\nのぐるほしけれ。"
	tinyfont.WriteLine(display, &shnmk12.Shnmk12, 3, 0x16, str, color.RGBA{0, 0, 0, 255})

	for {
		time.Sleep(time.Hour)
	}
}
