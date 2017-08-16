package popcount

import "testing"

func TestPopCount(t *testing.T) {
	var tests = []struct {
		input uint64
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{4, 1},
		{8, 1},
		{16, 1},
		{32, 1},
		{64, 1},
		{128, 1},
		{256, 1},
		{512, 1},
		{3, 2},
		{13, 3},
	}

	for _, test := range tests {
		if got := PopCount(test.input); got != test.want {
			t.Errorf("PopCount(%d) = %d, want %d\n", 
				test.input, got, test.want)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i:=0; i<b.N; i++ {
		PopCount(13)
	}
}
