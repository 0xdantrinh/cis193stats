package cis193stats

import (
	"fmt"

	"github.com/dantrinh/cis193stats"
	"math/rand"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
)

// randomPoints returns some random x, y points.
func randomPoints(n int, inpX data, inpY data) plotter.XYs {
	pts := make(plotter.XYs, inp.len())
	for i := range pts {
		/*if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}*/
		pts[i].X = inpX[i]
		pts[i].Y = inpY[i]
	}
	return pts
}

func plot(title string, inpX data, inpY data) {

	rand.Seed(int64(0))

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p,
		"First", randomPoints(15, inpX, inpY))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}
