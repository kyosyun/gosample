package solution21

import "fmt"

func Solution(A []int, K int) []int {

	// Aに入ってる配列をKの個数だけ動かす。

	// 同じ長さの配列の作成
	res := make([]int, len(A))

	// 回転数の確認
	rotate := K % len(A)

	for i := 0; i < len(A); i++ {
		newPosi := i + rotate
		if newPosi > len(A)-1 {
			newPosi -= len(A)
		}
		fmt.Printf("i: %d, rotate: %d, newPosi: %d\n", i, rotate, newPosi)
		fmt.Printf("★res[%d] <= %d \n", newPosi, A[i])
		res[newPosi] = A[i]
	}
	return res
}
