package luhn

import "strings"

// Valid determines whether a number is valid as per the Luhn formula. The number should be passed as a string type.
// It returns true if the number is valid, and false if not.
func Valid(n string) bool {
	n = strings.ReplaceAll(n, " ", "")
	if len(n) < 2 {
		return false
	}

	sum := 0
	double := len(n)%2 == 0
	for _, r := range n {
		digit := int(r - '0')
		if digit < 0 || digit > 9 {
			return false
		}
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}
	return sum%10 == 0
}
