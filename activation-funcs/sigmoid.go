package activationFuncs

import (
	"math"

	"github.com/thewizardplusplus/go-perceptron"
)

type Sigmoid[ValueType perceptron.ValueConstraint] struct{}

func (s Sigmoid[ValueType]) Apply(value ValueType) ValueType {
	if value >= 0 {
		return ValueType(1 / (1 + math.Exp(-float64(value))))
	}

	exponent := math.Exp(float64(value))
	return ValueType(exponent / (exponent + 1))
}

func (s Sigmoid[ValueType]) ApplyDerivative(value ValueType) ValueType {
	exponent := math.Exp(-math.Abs(float64(value)))
	return ValueType(exponent / math.Pow(exponent+1, 2)) //nolint:staticcheck
}
