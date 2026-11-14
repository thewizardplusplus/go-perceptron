package perceptron

type InitWeightParams struct {
	InputCount  int
	OutputCount int
}

type WeightInitializer[ValueType ValueConstraint] interface {
	InitWeight(params InitWeightParams) ValueType
}

type WeightInitializerFunc[ValueType ValueConstraint] func(params InitWeightParams) ValueType

func (f WeightInitializerFunc[ValueType]) InitWeight(params InitWeightParams) ValueType {
	return f(params)
}
