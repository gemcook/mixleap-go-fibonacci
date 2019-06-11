package fibonacci

import "testing"

func TestFib(t *testing.T) {
	tests := []struct {
		name string // テストの名前
		num  int    // 入力値
		want int    // 返却値
	}{
		{"1の時1を返す", 1, 1},
		{"2の時1を返す", 2, 1},
		{"3の時2を返す", 3, 2},
		{"7の時13を返す", 7, 13},
		{"20の時6765を返す", 20, 6765},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib(tt.num); got != tt.want {
				t.Errorf("Fib(%d) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}

func TestFibS(t *testing.T) {
	tests := []struct {
		name string // テストの名前
		num  int    // 入力値
		want int    // 返却値
	}{
		{"1の時1を返す", 1, 1},
		{"2の時1を返す", 2, 1},
		{"3の時2を返す", 3, 2},
		{"7の時13を返す", 7, 13},
		{"20の時6765を返す", 20, 6765},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibS(tt.num); got != tt.want {
				t.Errorf("Fib(%d) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}

func BenchmarkFib(b *testing.B) {
	// b.N はベンチマークをするためにいい感じの数値をパッケージ側で設定してくれる
	for i := 0; i < b.N; i++ {
		// 実行結果を比較しやすいように大きい数値を入れる
		Fib(35)
	}
}

func BenchmarkFibS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibS(35)
	}
}
