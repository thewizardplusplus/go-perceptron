package randomGenerators_test

import (
	"fmt"
	"math/rand/v2"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	perceptron "github.com/thewizardplusplus/go-perceptron"
	randomGenerators "github.com/thewizardplusplus/go-perceptron/random-generators"
)

const (
	sampleCount = 100
)

func TestGenerator_NextNormalValue(test *testing.T) {
	testGenerator_NextNormalValue[float64](test)
	testGenerator_NextNormalValue[float32](test)
	testGenerator_NextNormalValue[namedFloat64](test)
}

func testGenerator_NextNormalValue[ValueType perceptron.ValueConstraint](test *testing.T) {
	test.Helper()

	test.Run(reflect.TypeFor[ValueType]().String(), func(test *testing.T) {
		for _, testCase := range []struct {
			name   string
			params randomGenerators.NormalParams[ValueType]
		}{
			{
				name: "standard",
				params: randomGenerators.NormalParams[ValueType]{
					Mean:              0,
					StandardDeviation: 1,
				},
			},
			{
				name: "positive mean",
				params: randomGenerators.NormalParams[ValueType]{
					Mean:              5,
					StandardDeviation: 1,
				},
			},
			{
				name: "negative mean",
				params: randomGenerators.NormalParams[ValueType]{
					Mean:              -5,
					StandardDeviation: 1,
				},
			},
			{
				name: "scaled",
				params: randomGenerators.NormalParams[ValueType]{
					Mean:              2.5,
					StandardDeviation: 3.5,
				},
			},
			{
				name: "small deviation",
				params: randomGenerators.NormalParams[ValueType]{
					Mean:              0,
					StandardDeviation: 0.125,
				},
			},
		} {
			test.Run(testCase.name, func(test *testing.T) {
				generator := randomGenerators.NewGenerator[ValueType](randomSource())
				reference := rand.New(randomSource())
				for sampleIndex := range sampleCount {
					actual := generator.NextNormalValue(testCase.params)

					mean, standardDeviation :=
						float64(testCase.params.Mean), float64(testCase.params.StandardDeviation)
					expected := ValueType(reference.NormFloat64()*standardDeviation + mean)

					require.Equal(test, expected, actual, "sample #%d", sampleIndex)
				}
			})
		}
	})
}

func TestGenerator_NextUniformValue(test *testing.T) {
	testGenerator_NextUniformValue[float64](test)
	testGenerator_NextUniformValue[float32](test)
	testGenerator_NextUniformValue[namedFloat64](test)
}

func testGenerator_NextUniformValue[ValueType perceptron.ValueConstraint](test *testing.T) {
	test.Helper()

	test.Run(reflect.TypeFor[ValueType]().String(), func(test *testing.T) {
		for _, testCase := range []struct {
			name       string
			valueRange randomGenerators.Range[ValueType]
		}{
			{
				name:       "positive",
				valueRange: randomGenerators.Range[ValueType]{Min: 2, Max: 9},
			},
			{
				name:       "negative",
				valueRange: randomGenerators.Range[ValueType]{Min: -9, Max: -2},
			},
			{
				name:       "crossing zero",
				valueRange: randomGenerators.Range[ValueType]{Min: -3, Max: 7},
			},
			{
				name:       "fractional",
				valueRange: randomGenerators.Range[ValueType]{Min: 0.125, Max: 0.875},
			},
		} {
			test.Run(testCase.name, func(test *testing.T) {
				generator := randomGenerators.NewGenerator[ValueType](randomSource())
				reference := rand.New(randomSource())
				for sampleIndex := range sampleCount {
					actual := generator.NextUniformValue(testCase.valueRange)

					min, max := testCase.valueRange.Min, testCase.valueRange.Max
					expected := nextReferenceUniform[ValueType](reference)*(max-min) + min

					require.Equal(test, expected, actual, "sample #%d", sampleIndex)
					require.GreaterOrEqual(test, actual, testCase.valueRange.Min, "sample #%d", sampleIndex)
					require.LessOrEqual(test, actual, testCase.valueRange.Max, "sample #%d", sampleIndex)
				}
			})
		}
	})
}

func nextReferenceUniform[ValueType perceptron.ValueConstraint](generator *rand.Rand) ValueType {
	var randomValue ValueType
	switch reflect.TypeFor[ValueType]().Kind() {
	case reflect.Float32:
		randomValue = ValueType(generator.Float32())

	case reflect.Float64:
		randomValue = ValueType(generator.Float64())

	default:
		panic(fmt.Sprintf("unexpected type: %T", randomValue))
	}

	return randomValue
}
