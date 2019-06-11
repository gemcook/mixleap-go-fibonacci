package fibonacci

import (
	"testing"
)

func TestFib(t *testing.T) {
	// テストパターンを構造体で作成
	tests := []struct {
		name string // テストの名前
		num  int    // 入力値
		want int    // 返却値
	}{
		// TODO:テストパターンを作成する
		// 構造体の宣言の後に続けて波括弧を始めることで初期化することが出来る
		// 構造体のスライスなので {"name", num, want }, のように波括弧で囲う形で書く必要がある
		// 全てのフィールドに初期値を設定する場合フィールド名を省略可能
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if の判断文の前に処理をさせることが出来る
			// ここではテスト対象の関数の実行を行っている
			if got := Fib(tt.num); got != tt.want {
				t.Errorf("Fib(%d) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}

func TestFibS(t *testing.T) {
	// テストパターンを構造体で作成
	tests := []struct {
		name string // テストの名前
		num  int    // 入力値
		want int    // 返却値
	}{
		// TODO:テストパターンを作成する
		// 構造体の宣言の後に続けて波括弧を始めることで初期化することが出来る
		// 構造体のスライスなので {"name", num, want }, のように波括弧で囲う形で書く必要がある
		// 全てのフィールドに初期値を設定する場合フィールド名を省略可能
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if の判断文の前に処理をさせることが出来る
			// ここではテスト対象の関数の実行を行っている
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
