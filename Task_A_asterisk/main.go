package A_asterisk

import (
	"fundamental/tasks/shared"
	"github.com/go-echarts/go-echarts/v2/components"
)

// 5 Вариант
/*

Транзитивна ли по модулю 256 заданная функция f = f_N : Z_2 -> Z_2

f(x) = (x XOR 1) XOR (2 (x AND (1 + 2x) AND (3 + 4x) AND (7 + 8x) AND (
15 + 16x) AND (31 + 32x) AND (63 + 64x))) XOR (4(x^2 + 5))
*/

func f(x int) int {
	return (x ^ 1) ^ (2 * (x & (1 + 2*x) & (3 + 4*x) & (7 + 8*x) & (15 + 16*x) & (31 + 32*x) & (63 + 64*x))) ^ (4 * (x*x + 5))
	// return (x ^ 1) ^ (2 * (x & (1 + 2*x) & (3 + 4*x) & (7 + 8*x) & (
	//	15 + 16*x) & (31 + 32*x) & (63 + 64*x))) ^ (4 * (x*x + 11))
}

func checkTransitivityMod256(f func(int) int) (bool, *shared.LinkedList) {
	mod := 256
	list := shared.NewLinkedList()
	numbers := shared.MakeIntMap(mod)
	curPoint := 0

	return shared.FindCycleTransitivity(f, &mod, list, numbers, &curPoint)
}

func Check() *components.Page {

	status, list := checkTransitivityMod256(f)

	nodes := list.ToGraphNodes()
	links := list.GenGraphLinks()

	graph := shared.DrawGraph(nodes, links, "A*, "+func() string {
		if status {
			return "транзитивна"
		} else {
			return "не транзитивна"
		}
	}())

	return components.NewPage().AddCharts(graph)
}
