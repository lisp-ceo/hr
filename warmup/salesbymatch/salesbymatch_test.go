package salesbymatch

import (
	"fmt"
	"testing"
)

func Test_matchingSocks(t *testing.T) {
	tests := []struct {
		socks []int
		want  int
	}{
		{
			[]int{1, 2, 1, 2, 1, 3, 2},
			2,
		},
	}
	for n, tt := range tests {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			if got := matchingSocks(tt.socks); got != tt.want {
				t.Errorf("matchingSocks() = %v, want %v", got, tt.want)
			}
		})
	}
}
