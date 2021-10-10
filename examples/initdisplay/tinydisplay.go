//go:build !baremetal
// +build !baremetal

package initdisplay

import (
	"image/color"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/sago35/tinydisplay"
)

type TinyDisplay struct {
	*tinydisplay.Client
}

// InitDisplay initializes the display of each board.
func InitDisplay() *TinyDisplay {
	d, err := tinydisplay.NewClient("", 9812, 320, 240)
	if err != nil {
		log.Fatal(err)
	}
	d.FillScreen(color.RGBA{0, 0, 0, 255})
	time.Sleep(100 * time.Millisecond)
	d.FillScreen(color.RGBA{255, 255, 255, 255})

	return &TinyDisplay{
		Client: d,
	}
}

func (d *TinyDisplay) SaveAs(filename string) error {
	w, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer w.Close()

	return png.Encode(w, d)
}
