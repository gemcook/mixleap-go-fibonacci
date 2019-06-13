package fibonacci

// Fib はフィボナッチ数を返します
func Fib(num int) int {
	// 0番目は 0, 1番目は 1 なので 2未満はそのまま返却
	if num < 2 {
		return num
	}
	// [num]番目は[num-2]番目＋[num-1]番目の計算結果を返却する
	return Fib(num-2) + Fib(num-1)
}

// FibS はフィボナッチ数をマップで作成し、指定された値を返します
func FibS(num int) int {
	if num < 2 {
		return num
	}
	// num番目＋1 のスライスを作成
	memo := make([]int, num+1)
	memo[0] = 0 // 0番目には0を挿入
	memo[1] = 1 // 1番目には1を挿入

	// [num]番目は 2 ~ [num]番目まで、[index-2]番目＋[index-1]番目 の計算を
	// 繰り返しながら格納していくことで計算できる
	for i := 2; i <= num; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}

	// [num]番目に入っているのはFibonacciの[num]番目
	return memo[num]
}
