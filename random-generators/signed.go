package randomGenerators

import (
	"math"
	"math/rand/v2"

	"golang.org/x/exp/constraints"
)

var _ Interface[int64] = (*Signed[int64])(nil)

type Signed[ValueType constraints.Signed] struct {
	Base[ValueType]
}

func NewSigned[ValueType constraints.Signed](source rand.Source) Signed[ValueType] {
	return Signed[ValueType]{
		Base: NewBase[ValueType](source),
	}
}

func (generator Signed[ValueType]) NextNormalValue(params NormalParams[ValueType]) ValueType {
	return ValueType(math.Round(generator.NormFloat64(params)))
}

func (generator Signed[ValueType]) NextUniformValue(r Range[ValueType]) ValueType {
	return ValueType(generator.innerGenerator.Int64N(r.LenAsInt64())) + r.Min
}
