package randomGenerators

import (
	perceptron "github.com/thewizardplusplus/go-perceptron"
)

type NormalParams[ValueType perceptron.ValueConstraint] struct {
	Mean              ValueType
	StandardDeviation ValueType
}

type Interface[ValueType perceptron.ValueConstraint] interface {
	NextNormalValue(params NormalParams[ValueType]) ValueType
	NextUniformValue(r Range[ValueType]) ValueType
}
