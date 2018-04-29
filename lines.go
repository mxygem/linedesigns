package lines

import (
	"github.com/tidwall/pinhole"
)

// Design ...
type Design struct {
	LineWidth   float64
	ImageWidth  int
	ImageHeight int
	P           *pinhole.Pinhole
}

// New ...
func New(lineWidth float64, imageWidth, imageHeight int) *Design {
	return &Design{LineWidth: lineWidth, ImageWidth: imageWidth, ImageHeight: imageHeight, P: pinhole.New()}
}

// Rectangle creates a lined rectangle
func (d *Design) Rectangle(x1, y1, x2, y2, x3, y3, x4, y4, count float64) {
	left := d.dottedLine(x1, y1, x2, y2, count)
	bottom := d.dottedLine(x2, y2, x3, y3, count)
	right := d.dottedLine(x3, y3, x4, y4, count)
	top := d.dottedLine(x4, y4, x1, y1, count)

	d.connectDots(left, bottom)
	d.connectDots(bottom, right)
	d.connectDots(right, top)
	d.connectDots(top, left)
}

// Triangle creates a lined triangle
func (d *Design) Triangle(x1, y1, x2, y2, x3, y3, count float64) {
	l1 := d.dottedLine(x1, y1, x2, y2, count)
	l2 := d.dottedLine(x2, y2, x3, y3, count)
	l3 := d.dottedLine(x3, y3, x1, y1, count)

	d.connectDots(l1, l2)
	d.connectDots(l2, l3)
	d.connectDots(l3, l1)
}

// Save saves
func (d *Design) Save(fileName string) {
	opts := *pinhole.DefaultImageOptions
	opts.LineWidth = d.LineWidth
	d.P.SavePNG(fileName, d.ImageWidth, d.ImageHeight, &opts)
}
