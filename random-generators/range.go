package randomGenerators

import (
	perceptron "github.com/thewizardplusplus/go-perceptron"
)

type Range[ValueType perceptron.ValueConstraint] struct {
	Min ValueType
	Max ValueType
}

func (r Range[ValueType]) Len() ValueType {
	return r.Max - r.Min
}

func (r Range[ValueType]) LenAsInt64() int64 {
	return int64(r.Max) - int64(r.Min)
}
