package lines

import (
	"github.com/tidwall/pinhole"
)

// Design ...
type Design struct {
	LineWidth   float64
	ImageWidth  int
	ImageHeight int
}

// New ...
func New(lineWidth float64, imageWidth, imageHeight int) *Design {
	return &Design{LineWidth: lineWidth, ImageWidth: imageWidth, ImageHeight: imageHeight}
}

// Rectangle creates a lined rectangle
func (d *Design) Rectangle(x1, y1, x2, y2, x3, y3, x4, y4, count float64) {
	p := pinhole.New()
	left := d.DottedLine(x1, y1, x2, y2, count)
	bottom := d.DottedLine(x2, y2, x3, y3, count)
	right := d.DottedLine(x3, y3, x4, y4, count)
	top := d.DottedLine(x4, y4, x1, y1, count)

	d.ConnectDots(left, bottom, p)
	d.ConnectDots(bottom, right, p)
	d.ConnectDots(right, top, p)
	d.ConnectDots(top, left, p)

	d.save(p)
}

// Triangle creates a lined triangle
func (d *Design) Triangle(x1, y1, x2, y2, x3, y3, count float64) {
	p := pinhole.New()
	l1 := d.DottedLine(x1, y1, x2, y2, count)
	l2 := d.DottedLine(x2, y2, x3, y3, count)
	l3 := d.DottedLine(x3, y3, x1, y1, count)

	d.ConnectDots(l1, l2, p)
	d.ConnectDots(l2, l3, p)
	d.ConnectDots(l3, l1, p)

	d.save(p)
}

func (d *Design) save(p *pinhole.Pinhole) {
	opts := *pinhole.DefaultImageOptions
	opts.LineWidth = d.LineWidth
	p.SavePNG("dotline.png", d.ImageWidth, d.ImageHeight, &opts)
}
