package main

import (
	"image/color"
	"syscall/js"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/quanpixel8"
)

type canvasDisplay struct {
	width  int16
	height int16
	ctx    js.Value

	originX int
	originY int
	scale   int
}

func (d *canvasDisplay) Size() (int16, int16) {
	return d.width, d.height
}

func (d *canvasDisplay) SetPixel(x, y int16, c color.RGBA) {
	if c.A == 0 {
		return
	}

	px := d.originX + int(x)*d.scale
	py := d.originY + int(y)*d.scale

	if px < 0 || py < 0 || px >= int(d.width) || py >= int(d.height) {
		return
	}

	d.ctx.Set("fillStyle", "white")
	d.ctx.Call("fillRect", px, py, d.scale, d.scale)
}

func (d *canvasDisplay) Display() error {
	return nil
}

type probeDisplay struct {
	width  int16
	height int16
	hit    bool
}

func (d *probeDisplay) Size() (int16, int16) {
	return d.width, d.height
}

func (d *probeDisplay) SetPixel(x, y int16, c color.RGBA) {
	if c.A != 0 {
		d.hit = true
	}
}

func (d *probeDisplay) Display() error {
	return nil
}

func hasVisibleGlyph(font tinyfont.Fonter, r rune) bool {
	probe := &probeDisplay{width: 64, height: 64}
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	tinyfont.WriteLine(probe, font, 0, 32, string(r), white)
	return probe.hit
}

func countBlankGlyphs(font tinyfont.Fonter, runes []rune) int {
	count := 0
	for _, r := range runes {
		if !hasVisibleGlyph(font, r) {
			count++
		}
	}
	return count
}

func clear(ctx js.Value, width, height int) {
	ctx.Set("fillStyle", "black")
	ctx.Call("fillRect", 0, 0, width, height)
}

// drawItem draws a single packed row: rune code at the left edge,
// followed by the enlarged glyph. Blank glyphs still show their rune code.
func drawItem(ctx js.Value, display *canvasDisplay, font tinyfont.Fonter, r rune, x, y, itemW, itemH int) {
	ctx.Set("fillStyle", "#aaa")
	ctx.Set("font", "10px monospace")
	ctx.Call("fillText", codePoint(r), x+2, y+itemH-7)

	display.originX = x + 58
	display.originY = y + 3
	display.scale = 2

	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	tinyfont.WriteLine(display, font, 0, 8, string(r), white)

	ctx.Set("strokeStyle", "#222")
	ctx.Call("beginPath")
	ctx.Call("moveTo", float64(x), float64(y+itemH)-0.5)
	ctx.Call("lineTo", float64(x+itemW), float64(y+itemH)-0.5)
	ctx.Call("stroke")
}

func setText(id, value string) {
	js.Global().Get("document").Call("getElementById", id).Set("textContent", value)
}

func codePoint(r rune) string {
	const hex = "0123456789ABCDEF"

	if r <= 0xFFFF {
		buf := []byte{'U', '+', '0', '0', '0', '0'}
		v := uint32(r)
		for i := 5; i >= 2; i-- {
			buf[i] = hex[v&0xF]
			v >>= 4
		}
		return string(buf)
	}

	buf := []byte{'U', '+', '0', '0', '0', '0', '0', '0'}
	v := uint32(r)
	for i := 7; i >= 2; i-- {
		buf[i] = hex[v&0xF]
		v >>= 4
	}
	return string(buf)
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}

	neg := n < 0
	if neg {
		n = -n
	}

	buf := [21]byte{}
	i := len(buf)

	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}

	if neg {
		i--
		buf[i] = '-'
	}

	return string(buf[i:])
}

func newCanvas(doc, container js.Value, width, height int) js.Value {
	canvas := doc.Call("createElement", "canvas")
	canvas.Set("width", width)
	canvas.Set("height", height)
	canvas.Get("style").Set("display", "block")
	container.Call("appendChild", canvas)
	return canvas.Call("getContext", "2d")
}

func drawTestStrings(doc, parent js.Value, font tinyfont.Fonter, strs []string, scale int) {
	const gap = 16

	runeCellW := 8 * scale

	totalWidth := gap
	for _, s := range strs {
		totalWidth += len([]rune(s))*runeCellW + gap
	}
	height := 8 + 12*scale

	canvas := doc.Call("createElement", "canvas")
	canvas.Set("width", totalWidth)
	canvas.Set("height", height)

	style := canvas.Get("style")
	style.Set("display", "block")
	style.Set("position", "sticky")
	style.Set("top", "0")
	style.Set("zIndex", "10")
	style.Set("boxShadow", "0 2px 6px rgba(0,0,0,0.6)")

	parent.Call("appendChild", canvas)
	ctx := canvas.Call("getContext", "2d")

	clear(ctx, totalWidth, height)

	display := &canvasDisplay{
		width:  int16(totalWidth),
		height: int16(height),
		ctx:    ctx,
		scale:  scale,
	}

	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}

	x := gap
	for _, s := range strs {
		display.originX = x
		display.originY = 4
		tinyfont.WriteLine(display, font, 0, 8, s, white)

		x += len([]rune(s))*runeCellW + gap
	}
}

func main() {
	doc := js.Global().Get("document")
	container := doc.Call("getElementById", "viewport")

	font := &quanpixel8.QuanPixel8

	drawTestStrings(doc, container, font, []string{
		"あいうえお",
		"你好，中国",
		"你好，臺灣",
		"안녕하세요",
	}, 2)

	// allRunes is generated in runes.go from quanpixel8/quanpixel8.go.
	// This table intentionally displays every recorded rune, including blank glyphs.
	runes := allRunes
	blankGlyphs := countBlankGlyphs(font, runes)

	const itemW = 140
	const itemH = 22

	clientWidth := container.Get("clientWidth").Int()
	if clientWidth < itemW {
		clientWidth = itemW
	}

	cols := clientWidth / itemW
	if cols < 1 {
		cols = 1
	}
	canvasWidth := cols * itemW

	const maxChunkHeight = 8000

	totalRows := (len(runes) + cols - 1) / cols
	rowsPerChunk := maxChunkHeight / itemH
	if rowsPerChunk < 1 {
		rowsPerChunk = 1
	}

	totalHeight := totalRows * itemH

	display := &canvasDisplay{width: int16(canvasWidth), scale: 2}

	for chunkStart := 0; chunkStart < totalRows; chunkStart += rowsPerChunk {
		chunkRows := rowsPerChunk
		if chunkStart+chunkRows > totalRows {
			chunkRows = totalRows - chunkStart
		}
		chunkHeight := chunkRows * itemH

		ctx := newCanvas(doc, container, canvasWidth, chunkHeight)
		clear(ctx, canvasWidth, chunkHeight)

		display.ctx = ctx
		display.height = int16(chunkHeight)

		firstItem := chunkStart * cols
		lastItem := (chunkStart + chunkRows) * cols
		if lastItem > len(runes) {
			lastItem = len(runes)
		}

		for i := firstItem; i < lastItem; i++ {
			r := runes[i]
			row := i/cols - chunkStart
			col := i % cols

			x := col * itemW
			y := row * itemH

			drawItem(ctx, display, font, r, x, y, itemW, itemH)
		}
	}

	setText("count", "font entries: "+itoa(len(runes))+" / blank glyphs: "+itoa(blankGlyphs))
	setText("size", "canvas: "+itoa(canvasWidth)+" x "+itoa(totalHeight)+" (split into multiple canvases)")
	setText("note", "All runes present in QuanPixel8 are shown, including blank glyphs.")

	select {}
}
