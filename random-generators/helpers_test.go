package randomGenerators_test

import (
	"math/rand/v2"
)

type namedFloat64 float64

func randomSource() rand.Source {
	return rand.NewPCG(123, 456)
}
