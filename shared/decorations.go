package shared

import (
	"fmt"
	"strings"
)

// PrintLists
func _(lists []*LinkedList) {
	for i, l := range lists {
		fmt.Println("Цикл", i)
		PrintList(l)
	}
}

func PrintList(l *LinkedList) {
	start := l.Head

	fmt.Printf("%v", start.Value)
	start = start.Next
	for j := 1; j < l.Size; j++ {
		fmt.Printf(" -> %v", start.Value)
		start = start.Next
	}
	fmt.Println()
}

// PrintSeparator
func _() {
	fmt.Println(strings.Repeat("*", 60))
}
