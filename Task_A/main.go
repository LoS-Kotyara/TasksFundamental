package A

import (
	"fundamental/tasks/shared"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"strconv"
)

// 5 вариант
/*

Проверить на биективность/транзитивность полиномы f(x) и g(x)

f(x) = 12 + 3x - 14x^2
g(x) = 11 + 9x
*/

func f(x int) int {
	// return 12 + 3*x - 14*x*x
	// return 18 + x - 7*x*x
	return 11 - 3*x + 2*x*x
}

func g(x int) int {
	return 11 + 9*x
}

func checkBijectivity(f func(int) int) (bool, []*shared.LinkedList) {
	mod := 4

	var lists []*shared.LinkedList
	lists = append(lists, shared.NewLinkedList())
	listsLength := 1

	numbers := shared.MakeIntMap(mod)

	curPoint := 0
	lists[0].Append(0)
	count := 1

	status := true
	var nextPoint int
	for count != mod {
		nextPoint = shared.Mod(f(curPoint), mod)
		lists[listsLength-1].Append(nextPoint)
		numbers[curPoint] = 1
		count++
		if numbers[nextPoint] == 1 {
			if nextPoint == lists[listsLength-1].Head.Value {
				lists = append(lists, shared.NewLinkedList())
				listsLength++

				temp := -1
				for i := range numbers {
					if numbers[i] == 0 {
						temp = i
						break
					}
				}

				if temp == -1 {
					panic("temp = -1")
				}

				curPoint = temp

				lists[listsLength-1].Append(curPoint)

			} else {
				status = false
				break
			}
		} else {
			curPoint = nextPoint
		}
	}

	nextPoint = shared.Mod(f(curPoint), mod)

	if status && nextPoint == lists[listsLength-1].Head.Value {
		lists[listsLength-1].Append(nextPoint)
	}

	return status, lists
}

func checkTransitivity(f func(int) int) (bool, *shared.LinkedList) {
	mod := 8
	list := shared.NewLinkedList()
	numbers := shared.MakeIntMap(mod)
	curPoint := 0

	return shared.FindCycleTransitivity(f, &mod, list, numbers, &curPoint)
}

type funcCheck struct {
	bijectivity struct {
		status bool
		lists  []*shared.LinkedList
	}
	transitivity struct {
		status bool
		list   *shared.LinkedList
	}
}

type A struct {
	f funcCheck
	g funcCheck
}

func checkFunction(function func(int) int) funcCheck {
	var res funcCheck

	res.bijectivity.status, res.bijectivity.lists = checkBijectivity(function)

	if !res.bijectivity.status {
		return res
	}

	res.transitivity.status, res.transitivity.list = checkTransitivity(function)

	if !res.transitivity.status {
		return res
	}

	return res
}

func bijectivityToGraph(lists []*shared.LinkedList,
	title string) *charts.Graph {
	nodes := make([]opts.GraphNode, 0)
	links := make([]opts.GraphLink, 0)
	for _, e := range lists {
		nodes = append(
			nodes,
			e.ToGraphNodes()...,
		)
		links = append(links, e.GenGraphLinks()...)
	}

	return shared.DrawGraph(nodes, links, title)
}

func transitivityToGraph(list *shared.LinkedList, title string) *charts.Graph {
	nodes := list.ToGraphNodes()
	links := list.GenGraphLinks()

	return shared.DrawGraph(nodes, links, title)
}

func Check() *components.Page {
	var res A

	res.f = checkFunction(f)
	res.g = checkFunction(g)

	page := components.NewPage()

	page.Width = strconv.Itoa(1280)
	page.AddCharts(
		bijectivityToGraph(res.f.bijectivity.lists, "A, f, "+func() string {
			if res.f.bijectivity.status {
				return "биективна"
			} else {
				return "не биективна"
			}
		}()))

	page.AddCharts(
		transitivityToGraph(res.f.transitivity.list, "A, f, "+func() string {
			if res.f.transitivity.status {
				return "транзитивна"
			} else {
				return "не транзитивна"
			}
		}()))

	page.AddCharts(
		bijectivityToGraph(res.g.bijectivity.lists, "A, g, "+func() string {
			if res.g.bijectivity.status {
				return "биективна"
			} else {
				return "не биективна"
			}
		}()))

	page.AddCharts(
		transitivityToGraph(res.g.transitivity.list, "A, g, "+func() string {
			if res.g.transitivity.status {
				return "транзитивна"
			} else {
				return "не транзитивна"
			}
		}()))

	return page

}
