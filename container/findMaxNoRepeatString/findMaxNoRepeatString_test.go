package main

import (
	"testing"
)

func TestFindMaxNoRepeatString(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		//normal case
		{"abc", 3},
		{"abccba", 3},

		//edge case
		{"", 0},
		{"aaaaa", 1},
		{"asdfghjk", 8},

		//chinese case
		{"我是小智", 4},
		{"小智智小", 2},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 5},
	}

	for _, tt := range tests {
		ans := findMaxNoRepeatString(tt.s)
		if ans != tt.ans {
			t.Errorf("got %d for input %s,"+
				"expected %d", ans, tt.s, tt.ans)
		}
	}
}

/**
性能测试
*/
func BenchmarkFindMaxNoRepeatString(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	ans := 5
	for i := 0; i < b.N; i++ {
		actual := findMaxNoRepeatString(s)
		if actual != ans {
			b.Errorf("got %d for input %s,"+
				"expected %d", actual, s, ans)
		}
	}
}
