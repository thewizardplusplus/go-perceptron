package activationFuncs

import (
	"fmt"

	"github.com/samber/mo"
	"github.com/thewizardplusplus/go-perceptron"
)

type ZeroMode int

const (
	LeftHandZeroMode ZeroMode = iota
	RightHandZeroMode
	MeanZeroMode
)

type SurrogateGradientSource[ValueType perceptron.ValueConstraint] interface {
	ApplyDerivative(value ValueType) ValueType
}

type HeavisideOptions[ValueType perceptron.ValueConstraint] struct {
	ZeroMode                  ZeroMode
	ConstantSurrogateGradient ValueType
	SurrogateGradientSource   mo.Option[SurrogateGradientSource[ValueType]]
}

type Heaviside[ValueType perceptron.ValueConstraint] struct {
	options HeavisideOptions[ValueType]
}

func NewHeaviside[ValueType perceptron.ValueConstraint](
	options HeavisideOptions[ValueType],
) (Heaviside[ValueType], error) {
	if options.ZeroMode != LeftHandZeroMode &&
		options.ZeroMode != RightHandZeroMode &&
		options.ZeroMode != MeanZeroMode {
		return Heaviside[ValueType]{}, fmt.Errorf(
			"unexpected Heaviside zero mode: %d",
			options.ZeroMode,
		)
	}

	return Heaviside[ValueType]{options: options}, nil
}

func (h Heaviside[ValueType]) Apply(value ValueType) ValueType {
	const leftHandValue, rightHandValue = 0, 1

	if selectedValue, isSelected := selectByValue(value, selectionCases[ValueType]{
		leftHandValue:  leftHandValue,
		rightHandValue: rightHandValue,
	}); isSelected {
		return selectedValue
	}

	// value == 0
	switch h.options.ZeroMode {
	case LeftHandZeroMode:
		return leftHandValue

	case RightHandZeroMode:
		return rightHandValue

	case MeanZeroMode:
		return (leftHandValue + rightHandValue) / 2.0

	default:
		panic(fmt.Sprintf("unexpected Heaviside zero mode: %d", h.options.ZeroMode))
	}
}

func (h Heaviside[ValueType]) ApplyDerivative(value ValueType) ValueType {
	if surrogateGradientSource, isPresent := h.options.SurrogateGradientSource.Get(); isPresent {
		return surrogateGradientSource.ApplyDerivative(value)
	}

	if value != value { // NaN
		return value
	}

	return h.options.ConstantSurrogateGradient
}
