package main

import "testing"

func TestRriangle(t *testing.T) {

	tests := []struct {
		a, b, c int
	}{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}

	for _, tt := range tests {

		if c := triangle(tt.a, tt.b); c != tt.c {
			t.Errorf("triangle (%d,%d); "+
				"got %d; expected %d", tt.a, tt.b, c, tt.c)
		}

	}

}
