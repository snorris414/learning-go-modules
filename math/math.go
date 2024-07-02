// math provides functions for doing basic math operations.
package math

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add takes 2 Numbers and returns their sum
// For more information about adding numbers, see: https://www.mathisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
