package weightInitializers

import (
	"fmt"
	"math"

	perceptron "github.com/thewizardplusplus/go-perceptron"
	randomGenerators "github.com/thewizardplusplus/go-perceptron/random-generators"
)

var _ perceptron.WeightInitializer[float64] = Xavier[float64]{}

type Xavier[ValueType perceptron.ValueConstraint] struct {
	base[ValueType]
}

func NewXavier[ValueType perceptron.ValueConstraint](
	randomGenerator randomGenerators.Interface[ValueType],
	distribution randomGenerators.Distribution,
) Xavier[ValueType] {
	return Xavier[ValueType]{
		base: newBase(randomGenerator, distribution),
	}
}

func (initializer Xavier[ValueType]) InitWeight(params perceptron.InitWeightParams) ValueType {
	inputCount, outputCount := float64(params.InputCount), float64(params.OutputCount)

	switch initializer.distribution {
	case randomGenerators.NormalDistribution:
		return initializer.randomGenerator.NextNormalValue(
			randomGenerators.NormalParams[ValueType]{
				Mean:              0,
				StandardDeviation: ValueType(math.Sqrt(2 / (inputCount + outputCount))),
			},
		)

	case randomGenerators.UniformDistribution:
		limit := ValueType(math.Sqrt(6 / (inputCount + outputCount)))
		return initializer.randomGenerator.NextUniformValue(randomGenerators.Range[ValueType]{
			Min: -limit,
			Max: limit,
		})

	default:
		panic(fmt.Sprintf("unexpected distribution: %d", initializer.distribution))
	}
}
