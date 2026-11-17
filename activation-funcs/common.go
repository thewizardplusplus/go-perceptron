package activationFuncs

import (
	"github.com/thewizardplusplus/go-perceptron"
)

type selectionCases[ValueType perceptron.ValueConstraint] struct {
	leftHandValue  ValueType
	rightHandValue ValueType
}

func selectByValue[ValueType perceptron.ValueConstraint](
	value ValueType,
	selectionCases selectionCases[ValueType],
) (selectedValue ValueType, isSelected bool) {
	switch {
	case value < 0:
		return selectionCases.leftHandValue, true

	case value > 0:
		return selectionCases.rightHandValue, true

	case value != value: // NaN
		return value, true

	default:
		return 0, false
	}
}
