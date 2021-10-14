package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"io/ioutil"
	"log"
	"time"

	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
	"tinygo.org/x/tinyfont/examples/initdisplay"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	display := initdisplay.InitDisplay()
	display.FillScreen(color.RGBA{0, 0, 0, 255})

	//bb, err := ioutil.ReadFile("./_font/NotoSansJP-Regular.otf")
	bb, err := ioutil.ReadFile("./_font/NotoSansMono-Regular.ttf")
	if err != nil {
		return err
	}

	ft, err := opentype.Parse(bb)
	if err != nil {
		return err
	}
	//fmt.Printf("%#v\n", ft)

	face, err := opentype.NewFace(ft, &opentype.FaceOptions{Size: 48, DPI: 72})
	if err != nil {
		return err
	}

	face2, err := opentype.NewFace(ft, &opentype.FaceOptions{Size: 12, DPI: 72})
	if err != nil {
		return err
	}
	//fmt.Printf("%#v\n", face)

	if false {
		dr, mask, maskp, advance, ok := face.Glyph(fixed.Point26_6{}, 'あ')
		fmt.Printf("mask: %#v\n", mask)
		fmt.Printf("dr: %#v\n", dr)
		fmt.Printf("maskp: %#v\n", maskp)
		fmt.Printf("advance: %#v\n", advance)
		fmt.Printf("ok: %#v\n", ok)
	}

	sfntBuf := &sfnt.Buffer{}
	type xx struct {
		Rune  rune
		Index uint16
	}
	indexes := []xx{}
	max := rune(0x32000 * 2)
	max = 0x10000
	max = 0xFF
	for r := rune(0); r < max; r++ {
		idx, err := ft.GlyphIndex(sfntBuf, r)
		if err != nil {
			return err
		}
		if idx != 0 {
			indexes = append(indexes, xx{
				Rune:  r,
				Index: uint16(idx),
			})
			//cnt++
			//fmt.Printf("%08X:%c ", r, r)
		}
	}

	for _, mode := range []int{1, 0} {
		maxw, maxh := 0, 0
		offset := 0
		if mode == 0 {
			fmt.Printf("const cNotoSans12pt = \"\" +\n")
		} else if mode == 1 {
			fmt.Printf("var NotoSans12pt = Font{\n")
			fmt.Printf("	Glyphs: []Glypher{\n")
		}
		for i, xxx := range indexes {
			//if xxx.Rune != '/' {
			//	continue
			//}
			//if i == 10 {
			//	continue
			//}
			text := fmt.Sprintf("%c", xxx.Rune)
			if false {
				fmt.Printf("%d %s : ", xxx.Index, text)
			}

			dr, img, _, adv, ok := face.Glyph(fixed.Point26_6{}, xxx.Rune)
			if !ok {
				continue
			}
			w := adv.Ceil()
			h := adv.Ceil()
			if w > maxw {
				maxw = w
			}
			if h > maxh {
				maxh = h
			}
			if false {
				fmt.Printf("%08X  %c  %5d   %d %d\n", xxx.Rune, xxx.Rune, i, w, h)
				fmt.Printf("    %#v\n", dr)
			}

			img2 := image.NewPaletted(img.Bounds(), palette.WebSafe)
			for y := 0; y < img2.Bounds().Dy(); y++ {
				for x := 0; x < img2.Bounds().Dx(); x++ {
					img2.Set(x, y, color.RGBA{0x00, 0x00, 0x00, 0xFF})
				}
			}
			draw.CatmullRom.Scale(img2, img.Bounds(), img, img.Bounds(), draw.Over, nil)

			img3 := image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X/4, img.Bounds().Max.Y/4))
			draw.CatmullRom.Scale(img3, img3.Bounds(), img2, img2.Bounds(), draw.Over, nil)

			//display.FillScreen(color.RGBA{0, 0, 0, 255})
			var imgX image.Image
			imgX = img3
			sz := 12
			if false {
				imgX = img2
				sz = 48
			} else {
				if false {
					dr, img, _, adv, ok = face2.Glyph(fixed.Point26_6{}, xxx.Rune)
				}
			}

			if false {
				for y := 0; y < sz*2; y++ {
					for x := 0; x < sz*3; x++ {
						display.Set(x, y, color.RGBA{0x00, 0x00, 0x00, 0xFF})
					}
				}

				for xy := 0; xy < sz*2; xy++ {
					display.Set(xy, sz, color.RGBA{0xFF, 0xFF, 0xFF, 0xFF})
					display.Set(sz, xy, color.RGBA{0xFF, 0xFF, 0xFF, 0xFF})
					display.Set(sz+adv.Ceil(), xy, color.RGBA{0x00, 0xFF, 0x00, 0xFF})
				}
			}

			// font
			fontBuf := []uint8{}
			for y := imgX.Bounds().Min.Y; y < imgX.Bounds().Max.Y; y++ {
				for x := imgX.Bounds().Min.X; x < imgX.Bounds().Max.X; x++ {
					z := imgX.At(x, y)
					r, g, b, a := z.RGBA()
					if r != g || r != b {
						fmt.Printf("%#v\n", z)
						fmt.Scanln()
					}
					z = color.GrayModel.Convert(z)
					if true {
						r, _, _, _ = z.RGBA()
						r &= 0xC000
						if r == 0xC000 {
							r = 0xFFFF
						} else if r == 0x0000 {
							r = 0x0000
						} else if r == 0x8000 {
							r = 0x8000
						} else if r == 0x4000 {
							r = 0x4000
						} else {
							z = color.RGBA{uint8(r >> 8), 0, 0, uint8(a)}
							z = color.RGBA{0, 0, 0, uint8(a)}
						}
						fontBuf = append(fontBuf, byte((r&0xC000)>>14))
						z = color.RGBA{uint8(r >> 8), uint8(r >> 8), uint8(r >> 8), uint8(a)}
					}
					if false {
						display.Set(x+sz+dr.Min.X, y+sz+dr.Min.Y, z)
					}
				}
			}

			if false {
				// debug
				for y := imgX.Bounds().Min.Y; y < imgX.Bounds().Max.Y; y++ {
					for x := imgX.Bounds().Min.X; x < imgX.Bounds().Max.X; x++ {
						z := fontBuf[x+y*imgX.Bounds().Max.X]
						//fmt.Printf("%02d:%X ", x+y*imgX.Bounds().Max.X, z)
						fmt.Printf("%X", z)
					}
					fmt.Printf("\n")
				}
			}

			for len(fontBuf)%4 != 0 {
				fontBuf = append(fontBuf, 0)
			}

			if mode == 0 {
				// 0     : length
				// 1 - 4 : Rune
				// 5     : Width
				// 6     : Height
				// 7     : XAdvance
				// 8     : XOffset
				// 9     : YOffset
				// 10 -  : bmp
				switch xxx.Rune {
				case 0x00, 0x0D:
					fmt.Printf("	/* '\\x%02X' */ ", xxx.Rune)
				default:
					fmt.Printf("	/* '%c' */ ", xxx.Rune)
				}
				fmt.Printf("\"\\x%02X\" + ", len(fontBuf)/4+9)
				fmt.Printf("\"\\x%02X\\x%02X\\x%02X\\x%02X\" + ", byte(xxx.Rune>>24), byte(xxx.Rune>>16), byte(xxx.Rune>>8), byte(xxx.Rune))
				//width := dr.Max.X - dr.Min.X
				width := imgX.Bounds().Max.X - imgX.Bounds().Min.X
				//height := dr.Max.Y - dr.Min.Y
				height := imgX.Bounds().Max.Y - imgX.Bounds().Min.Y
				fmt.Printf("\"\\x%02X\\x%02X\" + ", width, height)
				fmt.Printf("\"\\x%02X\" + ", (adv.Ceil()+0)/4)
				fmt.Printf("\"\\x%02X\\x%02X\" + ", int2byte((dr.Min.X+0)/4), int2byte((dr.Min.Y+0)/4))
				fmt.Printf("\"")
				for ii := 0; ii < len(fontBuf); ii += 4 {
					fmt.Printf("\\x%02X", (fontBuf[ii]<<6)+(fontBuf[ii+1]<<4)+(fontBuf[ii+2]<<2)+fontBuf[ii+3])
				}
				fmt.Printf("\" +\n")
			} else if mode == 1 {
				if xxx.Rune <= 0x20 {
					fmt.Printf("		NotosansGlyph{Index: 0x%04X, Rune: 0x%08X}, // \\x%02X\n", offset, xxx.Rune, xxx.Rune)
				} else {
					fmt.Printf("		NotosansGlyph{Index: 0x%04X, Rune: 0x%08X}, // %c\n", offset, xxx.Rune, xxx.Rune)
				}
				offset += len(fontBuf)/4 + 9 + 1
			}

			//for y := dr.Min.Y; y < dr.Max.Y; y++ {
			//	display.Set(dr.Min.X+sz, y+sz, color.RGBA{0xFF, 0x00, 0x00, 0xFF})
			//	display.Set(dr.Max.X+sz, y+sz, color.RGBA{0xFF, 0x00, 0x00, 0xFF})
			//}
			//for x := dr.Min.X; x < dr.Max.X; x++ {
			//	display.Set(x+sz, dr.Min.Y+sz, color.RGBA{0xFF, 0x00, 0x00, 0xFF})
			//	display.Set(x+sz, dr.Max.Y+sz, color.RGBA{0xFF, 0x00, 0x00, 0xFF})
			//}

			//fmt.Scanln()

			//b, _ := font.BoundString(face, text)
			//w := b.Max.X - b.Min.X + fixed.I(1)
			//h := b.Max.Y - b.Min.Y + fixed.I(1)
			//if w.Ceil() > maxw {
			//	maxw = w.Ceil()
			//}
			//if h.Ceil() > maxh {
			//	maxh = h.Ceil()
			//}
			//fmt.Printf("%d %d\n", w.Ceil(), h.Ceil())
		}
		if mode == 0 {
			fmt.Printf("	\"\"\n")
		} else if mode == 1 {
			fmt.Printf("	},\n")
			fmt.Printf("	YAdvance: 12,\n")
			fmt.Printf("}\n")
		}

		if false {
			fmt.Printf("max : %d %d\n", maxw, maxh)
		}

		if false {
			text := "hello あいうえお"
			//text := "hello"
			b, a := font.BoundString(face, text)
			w := b.Max.X - b.Min.X + fixed.I(1)
			h := b.Max.Y - b.Min.Y + fixed.I(1)

			//img := image.NewRGBA(image.Rect(0, 0, w.Ceil(), h.Ceil()))
			img := image.NewRGBA(image.Rect(0, 0, 320, 240))
			for y := 0; y < img.Rect.Max.Y; y++ {
				for x := 0; x < img.Rect.Max.X; x++ {
					img.Set(x, y, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF})
				}
			}

			dr := &font.Drawer{
				Dst:  img,
				Src:  image.Black,
				Face: face,
				Dot:  fixed.Point26_6{},
			}
			dr.Src = image.NewUniform(color.RGBA{0, 0, 0, 255})
			//dr.Face = basicfont.Face7x13
			if true {
				dr.Dot.X = (w - a) / 2
				dr.Dot.Y = h - b.Max.Y + 1000
			} else {
				dr.Dot.X = 0
				dr.Dot.Y = 0
			}
			dr.DrawString(text)

			buf := &bytes.Buffer{}
			err = png.Encode(buf, img)
			if err != nil {
				return err
			}

			err = ioutil.WriteFile("out.png", buf.Bytes(), 0666)
			if err != nil {
				return err
			}
		}
		fmt.Printf("\n")
	}
	time.Sleep(100 * time.Millisecond)

	return nil
}

func int2byte(x int) byte {
	if x > 0 {
		return byte(x)
	}

	ret := 0xFFFFFFFF + x
	return byte(ret)
}
