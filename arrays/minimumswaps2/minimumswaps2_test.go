package minimumswaps2

import "testing"

func Test_minimumSwaps(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "second test case",
			args: args{
				[]int32{1,2,3,4,5},
			},
			want: 0,
		},
		{
			name: "first test case",
			args: args{
				[]int32{4, 3, 1, 2},
			},
			want: 3,
		},
		{
			name: "second test case",
			args: args{
				[]int32{2, 3, 4, 1, 5},
			},
			want: 3,
		},
		{
			name: "third test case",
			args: args{
				[]int32{2, 31, 1, 38, 29, 5, 44, 6, 12, 18, 39, 9, 48, 49, 13, 11, 7, 27, 14, 33, 50, 21, 46, 23, 15, 26, 8, 47, 40, 3, 32, 22, 34, 42, 16, 41, 24, 10, 4, 28, 36, 30, 37, 35, 20, 17, 45, 43, 25, 19},
			},
			want: 46,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSwaps(tt.args.arr); got != tt.want {
				t.Errorf("minimumSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
