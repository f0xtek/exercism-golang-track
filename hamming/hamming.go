package hamming

import (
	"fmt"
)

// Distance returns the Hamming distance for 2 given strings (DNA strands)
// or and error if the length of a and b are not equal
func Distance(a, b string) (int, error) {
	hammingDistance := 0
	if len(a) != len(b) {
		return 0, fmt.Errorf("len(a) != len(b)")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}
