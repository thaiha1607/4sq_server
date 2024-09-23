package misc

import (
	"math/rand"
)

func RandomIntInRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}
