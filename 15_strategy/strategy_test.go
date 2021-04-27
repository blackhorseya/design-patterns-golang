package strategy

import (
	"testing"
)

func TestPayByCash(t *testing.T) {
	type args struct {
		name     string
		cardID   string
		money    int
		strategy PaymentStrategy
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ada pay 123 by cash then success",
			args: args{name: "ada", cardID: "", money: 123, strategy: &Cash{}},
			want: "Pay $123 to ada by cash",
		},
		{
			name: "bob pay 888 by bank then success",
			args: args{name: "bob", cardID: "0002", money: 888, strategy: &Bank{}},
			want: "Pay $888 to bob by bank account 0002",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payment := NewPayment(tt.args.name, tt.args.cardID, tt.args.money, tt.args.strategy)

			if got := payment.Pay(); got != tt.want {
				t.Errorf("Pay() = %v, want %v", got, tt.want)
			}
		})
	}
}
