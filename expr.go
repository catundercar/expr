package expr

type Evaluator interface {
	Eval() (float64, error)
}
