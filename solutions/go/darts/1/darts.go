package darts

import "math"

func Score(x, y float64) int {
	distance := math.Sqrt(math.Pow((x - 0), 2) + math.Pow((y - 0), 2))
    switch {
        case distance >= 0 && distance <= 1:
        	return 10
        case distance > 1 && distance <= 5:
        	return 5
		case distance > 5 && distance <= 10:
        	return 1
        default:
        	return 0
    }
}
