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

func TestHe_InitWeight(test *testing.T) {
	testHe_InitWeight[float64](test)
	testHe_InitWeight[float32](test)
	testHe_InitWeight[namedFloat64](test)
}

func testHe_InitWeight[ValueType perceptron.ValueConstraint](test *testing.T) {
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
				name:                "normal distribution/single input",
				distribution:        randomGenerators.NormalDistribution,
				initWeightParams:    perceptron.InitWeightParams{InputCount: 1},
				expectedRandomScale: ValueType(math.Sqrt(2)),
			},
			{
				name:                "normal distribution/multiple inputs",
				distribution:        randomGenerators.NormalDistribution,
				initWeightParams:    perceptron.InitWeightParams{InputCount: 8},
				expectedRandomScale: ValueType(0.5),
			},

			// uniform distribution
			{
				name:                "uniform distribution/single input",
				distribution:        randomGenerators.UniformDistribution,
				initWeightParams:    perceptron.InitWeightParams{InputCount: 1},
				expectedRandomScale: ValueType(math.Sqrt(6)),
			},
			{
				name:                "uniform distribution/multiple inputs",
				distribution:        randomGenerators.UniformDistribution,
				initWeightParams:    perceptron.InitWeightParams{InputCount: 24},
				expectedRandomScale: ValueType(0.5),
			},
		} {
			test.Run(testCase.name, func(test *testing.T) {
				randomGenerator := randomGeneratorsMocks.NewMockInterface[ValueType](test)
				initializer := weightInitializers.NewHe(randomGenerator, testCase.distribution)

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
