package expr

import (
	"fmt"
	"github.com/catundercar/expr/pkg/container"
	"testing"
)

func TestNewRPN(t *testing.T) {
	ev, err := NewRPN("1+1*2-3+2/1+1-1+1/1") // 1
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ev)
	fmt.Println(ev.Eval())

	expr := "1+2*(1+2)-10"
	fmt.Println(expr)
	ev, err = NewRPN(expr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ev)
	fmt.Println(ev.Eval())
}

func TestRPN_Eval(t *testing.T) {
	type fields struct {
		queue container.Queue[any]
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpn := &RPN{}
			got, err := rpn.Eval()
			if (err != nil) != tt.wantErr {
				t.Errorf("Eval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Eval() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestRPN_String(t *testing.T) {
//	type fields struct {
//		queue container.Queue
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			rpn := &RPN{
//				queue: tt.fields.queue,
//			}
//			if got := rpn.String(); got != tt.want {
//				t.Errorf("String() = %v, want %v", got, tt.want)
//			}
//		})
//	}
////}
//
//func Test_binaryOpFunc(t *testing.T) {
//	type args struct {
//		fn func(v1, v2 float64) float64
//	}
//	tests := []struct {
//		name string
//		args args
//		want func(args ...float64) (float64, error)
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := binaryOpFunc(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("binaryOpFunc() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_hasPrecedence(t *testing.T) {
//	type args struct {
//		op1 string
//		op2 string
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := hasPrecedence(tt.args.op1, tt.args.op2); got != tt.want {
//				t.Errorf("hasPrecedence() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
