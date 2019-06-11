package fibonacci

// Fib はフィボナッチ数を再帰で計算しその値を返します
func Fib(num int) int {
	// 0番目は0
	// 1番目は1
	// [num]番目は[num-2]番目＋[num-1]番目

	// TODO:計算結果を返却する
	return num
}

// FibS はフィボナッチ数をスライスで作成し、指定された値を返します
func FibS(num int) int {
	// num番目＋1 のスライスを作成
	memo := make([]int, num+1)

	// 0番目には0を挿入
	// 1番目には1を挿入

	// [num]番目は 2 ~ [num]番目まで、[index-2]番目＋[index-1]番目 の計算を
	// 繰り返しながら格納していくことで計算できる

	// [num]番目に入っているのはFibonacciの[num]番目
	return memo[num]
}
