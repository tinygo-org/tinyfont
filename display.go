package tinyfont

import (
	"image/color"

	"tinygo.org/x/drivers"
)

type RotatedDisplay struct {
	display  drivers.Displayer
	Rotation Rotation
	X        int16
	Y        int16
}

func NewRotatedDisplay(display drivers.Displayer, rotation Rotation, x, y int16) RotatedDisplay {
	return RotatedDisplay{
		display:  display,
		Rotation: rotation,
		X:        x,
		Y:        y,
	}
}

func (d RotatedDisplay) Size() (int16, int16) {
	x, y := d.display.Size()
	if d.Rotation == NO_ROTATION {
		return x, y
	} else if d.Rotation == ROTATION_90 {
		return y, x
	} else if d.Rotation == ROTATION_180 {
		return x, y
	} else {
		return y, x
	}
}

func (d RotatedDisplay) SetPixel(x, y int16, c color.RGBA) {
	if d.Rotation == NO_ROTATION {
		d.display.SetPixel(d.X+x, d.Y+y, c)
	} else if d.Rotation == ROTATION_90 {
		d.display.SetPixel(d.X-y, d.Y+x, c)
	} else if d.Rotation == ROTATION_180 {
		d.display.SetPixel(d.X-x, d.Y-y, c)
	} else {
		d.display.SetPixel(d.X+y, d.Y-x, c)
	}
}

func (d RotatedDisplay) Display() error {
	return d.display.Display()
}
