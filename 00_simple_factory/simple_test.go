package simplefactory

import "testing"

func Test_Adventurer_Say(t *testing.T) {
	type field struct {
		a Adventurer
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields field
		args   args
		want   string
	}{
		{
			name:   "archer tom then success",
			fields: field{a: NewAdventurer("archer")},
			args:   args{name: "Tom"},
			want:   "Hi Tom, I'm an archer",
		},
		{
			name:   "warrior tom then success",
			fields: field{a: NewAdventurer("warrior")},
			args:   args{name: "Tom"},
			want:   "Hi Tom, I'm a warrior",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.a.Say(tt.args.name); got != tt.want {
				t.Errorf("Say() = %v, want %v", got, tt.want)
			}
		})
	}
}
