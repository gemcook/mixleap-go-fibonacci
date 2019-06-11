package fibonacci

import (
	"testing"
)

func TestFib(t *testing.T) {
	tests := []struct {
		name string // テストの名前
		num  int    // 入力値
		want int    // 返却値
	}{
		// TODO:テストパターンを作成する
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
		// TODO:テストパターンを作成する
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibS(tt.num); got != tt.want {
				t.Errorf("FibS(%v) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}

func BenchmarkFib(b *testing.B) {
	// b.N はベンチマークをするためにいい感じの数値をパッケージ側で設定してくれる
	for i := 0; i < b.N; i++ {
		// TODO:Fib関数を呼び出す
	}
}

func BenchmarkFibS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// TODO:FibS関数を呼び出す
	}
}
