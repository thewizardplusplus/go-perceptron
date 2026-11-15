package weightInitializers_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	perceptron "github.com/thewizardplusplus/go-perceptron"
	randomGeneratorsMocks "github.com/thewizardplusplus/go-perceptron/mocks/github.com/thewizardplusplus/go-perceptron/random-generators"
	randomGenerators "github.com/thewizardplusplus/go-perceptron/random-generators"
)

type namedFloat64 float64

type assertWeightParams[ValueType perceptron.ValueConstraint] struct {
	randomGenerator     *randomGeneratorsMocks.MockInterface[ValueType]
	distribution        randomGenerators.Distribution
	initializer         perceptron.WeightInitializer[ValueType]
	initWeightParams    perceptron.InitWeightParams
	expectedRandomScale ValueType
}

func assertWeight[ValueType perceptron.ValueConstraint](
	test *testing.T,
	params assertWeightParams[ValueType],
) {
	test.Helper()

	var expectedWeight ValueType
	switch params.distribution {
	case randomGenerators.NormalDistribution:
		expectedWeight = ValueType(-0.375)

		params.randomGenerator.EXPECT().
			NextNormalValue(randomGenerators.NormalParams[ValueType]{
				Mean:              0,
				StandardDeviation: params.expectedRandomScale,
			}).
			Return(expectedWeight).
			Once()

	case randomGenerators.UniformDistribution:
		expectedWeight = ValueType(0.375)

		params.randomGenerator.EXPECT().
			NextUniformValue(randomGenerators.Range[ValueType]{
				Min: -params.expectedRandomScale,
				Max: params.expectedRandomScale,
			}).
			Return(expectedWeight).
			Once()
	}

	actualWeight := params.initializer.InitWeight(params.initWeightParams)

	require.Equal(test, expectedWeight, actualWeight)
}
