package mink

import (
	"testing"
)

func Test_partion(t *testing.T) {
	type args struct {
		input []int
		l     int
		r     int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				input: []int{1, 2, 3, 4},
				l:     0,
				r:     3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := partion(tt.args.input, tt.args.l, tt.args.r)
			t.Log(tt.args.input, j)
		})
	}
}

func Test_bubbleSort(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				input: []int{1, 4, 2, 1, 4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bubbleSort(tt.args.input)
		})
	}
}

func Test_quickSort(t *testing.T) {
	type args struct {
		input []int
		l     int
		r     int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				input: []int{1, 4, 2, 1, 4, 5},
				l:     0,
				r:     5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.input, tt.args.l, tt.args.r)
		})
	}
}

func Test_heapSort(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				input: []int{1, 4, 2, 1, 2, 5, 6, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heapSort(tt.args.input)
		})
	}
}
