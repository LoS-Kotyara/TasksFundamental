package A

import (
	"reflect"
	"testing"
)

func TestF(t *testing.T) {
	input := []int{
		0, 1, 2, 3, 4, 5, 6, 7,
	}

	expected := []int{
		11, 10, 13, 20, 31, 46, 65, 88,
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

func TestCheckBijectivityStatuses(t *testing.T) {
	input := []func(int) int{
		func(x int) int {
			return 11 - 3*x + 2*x*x
		},
		func(x int) int {
			return 12 + 3*x - 14*x*x
		},
		func(x int) int {
			return 11 + 9*x
		},
		func(x int) int {
			return 18 + x - 7*x*x
		},
	}

	expected := []bool{true, true, true, false}
	result := make([]bool, 0)

	for _, v := range input {
		status, _ := checkBijectivity(v)
		result = append(result, status)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fail()
		t.Logf("Expected %v, got %v", expected, result)
	}

}

func TestCheckTransitivityStatuses(t *testing.T) {
	input := []func(int) int{
		func(x int) int {
			return 11 - 3*x + 2*x*x
		},
		func(x int) int {
			return 12 + 3*x - 14*x*x
		},
		func(x int) int {
			return 11 + 9*x
		},
		func(x int) int {
			return 18 + x - 7*x*x
		},
	}

	expected := []bool{false, false, true, false}
	result := make([]bool, 0)

	for _, v := range input {
		status, _ := checkTransitivity(v)
		result = append(result, status)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fail()
		t.Logf("Expected %v, got %v", expected, result)
	}

}
