package abstractfactory

import "testing"

func TestSaveOrderMain(t *testing.T) {
	type args struct {
		factory DAOFactory
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "rdb save main then success",
			args: args{factory: &RDBDAOFactory{}},
			want: "rdb main save",
		},
		{
			name: "xml save main then success",
			args: args{factory: &XMLDAOFactory{}},
			want: "xml main save",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.factory.CreateOrderMainDAO().SaveOrderMain(); got != tt.want {
				t.Errorf("SaveOrderMain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveOrderDetail(t *testing.T) {
	type args struct {
		factory DAOFactory
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "rdb save detail then success",
			args: args{factory: &RDBDAOFactory{}},
			want: "rdb detail save",
		},
		{
			name: "xml save detail then success",
			args: args{factory: &XMLDAOFactory{}},
			want: "xml detail save",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.factory.CreateOrderDetailDAO().SaveOrderDetail(); got != tt.want {
				t.Errorf("SaveOrderDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}
