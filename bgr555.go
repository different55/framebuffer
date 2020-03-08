// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package framebuffer

import (
	"image"
	"image/color"
)

type BGR555 struct {
	Pix    []byte
	Rect   image.Rectangle
	Stride int
}

func (i *BGR555) Bounds() image.Rectangle { return i.Rect }
func (i *BGR555) ColorModel() color.Model { return RGB555Model }

func (i *BGR555) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(i.Rect)) {
		return RGBColor{}
	}

	pix := i.Pix[i.PixOffset(x, y):]
	clr := uint16(pix[0])<<8 | uint16(pix[1])

	return RGBColor{
		uint8(clr) & mask5,
		uint8(clr>>5) & mask5,
		uint8(clr>>10) & mask5,
	}
}

func (i *BGR555) Set(x, y int, c color.Color) {
	i.SetRGB(x, y, RGB555Model.Convert(c).(RGBColor))
}

func (i *BGR555) SetRGB(x, y int, c RGBColor) {
	if !(image.Point{x, y}.In(i.Rect)) {
		return
	}

	n := i.PixOffset(x, y)
	pix := i.Pix[n:]
	clr := uint16(c.B<<10) | uint16(c.G<<5) | uint16(c.R)

	pix[0] = uint8(clr)
	pix[1] = uint8(clr >> 8)
}

func (i *BGR555) PixOffset(x, y int) int {
	return (y-i.Rect.Min.Y)*i.Stride + (x-i.Rect.Min.X)*2
}
