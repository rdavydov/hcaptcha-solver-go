package curves

import "math"

// Point represents a point in 2D space.
type Point struct {
	// x, y represent the coordinates of the point.
	x, y float64
}

// binomial returns the binomial coefficient "n choose k".
func binomial(n, k int) float64 {
	return float64(factorial(n)) / float64(factorial(k)*factorial(n-k))
}

// bernsteinPolynomialPoint calculates the i-th component of a bernstein polynomial of degree n.
func bernsteinPolynomialPoint(x float64, i, n int) float64 {
	return binomial(n, i) * math.Pow(x, float64(i)) * (math.Pow(1-x, float64(n-i)))
}

// factorial returns the factorial of n.
func factorial(n int) int {
	if n >= 1 {
		return n * factorial(n-1)
	}
	return 1
}
