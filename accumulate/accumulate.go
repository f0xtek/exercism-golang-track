package accumulate

// Accumulate accepts a collection and an operation
// and returns a new collection containing the result of
// applying that operation to each element of the input collection.
func Accumulate(s []string, f func(string) string) []string {
	result := make([]string, len(s))
	for _, v := range s {
		result = append(result, f(v))
	}
	return result
}
