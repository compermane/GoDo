package functions

import (
	"math/rand"
)

func ChooseRandom(fns []*Function) *Function {
	return fns[rand.Intn(len(fns))]
}