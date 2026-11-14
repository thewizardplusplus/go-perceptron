package randomGenerators_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	perceptron "github.com/thewizardplusplus/go-perceptron"
	randomGenerators "github.com/thewizardplusplus/go-perceptron/random-generators"
)

func TestRange_Len(test *testing.T) {
	testRange_Len[float64](test)
	testRange_Len[float32](test)
	testRange_Len[namedFloat64](test)
}

func testRange_Len[ValueType perceptron.ValueConstraint](test *testing.T) {
	test.Helper()

	test.Run(reflect.TypeFor[ValueType]().String(), func(test *testing.T) {
		for _, testCase := range []struct {
			name       string
			valueRange randomGenerators.Range[ValueType]
			expected   ValueType
		}{
			{
				name:       "positive",
				valueRange: randomGenerators.Range[ValueType]{Min: 2, Max: 9},
				expected:   7,
			},
			{
				name:       "negative",
				valueRange: randomGenerators.Range[ValueType]{Min: -9, Max: -2},
				expected:   7,
			},
			{
				name:       "crossing zero",
				valueRange: randomGenerators.Range[ValueType]{Min: -3, Max: 7},
				expected:   10,
			},
			{
				name:       "fractional",
				valueRange: randomGenerators.Range[ValueType]{Min: 0.125, Max: 0.875},
				expected:   0.75,
			},
		} {
			test.Run(testCase.name, func(test *testing.T) {
				actual := testCase.valueRange.Len()

				assert.Equal(test, testCase.expected, actual)
			})
		}
	})
}
