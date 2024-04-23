package util

import "math/rand"

func GenerateDigits() int {
	return 10000 + rand.Intn(10000-9999)
}
