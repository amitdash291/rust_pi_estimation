package go_utils

import "math"

func IsPointWithinCircle(xPos float64, yPos float64) bool {
	return math.Pow(xPos, 2)+math.Pow(yPos, 2) < 1.0
}
