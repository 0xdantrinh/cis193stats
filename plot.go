package cis193stats

import (
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
)

// randomPoints returns some random x, y points.
func RandomPoints(n int, inpX data, inpY data) plotter.XYs {
	l1 := inpX.Len()
	l2 := inpY.Len()

	if l1 == 0 || l2 == 0 {
		return nil
	}

	if l1 != l2 {
		return nil
	}

	pts := make(plotter.XYs, l1)
	for i := range pts {
		pts[i].X = inpX[i]
		pts[i].Y = inpY[i]
	}
	return pts
}

func PlotNow(title string, inpX data, inpY data) {

	//rand.Seed(int64(0))

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = title
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p,
		"First", RandomPoints(15, inpX, inpY))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}
