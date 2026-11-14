package perceptron

import (
	"golang.org/x/exp/constraints"
)

type ValueConstraint interface {
	constraints.Float
}
