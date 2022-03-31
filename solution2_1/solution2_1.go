package solution21

func Solution(A []int, K int) []int {

	// Aに入ってる配列をKの個数だけ動かす。

	// 同じ長さの配列の作成
	res := make([]int, len(A))

	// 配列の長さが0件だった場合は即座に終了
	if len(A) == 0 {
		return A
	}

	// 回転数の確認
	rotate := K % len(A)

	for i := 0; i < len(A); i++ {
		newPosi := i + rotate
		if newPosi > len(A)-1 {
			newPosi -= len(A)
		}
		res[newPosi] = A[i]
	}
	return res
}
