package loop

import (
	"regexp"
)

// StringMatchWithRegexp は文字列を受け取り16進数表記であれば true を返します
func StringMatchWithRegexp(str string) bool {
	r := regexp.MustCompile(`^0[xX][0-9A-Fa-f]+$`)
	return r.MatchString(str)
}

// StringMatchWithoutRegexp は文字列を受け取り16進数表記であれば true を返します
// 無名関数で実装すると更にちょっとだけ速くなります。
func StringMatchWithoutRegexp(str string) bool {
	// "0x"から始まらないパターンは false を返す
	if len(str) < 3 || str[0] != '0' || str[1] != 'x' && str[1] != 'X' {
		return false
	}
	// loop で一文字ずつ比較し、一文字でも当てはまらない文字があれば false を返す
	for _, c := range str[2:] {
		// ASCII : 0x30 ~ 0x39 , 0x41 ~ 0x5A , 0x61 ~ 0x7a のみ許可
		if c < '0' || '9' < c && c < 'A' || 'F' < c && c < 'a' || 'f' < c {
			return false
		}
	}
	return true
}
