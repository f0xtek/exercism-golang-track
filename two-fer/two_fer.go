// Package twofer returns the two-for-one string for a given name
package twofer

import "fmt"

// ShareWith returns the string "One for name, one for me." for the given argument "name",
// If no name is given, return the string "One for you, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
