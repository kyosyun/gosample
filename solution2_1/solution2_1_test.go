package solution21

import (
	"reflect"
	"testing"
)

func TestSoluction(t *testing.T) {

	tests := []struct {
		inArray  []int
		rotate   int
		expected []int
	}{
		{
			[]int{1, 2, 3},
			1,
			[]int{3, 1, 2},
		},
	}

	for _, test := range tests {
		res := Solution(test.inArray, test.rotate)

		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("expected: %d, result: %d", test.expected, res)
		}

	}

}
