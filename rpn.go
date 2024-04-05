package expr

import (
	"fmt"
	"github.com/catundercar/expr/pkg/container"
	"strconv"
)

type RPN struct {
	queue container.Queue[*token]
}

func (rpn *RPN) String() string {
	return fmt.Sprintf("%s", rpn.queue)
}

func (rpn *RPN) Eval() (float64, error) {
	stack := container.NewStack[float64]()
	for !rpn.queue.IsEmpty() {
		tk := rpn.queue.Pop()
		switch tk.kind {
		case tokenLiteral:
			v, err := strconv.ParseFloat(string(tk.byteValue), 64)
			if err != nil {
				return 0, fmt.Errorf("cannot parse value: %s err: %w", string(tk.byteValue), err)
			}
			stack.Push(v)
		default:
			v1, v2 := stack.Pop(), stack.Pop()
			f, err := op[tk.delimValue](v2, v1)
			if err != nil {
				return 0, fmt.Errorf("cannot evaluate op: %s err: %w", string(tk.delimValue), err)
			}
			stack.Push(f)
		}
	}
	return stack.Pop(), nil
}

func NewRPN(expression string) (Evaluator, error) {
	tokens, err := Parse(expression)
	if err != nil {
		return nil, err
	}

	stack := container.NewStack[*token]()
	queue := container.NewQueue[*token]()

	for i := range tokens {
		tk := tokens[i]
		switch tk.kind {
		case tokenLParen:
			stack.Push(&tk)
		case tokenRParen:
			for {
				if stack.IsEmpty() {
					return nil, fmt.Errorf("mismatched parenthese")
				}
				opToken := stack.Pop()
				if opToken.kind == tokenLParen {
					break // 左括号弹出后，停止弹出
				}
				queue.Push(opToken)
			}
		case tokenLiteral:
			queue.Push(&tk)
		case tokenOperator:
			// 先弹出优先级大于等于 tk的.
			for {
				if stack.IsEmpty() {
					break
				}
				opToken := stack.Peek()
				if opToken.kind == tokenLParen || !(opPriority[opToken.delimValue] >= opPriority[tk.delimValue]) {
					break // 左括号弹出后，停止弹出
				}
				queue.Push(stack.Pop())
			}
			stack.Push(&tk)
		default:
			return nil, fmt.Errorf("unknown token kind: %v", tk.kind)
		}
	}

	for !stack.IsEmpty() {
		queue.Push(stack.Pop())
	}
	return &RPN{queue: queue}, nil
}

var opPriority = map[byte]int{
	'+': 0, '-': 0,
	'*': 1, '/': 1,
	'(': 2, ')': 2,
}

var op = map[byte]func(args ...float64) (float64, error){
	'+': binaryOpFunc(func(v1, v2 float64) float64 {
		return v1 + v2
	}),
	'-': binaryOpFunc(func(v1, v2 float64) float64 {
		return v1 - v2
	}),
	'*': binaryOpFunc(func(v1, v2 float64) float64 {
		return v1 * v2
	}),
	'/': binaryOpFunc(func(v1, v2 float64) float64 {
		return v1 / v2
	}),
}

func binaryOpFunc(fn func(v1, v2 float64) float64) func(args ...float64) (float64, error) {
	return func(args ...float64) (float64, error) {
		if len(args) != 2 {
			return 0, nil
		}
		return fn(args[0], args[1]), nil
	}
}
