package randomGenerators

import (
	perceptron "github.com/thewizardplusplus/go-perceptron"
)

type Interface[ValueType perceptron.ValueConstraint] interface {
	NextNormalValue(params NormalParams[ValueType]) ValueType
	NextUniformValue(r Range[ValueType]) ValueType
}
