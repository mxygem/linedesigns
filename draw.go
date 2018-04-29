package lines

import (
	"fmt"

	"github.com/tidwall/pinhole"
)

// DottedLine ...
func (d *Design) DottedLine(x1, y1, x2, y2, count float64) [][]float64 {
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

// ConnectDots ...
func (d *Design) ConnectDots(c1, c2 [][]float64, p *pinhole.Pinhole) {
	fmt.Println(c1)
	fmt.Println(c2)
	for i, c1p := range c1 {
		fmt.Printf("i %d x: %.2f y: %.2f\n", i, c1p[0], c1p[1])
		if i < len(c2)-1 {
			c2p := c2[i+1]
			p.DrawLine(c1p[0], c1p[1], 0, c2p[0], c2p[1], 0)
		}
	}
}
