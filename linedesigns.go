package linedesigns

import (
	"github.com/tidwall/pinhole"
)

// Design ...
type Design struct {
	LineWidth   float64
	DotSize     float64
	ImageWidth  int
	ImageHeight int
	P           *pinhole.Pinhole
}

// New sets up
func New(lineWidth, dotSize float64, imageWidth, imageHeight int) *Design {
	return &Design{LineWidth: lineWidth, DotSize: dotSize, ImageWidth: imageWidth, ImageHeight: imageHeight, P: pinhole.New()}
}

// LineDotted creates a dotted line segment
func (d *Design) LineDotted(x1, y1, x2, y2, count float64) {
	d.drawPoints(d.dottedLine(x1, y1, x2, y2, count))
}

// AngleDotted creates two dotted line segments
func (d *Design) AngleDotted(x1, y1, x2, y2, x3, y3, count float64) {
	d.drawPoints(d.dottedLine(x1, y1, x2, y2, count))
	d.drawPoints(d.dottedLine(x2, y2, x3, y3, count))
}

// TriangleDotted creates a dotted triangle
func (d *Design) TriangleDotted(x1, y1, x2, y2, x3, y3, count float64) {
	points := [][]float64{}
	points = append(points, d.dottedLine(x1, y1, x2, y2, count)...)
	points = append(points, d.dottedLine(x2, y2, x3, y3, count)...)
	points = append(points, d.dottedLine(x3, y3, x1, y1, count)...)

	d.drawPoints(points)
}

// RectangleDotted creates a dotted rectangle
func (d *Design) RectangleDotted(x1, y1, x2, y2, x3, y3, x4, y4, count float64) {
	points := [][]float64{}
	points = append(points, d.dottedLine(x1, y1, x2, y2, count)...)
	points = append(points, d.dottedLine(x2, y2, x3, y3, count)...)
	points = append(points, d.dottedLine(x3, y3, x4, y4, count)...)
	points = append(points, d.dottedLine(x4, y4, x1, y1, count)...)

	d.drawPoints(points)
}

// CircleDotted creates a dotted circle
func (d *Design) CircleDotted(x, y, r, count float64) {
	d.drawPoints(d.dottedCircle(x, y, r, count))
}

// FreeformDotted takes >3 number of coordinates and returns those line segments as dots
func (d *Design) FreeformDotted(n [][]float64, count float64) {
	d.drawPoints(d.continuous(n, count))
}

// Angle creates a lined angle from three points
func (d *Design) Angle(x1, y1, x2, y2, x3, y3, count float64) {
	d.connectPoints(d.dottedLine(x1, y1, x2, y2, count), d.dottedLine(x2, y2, x3, y3, count))
}

// Triangle creates a lined triangle
func (d *Design) Triangle(x1, y1, x2, y2, x3, y3, count float64) {
	l1 := d.dottedLine(x1, y1, x2, y2, count)
	l2 := d.dottedLine(x2, y2, x3, y3, count)
	l3 := d.dottedLine(x3, y3, x1, y1, count)

	d.connectPoints(l1, l2)
	d.connectPoints(l2, l3)
	d.connectPoints(l3, l1)
}

// Rectangle creates a lined rectangle from four points
func (d *Design) Rectangle(x1, y1, x2, y2, x3, y3, x4, y4, count float64) {
	left := d.dottedLine(x1, y1, x2, y2, count)
	bottom := d.dottedLine(x2, y2, x3, y3, count)
	right := d.dottedLine(x3, y3, x4, y4, count)
	top := d.dottedLine(x4, y4, x1, y1, count)

	d.connectPoints(left, bottom)
	d.connectPoints(bottom, right)
	d.connectPoints(right, top)
	d.connectPoints(top, left)
}

// Circle creates a lined circle
func (d *Design) Circle(x, y, r, count, offset float64, wrap bool) {
	d.connectPointStream(d.dottedCircle(x, y, r, count), int(offset), wrap)
}

// Freeform takes >3 number of coordinates and connects them together.
func (d *Design) Freeform(n [][]float64, count, offset float64, wrap bool) {
	d.connectPointStream(d.continuous(n, count), int(offset), wrap)
}

// Save saves
func (d *Design) Save(fileName string) {
	opts := *pinhole.DefaultImageOptions
	opts.LineWidth = d.LineWidth
	d.P.SavePNG(fileName, d.ImageWidth, d.ImageHeight, &opts)
}
