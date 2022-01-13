package B

import (
	"reflect"
	"testing"
)

func TestProcess(t *testing.T) {
	input := []int{
		15, 16, 17, 18, 19, 20,
	}

	expected := make([][]float64, 0)

	result := make([][]float64, 0)
	for _, v := range input {
		status, _ := process(v)
		result = append(result, status)
		expected = append(expected, status)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fail()
		// t.Logf("Expected %v, got %v", expected, result)
	}

}
