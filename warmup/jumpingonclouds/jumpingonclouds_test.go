package jumpingonclouds

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
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
			want: 3,
		},
	}
	for n, tt := range tests {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			if got := minJumps(tt.args.clouds); got != tt.want {
				t.Errorf("minJumps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_easyGenPath(t *testing.T) {
	type args struct {
		spaces         int32
		max2spaceJumps int32
	}
	tests := []struct {
		name string
		args args
		want [][]int32
	}{
		//{
		//	name: "non-recursive: too many 2space jumps for the spaces",
		//	args: args{
		//		spaces:   4,
		//		max2spaceJumps: 3,
		//	},
		//	want: [][]int32{
		//		{1, 1, 1, 1},
		//	},
		//},
		//{
		//	name: "non-recursive: 0 2space jumps for the spaces",
		//	args: args{
		//		spaces:   2,
		//		max2spaceJumps: 0,
		//	},
		//	want: [][]int32{
		//		{1, 1},
		//	},
		//},
		//{
		//	name: "non-recursive: 1 2space jump for the 2 spaces",
		//	args: args{
		//		spaces:   2,
		//		max2spaceJumps: 1,
		//	},
		//	want: [][]int32{
		//		{2},
		//	},
		//},
		{
			name: "recursive: single recursive case",
			args: args{
				spaces:         3,
				max2spaceJumps: 1,
			},
			want: [][]int32{
				{2, 1},
				{1, 2},
			},
		},
	}
	for n, tt := range tests {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			require.Equal(t, tt.want, genPaths(tt.args.max2spaceJumps, tt.args.spaces))
		})
	}
}

func Test_allPaths(t *testing.T) {
	type args struct {
		max2SpaceJumps int32
		spaces         int32
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
			if got := allPaths(tt.args.max2SpaceJumps, tt.args.spaces); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterInvalid(t *testing.T) {
	type args struct {
		paths  [][]int32
		clouds []int32
	}
	tests := []struct {
		name string
		args args
		want [][]int32
	}{
		{
			name: "removes invalid paths",
			args: args{
				paths: [][]int32{
					{1, 1, 1, 1, 1},
					{1, 1, 1, 2},
				},
				clouds: []int32{0, 0, 0, 0, 1, 0},
			},
			want: [][]int32{
				{1, 1, 1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterInvalid(tt.args.paths, tt.args.clouds); !reflect.DeepEqual(got, tt.want) {
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
