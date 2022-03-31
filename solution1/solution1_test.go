package solution1

import (
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {

	tests := []struct {
		in     int
		expect int
	}{
		{
			9,
			2,
		},
		{
			529,
			4,
		}, {
			20,
			1,
		}, {
			1041,
			5,
		},
	}

	for _, test := range tests {
		res := Solution(test.in)
		if test.expect != res {
			t.Errorf("result: %s, expected: %s", strconv.Itoa(res), strconv.Itoa(test.expect))
		}
	}
}
