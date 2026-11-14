package randomGenerators

import (
	"math/rand/v2"

	perceptron "github.com/thewizardplusplus/go-perceptron"
)

type NormalParams[ValueType perceptron.ValueConstraint] struct {
	Mean              ValueType
	StandardDeviation ValueType
}

type Base[ValueType perceptron.ValueConstraint] struct {
	innerGenerator rand.Rand
}

func NewBase[ValueType perceptron.ValueConstraint](source rand.Source) Base[ValueType] {
	return Base[ValueType]{
		innerGenerator: *rand.New(source),
	}
}

func (generator Base[ValueType]) NormFloat64(params NormalParams[ValueType]) float64 {
	mean, standardDeviation := float64(params.Mean), float64(params.StandardDeviation)
	return generator.innerGenerator.NormFloat64()*standardDeviation + mean
}
