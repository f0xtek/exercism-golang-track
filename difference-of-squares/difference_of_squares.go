package diffsquares

// SquareOfSum returns the square of the sum of all integers from 1 to n
func SquareOfSum(n int) (square int) {
	square = (n * (n + 1)) / 2
	square *= square
	return square
}

// SumOfSquares returns the sum of each integer from 1 to n squared
func SumOfSquares(n int) (sum int) {
	sum = (n * (n + 1) * (2*n + 1) / 6)
	return sum
}

// Difference returns the difference between the returned values from SquareOfSum
// and SumOfSquares
func Difference(n int) (diff int) {
	diff = SquareOfSum(n) - SumOfSquares(n)
	return diff
}
