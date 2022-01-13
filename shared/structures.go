package shared

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"strconv"
)

type node struct {
	Value int
	Next  *node
}

func newNode(val int) *node {
	return &node{val, nil}
}

type LinkedList struct {
	Size int
	Head *node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) ToArray() []int {
	arr := make([]int, 0)
	if ll == nil {
		return arr
	}

	cur := ll.Head
	for ; cur.Next != nil; cur = cur.Next {
		arr = append(arr, cur.Value)
	}
	arr = append(arr, cur.Value)

	return arr
}

func (ll *LinkedList) ToGraphNodes() []opts.GraphNode {
	nodes := make([]opts.GraphNode, 0)
	if ll == nil {
		return nodes
	}
	arr := ll.ToArray()

	keys := make(map[int]bool)
	for _, e := range arr {
		if _, v := keys[e]; !v {
			keys[e] = true
			nodes = append(nodes, opts.GraphNode{Name: strconv.Itoa(e)})
		}
	}

	return nodes
}

func (ll *LinkedList) GenGraphLinks() []opts.GraphLink {
	links := make([]opts.GraphLink, 0)
	nodes := ll.ToArray()
	nodesLen := len(nodes)
	for i := 1; i < nodesLen; i++ {
		links = append(links, opts.GraphLink{Source: strconv.Itoa(nodes[i-1]),
			Target: strconv.Itoa(nodes[i])})
	}

	return links
}

func (ll *LinkedList) Append(val int) {
	n := newNode(val)
	ll.Size++
	if ll.Head == nil {
		ll.Head = n
		return
	}

	cur := ll.Head
	for ; cur.Next != nil; cur = cur.Next {
	}
	cur.Next = n
}

func (ll *LinkedList) inList(val int) bool {
	cur := ll.Head

	if cur == nil {
		return false
	}

	if cur.Value == val {
		return true
	}

	for ; cur.Next != nil; cur = cur.Next {
		if cur.Value == val {
			return true
		}
	}

	return false
}

func DrawGraph(nodes []opts.GraphNode, links []opts.GraphLink,
	title string) *charts.Graph {
	graph := charts.NewGraph()

	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: title}),
	)

	graph.AddSeries("graph", nodes, links,
		charts.WithGraphChartOpts(opts.
			GraphChart{
			Layout:         "",
			Force:          &opts.GraphForce{Repulsion: 1000},
			Roam:           true,
			EdgeSymbol:     []string{"circle", "arrow"},
			EdgeSymbolSize: []int{5, 10},
			Draggable:      false,
		}),
		charts.WithLabelOpts(opts.Label{Show: true}))

	return graph
}
