package kotori

import (
	"testing"
)

func TestSum(t *testing.T) {
	v := Sum([]int{1, 2, 4})

	if v != 7 {
		t.Error("Expected 7, but got ", v)
	}
}

func TestMin(t *testing.T) {
	v := Min([]float64{2, 21, 1})

	if v != 1 {
		t.Error("Expected 1, but got ", v)
	}
}

func TestMax(t *testing.T) {
	v := Max([]float64{2, 21, 1})

	if v != 21 {
		t.Error("Expected 1, but got ", v)
	}
}
