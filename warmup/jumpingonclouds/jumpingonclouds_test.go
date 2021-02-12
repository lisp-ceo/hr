package jumpingonclouds

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_minJumps(t *testing.T) {
	type args struct {
		clouds []int32
	}
	tests := []struct {
		args args
		want int32
	}{
		{
			args: args{
				[]int32{0, 0, 1, 0, 0, 1, 0},
			},
			want: 4,
		},
		{
			args: args{
				[]int32{0, 0, 0, 1, 0, 0},
			},
			want: 3,
		},
		{
			args: args{
				[]int32{
					0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0,
					0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0,
				},
			},
			want: 28,
		},
	}
	for n, tt := range tests {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			if got := jumpingOnClouds(tt.args.clouds); got != tt.want {
				t.Errorf("jumpingOnClouds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allPaths(t *testing.T) {
	type args struct {
		max2SpaceJumps int32
		spaces         int32
		clouds []int32
	}
	tests := []struct {
		name string
		args args
		want [][]int32
	}{
		{
			name: "generates all subtrees",
			args: args{
				max2SpaceJumps: 1,
				spaces:         3,
				clouds: []int32{0,0,0},
			},
			want: [][]int32{
				{2, 1},
				{1, 2},
				{1, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPaths(tt.args.max2SpaceJumps, tt.args.spaces, tt.args.clouds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSafePath(t *testing.T) {
	type args struct {
		path  []int32
		clouds []int32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "safe path",
			args: args{
				path: []int32{1, 1, 1, 2},
				clouds: []int32{0, 0, 0, 0, 1, 0},
			},
			want: true,
		},
		{
			name: "unsafe path",
			args: args{
				path: []int32{1, 1, 1, 1, 1},
				clouds: []int32{0, 0, 0, 0, 1, 0},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSafePath(tt.args.path, tt.args.clouds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterInvalid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minPathLength(t *testing.T) {
	type args struct {
		paths [][]int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "we calculate the length of smallest path as the minimum number of jumps",
			args: args{
				paths: [][]int32{
					{2, 2},
					{1, 1, 2},
					{1, 2, 1},
					{1, 1, 2},
					{1, 1, 1, 1},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathLength(tt.args.paths); got != tt.want {
				t.Errorf("minPathLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
