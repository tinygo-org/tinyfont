package main

import (
	"image/color"
	"time"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/examples/initdisplay"
	"tinygo.org/x/tinyfont/notoemoji"
)

func main() {
	display := initdisplay.InitDisplay()

	str := "\x23\x30\x31\x32\x33\x34\x35\x36\x37\n" +
		"\x38\x39\u00A9\u00AE\u203C\u2049\u20E3\u2122\u2139\n" +
		"\u2194\u2195\u2196\u2197\u2198\u2199\u21A9\u21AA\u231A\n" +
		"\u231B\u23E9\u23EA\u23EB\u23EC\u23F0\u23F3\u24C2\u25AA\n" +
		"\u25AB\u25B6\u25C0\u25CA\u25FB\u25FC\u25FD\u25FE\u2600\n" +
		"\u2601\u260E\u2611\u2614\u2615\u261D\u263A\u2648\u2649\n" +
		"\u264A\u264B\u264C\u264D\u264E\u264F\u2650\u2651\u2652\n" +
		"\u2653\u2660\u2663\u2665\u2666\u2668\u267B\u267F\u2693\n" +
		"\u26A0\u26A1\u26AA\u26AB\u26BD\u26BE\u26C4\u26C5\u26CE\n" +
		"\u26D4\u26EA\u26F2\u26F3\u26F5\u26FA\u26FD\u2702\u2705\n" +
		"\u2708\u2709\u270A\u270B\u270C\u270F\u2712\u2714\u2716\n" +
		"\u2728\u2733\u2734\u2744\u2747\u274C\u274E\u2753\u2754\n" +
		"\u2755\u2757\u2764\u2795\u2796\u2797\u27A1\u27B0\u27BF\n" +
		"\u2934\u2935\u2B05\u2B06\u2B07\u2B1B\u2B1C\u2B50\u2B55\n" +
		"\u3030\u303D\u3297\u3299\uFEFF"
	tinyfont.WriteLine(display, &notoemoji.NotoEmojiRegular20pt, 3, 0x16, str, color.RGBA{0, 0, 0, 255})

	for {
		time.Sleep(time.Hour)
	}
}
