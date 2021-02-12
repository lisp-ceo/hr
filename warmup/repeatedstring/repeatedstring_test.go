package repeatedstring

import "testing"

func Test_repeatedString(t *testing.T) {
	type args struct {
		s string
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "default test",
			args: args{
				s: "aba",
				n: 10,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedString(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("repeatedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
