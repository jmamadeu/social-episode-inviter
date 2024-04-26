package util

import (
	"math"
	"math/rand/v2"
)

func GenerateDigits() int {
	value := int(math.Floor(rand.Float64() * ((9999 - 1000) + 1000)))

	return value
}
