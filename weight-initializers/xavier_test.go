package weightInitializers_test

import (
	"math"
	"reflect"
	"testing"

	perceptron "github.com/thewizardplusplus/go-perceptron"
	randomGeneratorsMocks "github.com/thewizardplusplus/go-perceptron/mocks/github.com/thewizardplusplus/go-perceptron/random-generators"
	randomGenerators "github.com/thewizardplusplus/go-perceptron/random-generators"
	weightInitializers "github.com/thewizardplusplus/go-perceptron/weight-initializers"
)

func TestXavier_InitWeight(test *testing.T) {
	testXavier_InitWeight[float64](test)
	testXavier_InitWeight[float32](test)
	testXavier_InitWeight[namedFloat64](test)
}

func testXavier_InitWeight[ValueType perceptron.ValueConstraint](test *testing.T) {
	test.Helper()

	test.Run(reflect.TypeFor[ValueType]().String(), func(test *testing.T) {
		for _, testCase := range []struct {
			name                string
			distribution        randomGenerators.Distribution
			initWeightParams    perceptron.InitWeightParams
			expectedRandomScale ValueType
		}{
			// normal distribution
			{
				name:                "normal distribution/single input and output",
				distribution:        randomGenerators.NormalDistribution,
				initWeightParams:    perceptron.InitWeightParams{InputCount: 1, OutputCount: 1},
				expectedRandomScale: ValueType(1),
			},
			{
				name:                "normal distribution/equal layer sizes",
				distribution:        randomGenerators.NormalDistribution,
				initWeightParams:    perceptron.InitWeightParams{InputCount: 4, OutputCount: 4},
				expectedRandomScale: ValueType(0.5),
			},
			{
				name:                "normal distribution/expanding layer",
				distribution:        randomGenerators.NormalDistribution,
				initWeightParams:    perceptron.InitWeightParams{InputCount: 2, OutputCount: 4},
				expectedRandomScale: ValueType(math.Sqrt(1.0 / 3)),
			},
			{
				name:                "normal distribution/contracting layer",
				distribution:        randomGenerators.NormalDistribution,
				initWeightParams:    perceptron.InitWeightParams{InputCount: 4, OutputCount: 2},
				expectedRandomScale: ValueType(math.Sqrt(1.0 / 3)),
			},

			// uniform distribution
			{
				name:                "uniform distribution/single input and output",
				distribution:        randomGenerators.UniformDistribution,
				initWeightParams:    perceptron.InitWeightParams{InputCount: 1, OutputCount: 1},
				expectedRandomScale: ValueType(math.Sqrt(3)),
			},
			{
				name:                "uniform distribution/equal layer sizes",
				distribution:        randomGenerators.UniformDistribution,
				initWeightParams:    perceptron.InitWeightParams{InputCount: 12, OutputCount: 12},
				expectedRandomScale: ValueType(0.5),
			},
			{
				name:                "uniform distribution/expanding layer",
				distribution:        randomGenerators.UniformDistribution,
				initWeightParams:    perceptron.InitWeightParams{InputCount: 4, OutputCount: 12},
				expectedRandomScale: ValueType(math.Sqrt(0.375)),
			},
			{
				name:                "uniform distribution/contracting layer",
				distribution:        randomGenerators.UniformDistribution,
				initWeightParams:    perceptron.InitWeightParams{InputCount: 12, OutputCount: 4},
				expectedRandomScale: ValueType(math.Sqrt(0.375)),
			},
		} {
			test.Run(testCase.name, func(test *testing.T) {
				randomGenerator := randomGeneratorsMocks.NewMockInterface[ValueType](test)
				initializer := weightInitializers.NewXavier(randomGenerator, testCase.distribution)

				assertWeight(test, assertWeightParams[ValueType]{
					randomGenerator:     randomGenerator,
					distribution:        testCase.distribution,
					initializer:         initializer,
					initWeightParams:    testCase.initWeightParams,
					expectedRandomScale: testCase.expectedRandomScale,
				})
			})
		}
	})
}
