package perceptron_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	perceptron "github.com/thewizardplusplus/go-perceptron"
	perceptronMocks "github.com/thewizardplusplus/go-perceptron/mocks/github.com/thewizardplusplus/go-perceptron"
)

func TestWeightInitializerFunc_InitWeight(test *testing.T) {
	type namedFloat64 float64

	testWeightInitializerFunc_InitWeight[float64](test)
	testWeightInitializerFunc_InitWeight[float32](test)
	testWeightInitializerFunc_InitWeight[namedFloat64](test)
}

func testWeightInitializerFunc_InitWeight[ValueType perceptron.ValueConstraint](test *testing.T) {
	test.Helper()

	test.Run(reflect.TypeFor[ValueType]().String(), func(test *testing.T) {
		for _, testCase := range []struct {
			name             string
			initWeightParams perceptron.InitWeightParams
			expectedWeight   ValueType
		}{
			{
				name:             "positive weight",
				initWeightParams: perceptron.InitWeightParams{InputCount: 2, OutputCount: 4},
				expectedWeight:   ValueType(0.375),
			},
			{
				name:             "negative weight",
				initWeightParams: perceptron.InitWeightParams{InputCount: 4, OutputCount: 2},
				expectedWeight:   ValueType(-0.375),
			},
			{
				name:             "zero weight",
				initWeightParams: perceptron.InitWeightParams{InputCount: 1, OutputCount: 1},
				expectedWeight:   ValueType(0),
			},
		} {
			test.Run(testCase.name, func(test *testing.T) {
				mockInitializer := perceptronMocks.NewMockWeightInitializer[ValueType](test)
				mockInitializer.EXPECT().
					InitWeight(testCase.initWeightParams).
					Return(testCase.expectedWeight).
					Once()

				initializer := perceptron.WeightInitializerFunc[ValueType](mockInitializer.InitWeight)
				actualWeight := initializer.InitWeight(testCase.initWeightParams)

				require.Equal(test, testCase.expectedWeight, actualWeight)
			})
		}
	})
}
