package raindrops

import (
	"fmt"
	"strings"
)

// Convert returns Pling, Plang, and Plong if the input has 3, 5, or 7 as a factor
// return the input if not.
func Convert(input int) string {
	var str []string
	factors := [3]int{3, 5, 7}

	if input%factors[0] == 0 {
		str = append(str, "Pling")
	}
	if input%factors[1] == 0 {
		str = append(str, "Plang")
	}
	if input%factors[2] == 0 {
		str = append(str, "Plong")
	}
	if strings.Join(str, "") == "" {
		return fmt.Sprintf("%d", input)
	}
	return strings.Join(str, "")
}
