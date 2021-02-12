package twodarray

import "testing"

func Test_hourglassSum(t *testing.T) {
	type args struct {
		arr [][]int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "1st test case",
			args: args{
				arr: [][]int32{
					{1,1,1,0,0,0},
					{0,1,0,0,0,0},
					{1,1,1,0,0,0},
					{0,0,2,4,4,0},
					{0,0,0,2,0,0},
					{0,0,1,2,4,0},
				},
			},
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hourglassSum(tt.args.arr); got != tt.want {
				t.Errorf("hourglassSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
