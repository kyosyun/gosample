package solution1

import (
	"math"
	"strconv"
)

func Solution(N int) int {
	// binaryGapの最大の長さを返却する
	max := 0
	bin := strconv.FormatInt(int64(N), 2)
	tempMax := 0

	for i := 0; i < len(bin); i++ {
		// 1の場合
		if bin[i:i+1] == "1" {
			max = int(math.Max(float64(max), float64(tempMax)))
			tempMax = 0
		} else {
			// 0の場合
			tempMax += 1
		}
	}
	return max
}
