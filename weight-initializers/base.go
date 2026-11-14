package weightInitializers

import (
	perceptron "github.com/thewizardplusplus/go-perceptron"
	randomGenerators "github.com/thewizardplusplus/go-perceptron/random-generators"
)

type base[valueType perceptron.ValueConstraint] struct {
	randomGenerator randomGenerators.Interface[valueType]
	distribution    randomGenerators.Distribution
}

func newBase[valueType perceptron.ValueConstraint](
	randomGenerator randomGenerators.Interface[valueType],
	distribution randomGenerators.Distribution,
) base[valueType] {
	return base[valueType]{
		randomGenerator: randomGenerator,
		distribution:    distribution,
	}
}
