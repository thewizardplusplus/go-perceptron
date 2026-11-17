package activationFuncs

import (
	"errors"
	"fmt"
	"math"

	"github.com/samber/mo"
	"github.com/thewizardplusplus/go-perceptron"
)

type ZeroDerivativeMode int

const (
	LeftHandZeroDerivativeMode ZeroDerivativeMode = iota
	RightHandZeroDerivativeMode
	MeanZeroDerivativeMode
	CustomZeroDerivativeMode
)

type ReLUOptions[ValueType perceptron.ValueConstraint] struct {
	NegativeSlope        ValueType // 0 for ReLU, (0, 1) for Leaky ReLU, 1 for Identity
	ZeroDerivativeMode   ZeroDerivativeMode
	CustomZeroDerivative mo.Option[ValueType] // must be in [NegativeSlope, 1]
}

type ReLU[ValueType perceptron.ValueConstraint] struct {
	options ReLUOptions[ValueType]
}

func NewReLU[ValueType perceptron.ValueConstraint](
	options ReLUOptions[ValueType],
) (ReLU[ValueType], error) {
	// use this form to check for NaN
	if !(options.NegativeSlope >= 0 && options.NegativeSlope <= 1) { //nolint:staticcheck
		return ReLU[ValueType]{}, errors.New("negative slope is out of range")
	}

	if options.ZeroDerivativeMode != LeftHandZeroDerivativeMode &&
		options.ZeroDerivativeMode != RightHandZeroDerivativeMode &&
		options.ZeroDerivativeMode != MeanZeroDerivativeMode &&
		options.ZeroDerivativeMode != CustomZeroDerivativeMode {
		return ReLU[ValueType]{}, fmt.Errorf(
			"unexpected ReLU derivative mode: %d",
			options.ZeroDerivativeMode,
		)
	}

	if options.ZeroDerivativeMode == CustomZeroDerivativeMode {
		zeroDerivative, isPresent := options.CustomZeroDerivative.Get()
		if !isPresent {
			return ReLU[ValueType]{}, errors.New(
				"zero derivative is required for the custom ReLU derivative mode",
			)
		}
		// use this form to check for NaN
		if !(zeroDerivative >= options.NegativeSlope && zeroDerivative <= 1) { //nolint:staticcheck
			return ReLU[ValueType]{}, errors.New("zero derivative is out of range")
		}
	}

	return ReLU[ValueType]{options: options}, nil
}

func (r ReLU[ValueType]) Apply(value ValueType) ValueType {
	if value > 0 {
		return value
	}

	// separate check for infinity
	if r.options.NegativeSlope == 0 && math.IsInf(float64(value), -1) {
		return 0
	}

	return r.options.NegativeSlope * value
}

func (r ReLU[ValueType]) ApplyDerivative(value ValueType) ValueType {
	leftHandDerivative, rightHandDerivative := r.options.NegativeSlope, ValueType(1)

	if selectedDerivative, isSelected := selectByValue(value, selectionCases[ValueType]{
		leftHandValue:  leftHandDerivative,
		rightHandValue: rightHandDerivative,
	}); isSelected {
		return selectedDerivative
	}

	// value == 0
	switch r.options.ZeroDerivativeMode {
	case LeftHandZeroDerivativeMode:
		return leftHandDerivative

	case RightHandZeroDerivativeMode:
		return rightHandDerivative

	case MeanZeroDerivativeMode:
		return (leftHandDerivative + rightHandDerivative) / 2

	case CustomZeroDerivativeMode:
		return r.options.CustomZeroDerivative.MustGet()

	default:
		panic(fmt.Sprintf("unexpected ReLU derivative mode: %d", r.options.ZeroDerivativeMode))
	}
}
