package observer

import "testing"

func TestObserver(t *testing.T) {
	type fields struct {
		subject *Subject
		readers []*Reader
	}
	type args struct {
		context string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "subscribe then message",
			fields: fields{
				subject: NewSubject(),
				readers: []*Reader{
					NewReader("reader1"),
					NewReader("reader2"),
					NewReader("reader3"),
				},
			},
			args: args{context: "observer mode"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, r := range tt.fields.readers {
				tt.fields.subject.Subscribe(r)
			}

			tt.fields.subject.UpdateContext(tt.args.context)
		})
	}
}
