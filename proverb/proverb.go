package proverb

import "fmt"

// Proverb returns the relevant proverb for the given slice of strings
func Proverb(rhyme []string) []string {
	var s []string
	if len(rhyme) == 0 {
		return s
	} else if len(rhyme) > 1 {
		for i := 0; i < len(rhyme)-1; i++ {
			s = append(s, fmt.Sprintf("For want of a %v the %v was lost.", rhyme[i], rhyme[i+1]))
		}
	}
	s = append(s, fmt.Sprintf("And all for the want of a %v.", rhyme[0]))
	return s
}
