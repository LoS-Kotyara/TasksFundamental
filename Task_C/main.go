package C

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"math"
	"strconv"
)

// 5 вариант
/*

Построить betta-проекцию с точностью до заданных k разрядов для функции f(x)
	при betta=2^{1 / n} (для n=2,3, ..., 100)

*/

func f(x int) int {
	return (x ^ 1) ^ (2 * (x & (1 + 2*x) & (3 + 4*x) & (7 + 8*x) & (15 + 16*x) & (31 + 32*x) & (63 + 64*x))) ^ (4 * (x*x + 5))
	// return (x ^ 1) ^ (2 * (x & (1 + 2*x) & (3 + 4*x) & (7 + 8*x) & (
	//	15 + 16*x) & (31 + 32*x) & (63 + 64*x))) ^ (4 * (x*x + 11))
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func toBeta(x int, k int, beta float64) float64 {
	reduction := reverseString(fmt.Sprintf("%0*b", k, x))

	result := 0.0

	for i, e := range reduction {
		v, _ := strconv.Atoi(string(e))
		result += float64(v) * math.Pow(beta, float64(i))
	}

	return result
}

func process(k int, n int) ([]float64, []float64) {

	beta := math.Pow(2.0, 1.0/float64(n))

	betaToK := math.Pow(beta, float64(k))

	var k2 float64 = float64(int(1) << k)
	xS := make([]float64, int(k2))
	fX := make([]float64, int(k2))

	for i := range xS {
		xS[i] = math.Mod(toBeta(i, k, beta)/betaToK, 1.0)
		fX[i] = math.Mod(toBeta(f(i), k, beta)/betaToK, 1.0)
	}

	return xS, fX
}

func Draw() *components.Page {
	page := components.NewPage()

	for i := 10; i < 20; i += 5 {
		for j := i; j < 20; j += 2 {
			xS, fX := process(i, j)

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
					Sprintf("C, k=%d, n=%d", i, j)}),
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
	}

	return page
}
