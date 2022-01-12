package E

import (
	"fmt"
	"math"
	"net/http"
	"sort"
	"strings"
)

/*

Вычислить (бинормализированный) параметр стахостичности для орбиты {a^l mon n}_{l = 0}^{inf}
	динамической системы на группе Г(n) взаимно простых вычетов по модулю натурального n
	при эволюции f(x) = ax mod n, где (a, n) = 1 (т.е. a и n взаимно простые)
*/

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func containsInt(arr []int, e int) bool {
	for _, v := range arr {
		if v == e {
			return true
		}
	}

	return false
}

func indexOfInt(arr []int, e int) int {
	for k, v := range arr {
		if e == v {
			return k
		}
	}
	return -1
}

func f(x, a, n int) int {
	return a * x % n
}

func processOrbit(orbit []int, g []int, m int) (float64, float64, float64) {
	distances := make([]int, 0)
	for _, e := range orbit {
		distances = append(distances, indexOfInt(g, e))
	}

	R := math.Pow(float64(m-distances[len(distances)-1]+distances[0]), 2)
	for i := 0; i < len(distances)-1; i++ {
		R += math.Pow(float64(distances[i+1]-distances[i]), 2)
	}
	r := R / math.Pow(float64(m), 2)
	s := r * float64(len(orbit))
	return R, r, s
}

func getGamma(n int) []int {
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		if gcd(i, n) == 1 {
			result = append(result, i)
		}
	}

	return result
}

func getOrbits(g []int, a, n int) [][]int {
	orbs := make([][]int, 0)
	usedGElems := make([]int, 0)

	for _, e := range g {
		x := e
		if containsInt(usedGElems, x) {
			continue
		}
		usedGElems = append(usedGElems, x)

		newOrb := make([]int, 1)
		newOrb[0] = x

		x = f(x, a, n)
		usedGElems = append(usedGElems, x)

		for newOrb[0] != x {
			newOrb = append(newOrb, x)
			x = f(x, a, n)
			usedGElems = append(usedGElems, x)
		}

		sort.Ints(newOrb)

		orbs = append(orbs, newOrb)
	}

	return orbs
}

func run(a, n int, w http.ResponseWriter) {
	if gcd(a, n) == 1 {
		g := getGamma(n)
		m := len(g)

		fmt.Fprintf(w, "a = %d\tn = %d\n", a, n)
		fmt.Fprintf(w, "Г(n): %v\n", g)
		fmt.Fprintln(w, strings.Repeat("-", 60))

		orbs := getOrbits(g, a, n)

		for i, v := range orbs {
			fmt.Fprintf(w, "Orb%d: %v\n", i, v)
			R, r, s := processOrbit(v, g, m)
			fmt.Fprintf(w, "\tR = %f\tr = %f\ts = %f\n", R, r, s)
		}
	} else {
		fmt.Fprintln(w, "a and b is not co-prime")
	}

}

func Process(w http.ResponseWriter) {
	fmt.Fprintf(w, "Task E\n\n")

	for i := 2; i <= 30; i += 7 {
		for j := 30; j <= 58; j += 7 {
			if gcd(i, j) == 1 {
				run(i, j, w)
				fmt.Fprintln(w, strings.Repeat("*", 80))
				fmt.Fprintln(w, "")
			}
		}
	}
}
