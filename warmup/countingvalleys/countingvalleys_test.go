package countingvalleys

import (
	"fmt"
	"testing"
)

func Test_countValleys(t *testing.T) {
	type args struct {
		steps int
		p     string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				steps: 8,
				p:     "UDDDUDUU",
			},
			want: 1,
		},
		{
			args: args{
				steps: 12,
				p:     "DDUUDDUDUUUD",
			},
			want: 2,
		},
	}
	for n, tt := range tests {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			if got := countValleys(tt.args.steps, tt.args.p); got != tt.want {
				t.Errorf("countValleys() = %v, want %v", got, tt.want)
			}
		})
	}
}
