package linedesigns

import (
	"math"
)

func (d *Design) dottedLine(x1, y1, x2, y2, count float64) (points [][]float64) {
	for i := float64(0); i <= count; i++ {
		// the equation below returns a point that equates to the ratio of the total line. i.e.:
		// |---|---|---|---|---|---|---|
		// 0       2                  100
		x := x1 + (count-i)*(x2-x1)/count
		y := y1 + (count-i)*(y2-y1)/count
		points = append(points, []float64{x, y})
	}
	return
}

func (d *Design) dottedCircle(x, y, r, count float64) (points [][]float64) {
	for i := float64(0); i < count; i++ {
		deg := i * 360.0 / count
		X := x + (r * math.Cos(deg*(math.Pi/180)))
		Y := y + (r * math.Sin(deg*(math.Pi/180)))
		points = append(points, []float64{X, Y})
	}
	return
}

func (d *Design) connectPoints(c1, c2 [][]float64) {
	for i, c1p := range c1 {
		if i < len(c2)-1 {
			c2p := c2[i+1]
			d.P.DrawLine(c1p[0], c1p[1], 0, c2p[0], c2p[1], 0)
		}
	}
}

func (d *Design) continuous(n [][]float64, count float64) (points [][]float64) {
	for i, point := range n {
		if i < len(n)-1 {
			points = append(points, d.dottedLine(point[0], point[1], n[i+1][0], n[i+1][1], count)...)
		} else {
			points = append(points, d.dottedLine(point[0], point[1], n[0][0], n[0][1], count)...)
		}
	}
	return
}

func (d *Design) connectPointStream(points [][]float64, offset int, wrap bool) {
	if offset >= len(points) {
		offset = len(points) - 1
	}
	for i, point := range points {
		currentOffset := i + offset
		if currentOffset < len(points) {
			offsetPoints := points[currentOffset]
			d.P.DrawLine(point[0], point[1], 0, offsetPoints[0], offsetPoints[1], 0)
		} else if wrap {
			adjustedOffsetPoints := points[0+(offset)-(len(points)-i)]
			d.P.DrawLine(point[0], point[1], 0, adjustedOffsetPoints[0], adjustedOffsetPoints[1], 0)
		}
	}
}

func (d *Design) drawPoints(points [][]float64) {
	for _, p := range points {
		d.P.DrawDot(p[0], p[1], 0, d.DotSize)
	}
}
