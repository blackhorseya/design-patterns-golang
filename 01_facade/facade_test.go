package facade

import "testing"

func TestEcho(t *testing.T) {
	type fields struct {
		api API
	}
	type args struct {
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "api echo then success",
			fields: fields{api: NewAPIImpl()},
			args:   args{},
			want:   "A running\nB running",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.api.Echo(); got != tt.want {
				t.Errorf("Echo() = %v, want %v", got, tt.want)
			}
		})
	}
}
