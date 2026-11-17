package activationFuncs

import (
	"math"

	"github.com/thewizardplusplus/go-perceptron"
)

type Tanh[ValueType perceptron.ValueConstraint] struct{}

func (t Tanh[ValueType]) Apply(value ValueType) ValueType {
	return ValueType(math.Tanh(float64(value)))
}

func (t Tanh[ValueType]) ApplyDerivative(value ValueType) ValueType {
	exponent := math.Exp(-2 * math.Abs(float64(value)))
	return ValueType(4 * exponent / math.Pow(1+exponent, 2)) //nolint:staticcheck
}
