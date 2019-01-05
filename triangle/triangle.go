/*
Package triangle implements a simple library to determine
the kind of triangle given the length of its sides
*/
package triangle

import (
	"math"
)

type Kind string

// Kinds of triangle
const (
	NaT = "NaT"
	Equ = "Equ"
	Iso = "Iso"
	Sca = "Sca"
)

// KindFromSides returns the Kind of triangle given the length of its sides
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	switch {
	case math.IsNaN(a + b + c):
		k = NaT
	case math.IsInf(a+b+c, 1) || math.IsInf(a+b+c, -1):
		k = NaT
	case a <= 0 || b <= 0 || c <= 0:
		k = NaT
	case a+b < c || b+c < a || a+c < b:
		k = NaT
	case a == b && b == c:
		k = Equ
	case a == b || b == c || a == c:
		k = Iso
	default:
		k = Sca
	}

	return k
}
