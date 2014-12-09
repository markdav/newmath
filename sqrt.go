package newmath

import (
	"math"
)

const default_accuracy = 0.0001

func SqrtTo(x, accuracy float64) float64 {
	var z, z2 float64

	for z = x / 2; ; {
		z2 = z - (z*z-x)/(2.0*z)
		if math.Abs(z2-z) < accuracy {
			return z2
		}
		z = z2
	}

	return z
}

/**
* set the accura
**/
func Sqrt(x float64) float64 {
	return SqrtTo(x, default_accuracy)
}
