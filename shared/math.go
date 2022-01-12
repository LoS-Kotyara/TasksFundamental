package shared

func Mod(a, n int) int {
	return ((a % n) + n) % n
}

func MakeIntMap(a int) map[int]int {
	numbers := make(map[int]int, a)
	for i := 0; i < a; i++ {
		numbers[i] = 0
	}

	return numbers
}

func FindCycleTransitivity(
	f func(int) int,
	mod *int,
	list *LinkedList,
	numbers map[int]int,
	curPoint *int) (bool, *LinkedList) {

	if *curPoint == 0 {
		list.Append(*curPoint)
		numbers[*curPoint] = 1
	}

	count := &list.Size
	nextPoint := Mod(f(*curPoint), *mod)
	if *count == *mod {
		list.Append(nextPoint)
		if nextPoint == list.Head.Value {
			return true, list
		} else {
			return false, list
		}
	} else {
		list.Append(nextPoint)
		if numbers[nextPoint] == 1 {
			return false, list
		} else {
			numbers[nextPoint] = 1
			*curPoint = nextPoint
			return FindCycleTransitivity(f, mod, list, numbers, curPoint)
		}
	}
}
