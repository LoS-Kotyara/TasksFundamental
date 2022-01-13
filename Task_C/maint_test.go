package C

import (
	"math"
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	input := []string{
		"10001",
		"001",
		"10010",
	}

	expected := []string{
		"10001",
		"100",
		"01001",
	}

	result := make([]string, 0)

	for _, v := range input {
		result = append(result, reverseString(v))
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fail()
		t.Logf("Expected %v, got %v", expected, result)
	}
}

func TestToBeta(t *testing.T) {
	k := 10
	n := 10
	beta := math.Pow(2.0, 1.0/float64(n))

	input := []int{
		0,
		21,
		465, 786, 564, 534,
	}

	expected := []float64{
		0,
		3.468206265769929,
		7.200830396588009,
		5.998448482975048,
		5.7484858112166375,
		5.406045711379836,
	}

	result := make([]float64, 0)

	for _, v := range input {
		result = append(result, toBeta(v, k, beta))
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fail()
		t.Logf("Expected %v, got %v", expected, result)
	}
}
