package newyearchaos

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minimumBribes(t *testing.T) {
	type args struct {
		q []int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "too chaotic",
			args: args{
				[]int32{4,1,2,3},
			},
			want: "Too chaotic\n",
		},
		{
			name: "4 swaps",
			args: args{
				[]int32{1,2,5,3,4,7,8,6},
			},
			want: "4\n",
		},
		{
			name: "7 swaps",
			args: args{
				[]int32{1,2,5,3,7,8,6,4},
			},
			want: "7\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.Buffer{}
			stdout = &buf
			minimumBribes(tt.args.q)
			assert.Equal(t, tt.want, buf.String())
		})
	}
}
