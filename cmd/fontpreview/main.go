// fontpreview renders a sample of each tinyfont font to a PNG in the images/ directory.
// Run from the repo root: go run ./cmd/fontpreview
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/freeserif"
	"tinygo.org/x/tinyfont/gophers"
	"tinygo.org/x/tinyfont/notoemoji"
	"tinygo.org/x/tinyfont/notosans"
	"tinygo.org/x/tinyfont/proggy"
	"tinygo.org/x/tinyfont/shnm"
)

// imgDisplay is an in-memory display backed by an image.RGBA.
type imgDisplay struct {
	img *image.RGBA
}

func (d *imgDisplay) SetPixel(x, y int16, c color.RGBA) {
	b := d.img.Bounds()
	if int(x) >= 0 && int(x) < b.Max.X && int(y) >= 0 && int(y) < b.Max.Y {
		d.img.SetRGBA(int(x), int(y), c)
	}
}

func (d *imgDisplay) Size() (int16, int16) {
	b := d.img.Bounds()
	return int16(b.Max.X), int16(b.Max.Y)
}

func (d *imgDisplay) Display() error { return nil }

// measureWidth returns the pixel width of text rendered in font.
func measureWidth(font tinyfont.Fonter, text string) int {
	w := 0
	for _, r := range text {
		g := font.GetGlyph(r)
		w += int(g.Info().XAdvance)
	}
	return w
}

// scaleImage returns a nearest-neighbour upscaled copy of src.
func scaleImage(src *image.RGBA, scale int) *image.RGBA {
	b := src.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Max.X*scale, b.Max.Y*scale))
	for y := 0; y < b.Max.Y; y++ {
		for x := 0; x < b.Max.X; x++ {
			c := src.RGBAAt(x, y)
			for dy := 0; dy < scale; dy++ {
				for dx := 0; dx < scale; dx++ {
					dst.SetRGBA(x*scale+dx, y*scale+dy, c)
				}
			}
		}
	}
	return dst
}

// renderFontRows renders multiple rows of text in font and saves a PNG to outDir/name.png.
// Small fonts are scaled up so they remain visible at a reasonable size.
func renderFontRows(name string, font tinyfont.Fonter, rows []string, outDir string, extraBottomPadding int) error {
	const padding = 4
	const fgVal = 0x00
	const bg = 0xFF

	yAdv := int(font.GetYAdvance())
	if yAdv == 0 {
		yAdv = 16
	}

	// Calculate max width and total height.
	maxW := 0
	for _, row := range rows {
		w := measureWidth(font, row)
		if w > maxW {
			maxW = w
		}
	}

	w := maxW + padding*2
	h := len(rows)*yAdv + padding*2 + extraBottomPadding

	img := image.NewRGBA(image.Rect(0, 0, w, h))
	// White background.
	white := color.RGBA{bg, bg, bg, 0xFF}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.SetRGBA(x, y, white)
		}
	}

	d := &imgDisplay{img}
	fgColor := color.RGBA{fgVal, fgVal, fgVal, 0xFF}
	for i, row := range rows {
		y := int16(padding + i*yAdv + yAdv)
		tinyfont.WriteLine(d, font, int16(padding), y, row, fgColor)
	}

	// Scale up small fonts so the image is at least ~48px tall.
	scale := 1
	switch {
	case h <= 16:
		scale = 5
	case h <= 24:
		scale = 4
	case h <= 48:
		scale = 2
	}
	if scale > 1 {
		img = scaleImage(img, scale)
	}

	path := filepath.Join(outDir, name+".png")
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, img)
}

func main() {
	const outDir = "images"

	asciiRows := []string{
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"abcdefghijklmnopqrstuvwxyz",
		"0123456789",
		"!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~",
	}

	// Gophers font only shows uppercase.
	gophersRows := []string{
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}

	// A selection of emoji known to exist in the 12pt NotoEmoji font.
	emojiRows := []string{
		"☀♠♥♦⭐❤✅✨🎉🎊🔥💡",
	}

	// Japanese hiragana characters for shnm font.
	japaneseRows := []string{
		"あいうえおかきくけこ",
		"さしすせそたちつてと",
		"なにぬねのはひふへほ",
	}

	type fontEntry struct {
		name               string
		font               tinyfont.Fonter
		rows               []string
		extraBottomPadding int
	}

	fonts := []fontEntry{
		{"freemono", &freemono.Regular12pt7b, asciiRows, 8},
		{"freesans", &freesans.Regular12pt7b, asciiRows, 8},
		{"freeserif", &freeserif.Regular12pt7b, asciiRows, 8},
		{"gophers", &gophers.Regular32pt, gophersRows, 0},
		{"notoemoji", &notoemoji.NotoEmojiRegular12pt, emojiRows, 0},
		{"notosans", &notosans.Notosans12pt, asciiRows, 0},
		{"proggy", &proggy.TinySZ8pt7b, asciiRows, 0},
		{"shnm", &shnm.Shnmk12, japaneseRows, 0},
		{"org_01", &tinyfont.Org01, asciiRows, 4},
		{"picopixel", &tinyfont.Picopixel, asciiRows, 4},
		{"tiny3x3a2pt7b", &tinyfont.Tiny3x3a2pt7b, asciiRows, 0},
		{"tomthumb", &tinyfont.TomThumb, asciiRows, 4},
	}

	for _, fe := range fonts {
		if err := renderFontRows(fe.name, fe.font, fe.rows, outDir, fe.extraBottomPadding); err != nil {
			fmt.Fprintf(os.Stderr, "error rendering %s: %v\n", fe.name, err)
		} else {
			fmt.Printf("rendered %s → %s/%s.png\n", fe.name, outDir, fe.name)
		}
	}
}
