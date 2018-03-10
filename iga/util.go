package iga

import (
	"math/rand"
)

func randInt(fromNum int, toNum int) int {
	return rand.Intn(toNum-fromNum) + fromNum
}
