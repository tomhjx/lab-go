package random

import (
	"math/rand"
	"time"
)

func GetIntn(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}
