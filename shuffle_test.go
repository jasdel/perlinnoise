package perlinnoise

import (
	"math/rand"
	"testing"
)

// Simple structure interface for
type testArray []int

func (a testArray) Len() int {
	return len(a)
}
func (a testArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type testCase struct {
	in   testArray
	out  testArray
	seed int64
}

var testData = []testCase{
	testCase{in: testArray{0, 1, 2, 3, 4, 5, 6, 7, 8}, out: testArray{3, 0, 4, 2, 6, 8, 1, 7, 5}, seed: 1},
	testCase{in: testArray{0, 1, 2, 3, 4, 5, 6, 7, 8}, out: testArray{4, 8, 1, 7, 6, 5, 3, 0, 2}, seed: 10},
	testCase{in: testArray{0, 1, 2, 3, 4, 5, 6, 7}, out: testArray{6, 2, 4, 1, 5, 7, 3, 0}, seed: 3},
	testCase{in: testArray{0, 1}, out: testArray{1, 0}, seed: 0},
	testCase{in: testArray{0, 1}, out: testArray{0, 1}, seed: 1},
}

// Verifies shuffling works correctly.
func TestShuffle(t *testing.T) {
	for testNum, test := range testData {
		in := make(testArray, len(test.in))
		copy(in, test.in)

		Shuffle(in, rand.New(rand.NewSource(test.seed)))

		if in.Len() != test.out.Len() {
			t.Fatal("Case", testNum, "In and Out collection lengths do not match", in.Len(), test.out.Len())
		}

		for i, v := range test.out {
			if in[i] != v {
				t.Error("Case", testNum, "Expected", test.out, "got", in)
				break
			}
		}
	}
}
