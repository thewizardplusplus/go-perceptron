package randomGenerators

import (
	"fmt"
	"math/rand/v2"
	"reflect"

	"golang.org/x/exp/constraints"
)

var _ Interface[float64] = (*Float[float64])(nil)

type Float[ValueType constraints.Float] struct {
	Base[ValueType]

	valueTypeKind reflect.Kind
}

func NewFloat[ValueType constraints.Float](source rand.Source) Float[ValueType] {
	return Float[ValueType]{
		Base:          NewBase[ValueType](source),
		valueTypeKind: reflect.TypeFor[ValueType]().Kind(),
	}
}

func (generator Float[ValueType]) NextNormalValue(params NormalParams[ValueType]) ValueType {
	return ValueType(generator.NormFloat64(params))
}

func (generator Float[ValueType]) NextUniformValue(r Range[ValueType]) ValueType {
	var randomValue ValueType
	switch generator.valueTypeKind {
	case reflect.Float32:
		randomValue = ValueType(generator.innerGenerator.Float32() * float32(r.Len()))

	case reflect.Float64:
		randomValue = ValueType(generator.innerGenerator.Float64() * float64(r.Len()))

	default:
		panic(fmt.Sprintf("unexpected type: %T", randomValue))
	}

	return randomValue + r.Min
}
