package main

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"log"

	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	bb, err := ioutil.ReadFile("./_font/NotoSansJP-Regular.otf")
	//bb, err := ioutil.ReadFile("./_font/NotoSansMono-Regular.ttf")
	if err != nil {
		return err
	}

	ft, err := opentype.Parse(bb)
	if err != nil {
		return err
	}
	//fmt.Printf("%#v\n", ft)

	face, err := opentype.NewFace(ft, &opentype.FaceOptions{Size: 12, DPI: 72})
	if err != nil {
		return err
	}

	sfntBuf := &sfnt.Buffer{}
	type xx struct {
		Rune  rune
		Index uint16
	}
	indexes := []xx{}
	max := rune(0x32000 * 2)
	max = 0x10000
	//max = 0xFF
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
		for _, xxx := range indexes {
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

			var imgX image.Image
			imgX = img

			// font
			fontBuf := []uint8{}
			for y := imgX.Bounds().Min.Y; y < imgX.Bounds().Max.Y; y++ {
				for x := imgX.Bounds().Min.X; x < imgX.Bounds().Max.X; x++ {
					z := imgX.At(x, y)
					r, g, b, a := z.RGBA()
					if r != g || r != b {
						fmt.Printf("rgb errror : %#v\n", z)
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
				fmt.Printf("\"\\x%02X\" + ", len(fontBuf)+9)
				fmt.Printf("\"\\x%02X\\x%02X\\x%02X\\x%02X\" + ", byte(xxx.Rune>>24), byte(xxx.Rune>>16), byte(xxx.Rune>>8), byte(xxx.Rune))
				//width := dr.Max.X - dr.Min.X
				width := imgX.Bounds().Max.X - imgX.Bounds().Min.X
				//height := dr.Max.Y - dr.Min.Y
				height := imgX.Bounds().Max.Y - imgX.Bounds().Min.Y
				fmt.Printf("\"\\x%02X\\x%02X\" + ", width, height)
				fmt.Printf("\"\\x%02X\" + ", (adv.Ceil() + 0))
				fmt.Printf("\"\\x%02X\\x%02X\" + ", int2byte(dr.Min.X), int2byte(dr.Min.Y))
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

		}
		if mode == 0 {
			fmt.Printf("	\"\"\n")
		} else if mode == 1 {
			fmt.Printf("	},\n")
			fmt.Printf("	YAdvance: 12,\n")
			fmt.Printf("}\n")
		}

		fmt.Printf("\n")
	}

	return nil
}

func int2byte(x int) byte {
	if x > 0 {
		return byte(x)
	}

	ret := 0xFFFFFFFF + x
	return byte(ret)
}
