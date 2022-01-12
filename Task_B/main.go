package B

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"math"
	"sort"
)

// 5 Вариант
/*

Построить проекцию в единичном квадрате плоскости для функции f(
x) с точностью до заданных k=1, 2, 3, ... разрядов

f(x) = (x XOR 1) XOR (2 (x AND (1 + 2x) AND (3 + 4x) AND (7 + 8x) AND (
15 + 16x) AND (31 + 32x) AND (63 + 64x))) XOR (4(x^2 + 5))
*/

func f(x int) float64 {
	return float64((x ^ 1) ^ (2 * (x & (1 + 2*x) & (3 + 4*x) & (7 + 8*x) & (15 + 16*x) & (31 + 32*x) & (63 + 64*x))) ^ (4 * (x*x + 5)))
	// return (x ^ 1) ^ (2 * (x & (1 + 2*x) & (3 + 4*x) & (7 + 8*x) & (
	//	15 + 16*x) & (31 + 32*x) & (63 + 64*x))) ^ (4 * (x*x + 11))
}

func _(x int) float64 {
	return float64(x*x + 3)
}

func process(k int) ([]float64, []float64) {
	var k2 float64 = float64(int(1) << k)

	xS := make([]float64, int(k2))
	fX := make([]float64, int(k2))
	q := make([]float64, int(k2))

	for i := range xS {
		xS[i] = float64(i) / k2
		fX[i] = math.Mod(f(i), k2) / k2
		q[i] = f(i)
	}

	sort.Float64s(q)
	return xS, fX
}

func Draw() *components.Page {
	page := components.NewPage()

	for i := 15; i < 20; i++ {
		xS, fX := process(i)

		scatterData := make([]opts.ScatterData, len(fX))
		for i, v := range fX {
			scatterData[i] = opts.ScatterData{
				Value:      []float64{xS[i], v},
				SymbolSize: 3,
			}
		}

		scatter := charts.NewScatter()

		scatter.SetGlobalOptions(
			charts.WithTitleOpts(opts.Title{Title: fmt.
				Sprintf("B, k=%d", i)}),
			charts.WithXAxisOpts(opts.XAxis{
				Name:      "x",
				SplitLine: &opts.SplitLine{Show: false},
			}),
			charts.WithYAxisOpts(opts.YAxis{
				Name:      "f(x)",
				SplitLine: &opts.SplitLine{Show: false},
			}),
		)

		scatter.AddSeries("", scatterData)

		page.AddCharts(scatter)
	}

	return page
}
