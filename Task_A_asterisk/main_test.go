package A_asterisk

import (
	"reflect"
	"testing"
)

func TestF1(t *testing.T) {
	f := func(x int) int {
		return (x ^ 1) ^ (2 * (x & (1 + 2*x) & (3 + 4*x) & (7 + 8*x) &
			(15 + 16*x) & (31 + 32*x) & (63 + 64*x))) ^ (4 * (x*x + 5))
	}

	input := []int{
		0, 4, 8, 12, 16, 20, 24, 28, 32,
	}

	expected := []int{
		21, 81, 285, 601, 1029, 1601, 2317, 3145, 4149,
	}

	result := make([]int, 0)

	for _, v := range input {
		result = append(result, f(v))
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fail()
		t.Logf("Expected %v, got %v", expected, result)
	}
}

func TestF2(t *testing.T) {
	f := func(x int) int {
		return (x ^ 1) ^ (2 * (x & (1 + 2*x) & (3 + 4*x) & (7 + 8*x) &
			(15 + 16*x) & (31 + 32*x) & (63 + 64*x))) ^ (4 * (x*x + 11))
	}

	input := []int{
		0, 4, 8, 12, 16, 20, 24, 28, 32,
	}

	expected := []int{
		45, 105, 293, 609, 1085, 1657, 2357, 3185, 4109,
	}

	result := make([]int, 0)

	for _, v := range input {
		result = append(result, f(v))
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fail()
		t.Logf("Expected %v, got %v", expected, result)
	}
}

func TestCheckTransitivityMod256F1(t *testing.T) {
	input := []func(int) int{
		func(x int) int {
			return (x ^ 1) ^ (2 * (x & (1 + 2*x) & (3 + 4*x) & (7 + 8*x) &
				(15 + 16*x) & (31 + 32*x) & (63 + 64*x))) ^ (4 * (x*x + 5))
		},
		func(x int) int {
			return (x ^ 1) ^ (2 * (x & (1 + 2*x) & (3 + 4*x) & (7 + 8*x) &
				(15 + 16*x) & (31 + 32*x) & (63 + 64*x))) ^ (4 * (x*x + 11))
		},
	}

	expected := []bool{true, true}

	result := make([]bool, 0)
	for _, v := range input {
		status, _ := checkTransitivityMod256(v)
		result = append(result, status)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fail()
		t.Logf("Expected %v, got %v", expected, result)
	}

}
