// これは「Goのループはそこそこ速い」ということを比べやすくするために用意したものです。
// こちらの記事がとても良かったので共有したいと思い一部をお借りして作成しました。
// https://qiita.com/naoina/items/d71ddfab31f4b29f6693
// ピックアップしたのは文字列が16進数表記になっていれば true を返すという関数です。

// 標準パッケージの regexp を使用したものと、ループで一文字づつ判断していくものを用意しました。
// 実行してもらえればわかりますが、Goはループでゴリゴリ実装してもそこまでパフォーマンスが落ちません。
// (regexp は内部でアロケーションしているので速度が出ないというのもあります。)
// 再帰フィボナッチの例では、同じ計算を 2回以上呼ぶ無駄が発生していたので、とんでもなく遅い結果となりました。
// 比べて、ループで計算していたものは、引数を大きくしても全然スピードが落ちていませんでしたよね？

// 当たり前ですが、パフォーマンスを上げたい場合は無駄な処理をしないことや
// なるべくアロケーションをしない実装を検討する必要があります。
// そうするとこの例のようにループで地道に回したほうが速いということになります。

// C言語では malloc や realloc などで苦しみながらメモリ管理をしていましたが
// それを気にしないで良い分「なんか遅い」みたいな状況に陥ることも出てきそうです。

// performance_tuningディレクトリ直下で go test -bench . -benchmem で実行してみてください。
// allocs/op と出力された値が 1回の実行でメモリアロケーションした回数です。

package loop

import (
	"testing"
)

func TestStringMatchWithRegexp(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"[success] 16進数に現れる物のみ", args{str: "0x1234567890ABCDEF"}, true},
		{"[fail] 最後だけ16進数表記ではない", args{str: "0x1234567890ABCDEFG"}, false},
		{"[fail] 最初だけ16進数表記ではない", args{str: "1x1234567890ABCDEF"}, false},
		{"[fail] 真ん中の一個だけ16進数表記ではない", args{str: "0x12345678H0ABCDEF"}, false},
		{"[fail] 全て16進数表記ではない", args{str: "FriskStrongMint"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringMatchWithRegexp(tt.args.str); got != tt.want {
				t.Errorf("StringMatchWithRegexp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringMatchWithoutRegexp(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"[success] 16進数に現れる物のみ", args{str: "0x1234567890ABCDEF"}, true},
		{"[fail] 最後だけ16進数表記ではない", args{str: "0x1234567890ABCDEFG"}, false},
		{"[fail] 最初だけ16進数表記ではない", args{str: "1x1234567890ABCDEF"}, false},
		{"[fail] 真ん中の一個だけ16進数表記ではない", args{str: "0x12345678H0ABCDEF"}, false},
		{"[fail] 全て16進数表記ではない", args{str: "FriskStrongMint"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringMatchWithoutRegexp(tt.args.str); got != tt.want {
				t.Errorf("StringMatchWithoutRegexp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkStringMatchWithRegexp(b *testing.B) {
	s := "0x1234567890ABCDEF"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringMatchWithRegexp(s)
	}
}

func BenchmarkStringMatchWithoutRegexp(b *testing.B) {
	s := "0x1234567890ABCDEF"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringMatchWithoutRegexp(s)
	}
}
