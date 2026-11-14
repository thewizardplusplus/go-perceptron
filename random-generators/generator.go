package randomGenerators

import (
	"fmt"
	"math/rand/v2"
	"reflect"

	perceptron "github.com/thewizardplusplus/go-perceptron"
)

var _ Interface[float64] = (*Generator[float64])(nil)

type Generator[ValueType perceptron.ValueConstraint] struct {
	innerGenerator rand.Rand
	valueTypeKind  reflect.Kind
}

func NewGenerator[ValueType perceptron.ValueConstraint](source rand.Source) Generator[ValueType] {
	return Generator[ValueType]{
		innerGenerator: *rand.New(source),
		valueTypeKind:  reflect.TypeFor[ValueType]().Kind(),
	}
}

func (generator Generator[ValueType]) NextNormalValue(params NormalParams[ValueType]) ValueType {
	mean, standardDeviation := float64(params.Mean), float64(params.StandardDeviation)
	return ValueType(generator.innerGenerator.NormFloat64()*standardDeviation + mean)
}

func (generator Generator[ValueType]) NextUniformValue(r Range[ValueType]) ValueType {
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
