package perceptron

type ActivationFunc[ValueType ValueConstraint] interface {
	Apply(value ValueType) ValueType
	ApplyDerivative(value ValueType) ValueType
}

type ActivationFuncWrapper[ValueType ValueConstraint] struct {
	FuncValue           func(value ValueType) ValueType
	DerivativeFuncValue func(value ValueType) ValueType
}

func (wrapper ActivationFuncWrapper[ValueType]) Apply(value ValueType) ValueType {
	return wrapper.FuncValue(value)
}

func (wrapper ActivationFuncWrapper[ValueType]) ApplyDerivative(value ValueType) ValueType {
	return wrapper.DerivativeFuncValue(value)
}
