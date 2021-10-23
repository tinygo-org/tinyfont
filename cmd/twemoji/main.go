package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"golang.org/x/image/draw"
)

func main() {
	err := run(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second)
}

type st struct {
	Rune  rune
	Bytes []byte
}

func run(dir string) error {
	os.MkdirAll(filepath.Dir("_out/const.go"), 0775)
	w, err := os.Create("_out/const.go")
	if err != nil {
		return err
	}
	defer w.Close()

	matches, err := filepath.Glob(fmt.Sprintf("%s/*.png", dir))
	//matches, err := glob(fmt.Sprintf("%s/*.png", dir))
	if err != nil {
		return err
	}

	dict := map[rune]st{}
	for i, match := range matches {
		if strings.Contains(match, "-") {
			continue
		}
		if i > 20 {
			//break
		}
		var r rune
		fmt.Sscanf(filepath.Base(match), "%X.png", &r)

		//render(match)
		buf, err := mktinyfontFromPng72x72(match)
		if err != nil {
			return err
		}
		dict[r] = st{
			Rune:  r,
			Bytes: buf,
		}

		//// fmt.Fprintf(w, "	/* %08X %c */ \"", r, r)
		//// for _, b := range buf {
		//// 	fmt.Fprintf(w, "\\x%02X", b)
		//// }
		//// fmt.Fprintf(w, "\" +\n")
		//// //fmt.Fprintf(w, "%#v\n", buf)
	}
	//fmt.Fprintf(w, "	\"\"\n")

	keys := []rune{}
	for k := range dict {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	fmt.Fprintf(w, "package twemoji\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "var TinyFont = sTinyFont{}\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "type sTinyFont struct {\n")
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "const mTinyFont = \"\" +\n")
	ofs := 0
	for _, k := range keys {
		v := dict[k]
		fmt.Fprintf(w, "	\"")
		fmt.Fprintf(w, "\\x%02X\\x%02X\\x%02X", byte((v.Rune&0x00FF0000)>>16), byte((v.Rune&0x0000FF00)>>8), byte(v.Rune))
		fmt.Fprintf(w, "\" + \"")
		fmt.Fprintf(w, "\\x%02X\\x%02X\\x%02X", byte((ofs&0x00FF0000)>>16), byte((ofs&0x0000FF00)>>8), byte(ofs))
		fmt.Fprintf(w, "\" + ")
		fmt.Fprintf(w, "/* %c */\n", v.Rune)

		ofs += len(v.Bytes)
	}
	fmt.Fprintf(w, "	\"\"\n")
	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "const cTinyFont = \"\" +\n")
	for _, k := range keys {
		v := dict[k]

		fmt.Fprintf(w, "	/* %08X %c */ \"", v.Rune, v.Rune)
		for _, b := range v.Bytes {
			fmt.Fprintf(w, "\\x%02X", b)
		}
		fmt.Fprintf(w, "\" +\n")
	}
	fmt.Fprintf(w, "	\"\"\n")

	return nil
}

func glob(target string) ([]string, error) {
	matches, err := filepath.Glob(target)
	if err != nil {
		return nil, err
	}

	sort.Slice(matches, func(i, j int) bool {
		ii := uint32(0)
		jj := uint32(0)
		fmt.Sscanf(filepath.Base(matches[i]), "%X.png", &ii)
		fmt.Sscanf(filepath.Base(matches[j]), "%X.png", &jj)

		return ii < jj
	})

	return matches, nil
}

func mktinyfontFromPng72x72(file string) ([]byte, error) {
	r, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	img, err := png.Decode(r)
	if err != nil {
		return nil, err
	}

	img2 := image.NewPaletted(img.Bounds(), palette.WebSafe)
	for y := 0; y < img2.Bounds().Dy(); y++ {
		for x := 0; x < img2.Bounds().Dx(); x++ {
			img2.Set(x, y, color.RGBA{0xFF, 0xFF, 0xFF, 0xFF})
		}
	}
	draw.CatmullRom.Scale(img2, img.Bounds(), img, img.Bounds(), draw.Over, nil)

	img3 := image.NewRGBA(image.Rect(0, 0, 12, 12))
	draw.CatmullRom.Scale(img3, img3.Bounds(), img2, img2.Bounds(), draw.Over, nil)

	ret := make([]byte, 0, 12*12*2)
	for y := 0; y < img3.Bounds().Dy(); y++ {
		for x := 0; x < img3.Bounds().Dx(); x++ {
			z := RGBATo565(img3.At(x, y))
			ret = append(ret, byte(z>>8), byte(z))
		}
	}

	return ret, nil
}

func RGBATo565(c color.Color) uint16 {
	r, g, b, _ := c.RGBA()
	return uint16((r & 0xF800) +
		((g & 0xFC00) >> 5) +
		((b & 0xF800) >> 11))
}

func RGBATo55a5(c color.Color) uint16 {
	r, g, b, a := c.RGBA()
	return uint16((r & 0xF800) +
		((g & 0xF800) >> 5) +
		((a & 0x8000) >> 10) +
		((b & 0xF800) >> 11))
}
