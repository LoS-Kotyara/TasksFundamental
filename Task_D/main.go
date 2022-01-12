package D

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"math"
	"strconv"
)

/*

Построить последовательность ван дер Корпута с помощью отображения Какутани-фон Неймана. Построить
	отображение на Z2, индуцированное отображением Какутани-фон Неймана с помощью отображения Монна.
*/

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func kakutaniVonNeumann(x float64) float64 {
	b := 2.0
	k := 0.0

	for !(x >= 1.0-1.0/math.Pow(b, k) && x < 1.0-1.0/math.Pow(b, k+1)) {
		k = k + 1
	}
	return x - 1 + 1.0/math.Pow(b, k) + 1.0/math.Pow(b, k+1)
}

func monna(x int) float64 {
	binary := reverseString(fmt.Sprintf("%0b", x))

	sum := 0.0

	for i, e := range binary {
		v, _ := strconv.Atoi(string(e))
		sum += float64(v) * math.Pow(2, float64(-(i+1)))
	}
	return sum
}

func vanDerCorput(n int) ([]float64, []float64, []int, []float64) {
	kakutaniX := make([]float64, n)
	monnaX := make([]int, n)

	kakutaniY := make([]float64, n)
	monnaY := make([]float64, n)

	x := 0.0
	for i := 0; i < n; i++ {
		kakutaniX[i] = x
		x = kakutaniVonNeumann(x)
		kakutaniY[i] = x

		monnaX[i] = i
		monnaY[i] = monna(i + 1)
	}

	return kakutaniX, kakutaniY, monnaX, monnaY
}

func process(n int) ([]float64, []float64, []int, []float64) {
	return vanDerCorput(n)
}

func Draw() *components.Page {
	page := components.NewPage()

	for i := 30; i <= 100; i += 10 {
		kakutaniX, kakutaniY, monnaX, monnaY := process(i)

		kakutaniScatterData := make([]opts.ScatterData, len(kakutaniY))
		for i, v := range kakutaniY {
			kakutaniScatterData[i] = opts.ScatterData{
				Value:      []float64{kakutaniX[i], v},
				SymbolSize: 5,
			}
		}
		kakutaniScatter := charts.NewScatter()
		kakutaniScatter.SetGlobalOptions(
			charts.WithTitleOpts(opts.Title{Title: fmt.
				Sprintf("D, Kakutani Von Neumann, n=%d", i)}),
			charts.WithXAxisOpts(opts.XAxis{
				Name:        "",
				SplitLine:   &opts.SplitLine{Show: true},
				MaxInterval: 0.1,
				Min:         0.0,
				Max:         1.0,
			}),
			charts.WithYAxisOpts(opts.YAxis{
				Name:      "",
				SplitLine: &opts.SplitLine{Show: true},
			}),
		)
		kakutaniScatter.AddSeries("kakutani", kakutaniScatterData)

		monnaLineData := make([]opts.LineData, len(monnaY))
		for i, v := range monnaY {
			monnaLineData[i] = opts.LineData{
				Value:      []float64{float64(monnaX[i]), v},
				Symbol:     "roundRect",
				SymbolSize: 5,
			}
		}
		monnaLine := charts.NewLine()
		monnaLine.SetGlobalOptions(
			charts.WithTitleOpts(opts.Title{Title: fmt.
				Sprintf("D, Monna, n=%d", i)}),
			charts.WithXAxisOpts(opts.XAxis{
				Name:        "",
				SplitLine:   &opts.SplitLine{Show: false},
				MinInterval: 1.0,
				MaxInterval: 5.0,
			}),
			charts.WithYAxisOpts(opts.YAxis{
				Name:      "",
				SplitLine: &opts.SplitLine{Show: true},
			}),
		)
		monnaLine.AddSeries("monna", monnaLineData)

		page.AddCharts(kakutaniScatter, monnaLine)
	}

	return page
}
