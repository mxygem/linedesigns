package linedesigns

func (d *Design) dottedLine(x1, y1, x2, y2, count float64) [][]float64 {
	points := [][]float64{}
	i := 0.0
	for i <= count {
		x := x1 + (count-i)*(x2-x1)/count
		y := y1 + (count-i)*(y2-y1)/count
		points = append(points, []float64{x, y})
		i++
	}

	return points
}

func (d *Design) connectDots(c1, c2 [][]float64) {
	for i, c1p := range c1 {
		if i < len(c2)-1 {
			c2p := c2[i+1]
			d.P.DrawLine(c1p[0], c1p[1], 0, c2p[0], c2p[1], 0)
		}
	}
}

func (d *Design) drawDots(dots [][]float64) {
	for _, pos := range dots {
		d.P.DrawDot(pos[0], pos[1], 0, 0.02)
	}
}
