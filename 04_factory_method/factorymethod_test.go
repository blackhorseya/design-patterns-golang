package factorymethod

import "testing"

func TestOperator(t *testing.T) {
	type args struct {
		factory OperatorFactory
		a       int
		b       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "plus 1 2 then 3",
			args: args{factory: PlusOperatorFactory{}, a: 1, b: 2},
			want: 3,
		},
		{
			name: "minus 4 2 then 2",
			args: args{factory: MinusOperatorFactory{}, a: 4, b: 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := tt.args.factory.Create()
			op.SetA(tt.args.a)
			op.SetB(tt.args.b)

			if got := op.Result(); got != tt.want {
				t.Errorf("Operator() = %v, want %v", got, tt.want)
			}
		})
	}
}
