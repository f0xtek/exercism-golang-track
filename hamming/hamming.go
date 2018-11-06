package hamming

import "errors"

// Distance returns the Hamming distance for 2 given strings (DNA strands)
// or and error if the length of a and b are not equal
func Distance(a, b string) (int, error) {
	hammingDistance := 0
	if len(a) != len(b) {
		return 0, errors.New("len(a) != len(b)")
	}
	for i := range a {
		if a[i] != b[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}
