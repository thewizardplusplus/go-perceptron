package weightInitializers

import (
	"fmt"
	"math"

	perceptron "github.com/thewizardplusplus/go-perceptron"
	randomGenerators "github.com/thewizardplusplus/go-perceptron/random-generators"
)

var _ perceptron.WeightInitializer[float64] = He[float64]{}

type He[ValueType perceptron.ValueConstraint] struct {
	base[ValueType]
}

func NewHe[ValueType perceptron.ValueConstraint](
	randomGenerator randomGenerators.Interface[ValueType],
	distribution randomGenerators.Distribution,
) He[ValueType] {
	return He[ValueType]{
		base: newBase(randomGenerator, distribution),
	}
}

func (initializer He[ValueType]) InitWeight(params perceptron.InitWeightParams) ValueType {
	switch initializer.distribution {
	case randomGenerators.NormalDistribution:
		return initializer.randomGenerator.NextNormalValue(
			randomGenerators.NormalParams[ValueType]{
				Mean:              0,
				StandardDeviation: ValueType(math.Sqrt(2 / float64(params.InputCount))),
			},
		)

	case randomGenerators.UniformDistribution:
		limit := ValueType(math.Sqrt(6 / float64(params.InputCount)))
		return initializer.randomGenerator.NextUniformValue(randomGenerators.Range[ValueType]{
			Min: -limit,
			Max: limit,
		})

	default:
		panic(fmt.Sprintf("unexpected distribution: %d", initializer.distribution))
	}
}
